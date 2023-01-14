//Vulkan Memory Allocator with Fragmentation Statistics
package dieselvk

import (
	"fmt"
	"unsafe"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
	vk "github.com/vulkan-go/vulkan"
)

/*
DieselVK Memory allocation does basic memory constraint handling with the GPU. From a
wide perspective the allocator views the GPU dedicated memory as a limited resource. Allocated
memory blocks therefore can be mapped as dedicated memory but priority access directs GPU eviction
to host visible VM memory.

Different memory allocation strategies can be employed by user implemented interfaces. We implement
three distinct interfaces. CoreAllocator which is a default Host Visible Memory allocation. CoreStageAllocator
which is a dedicated GPU memory backed by a Host Visible staging buffer. CoreImageAllocator which is similar
to core stage allocator but utilizes the vkImage* API bindings.

Staged allocators will fall back to host visible memory when dedicated pool memory is filled. With the exception
of Image Staged Allocators.

An CoreOptimalAllocator priority ranks staged buffers and uses priority heuristics to move staged buffers into
GPU dedicated memory, where larger buffer chunks are preferred.

It should be noted that the static initialization functions for each "class" are responsible for querying physical
device memory types for a suitable memory layout

Note it is probably advantageous that we run a defrag / block analysis threaded routine to monitor the memory composition
so we run it in the interface with Run() with an int channel to monitor memory coherency so that main thread allocations
can only worry about allocations which will mark the memory pool as incoherent/coherent. The

*/
const (
	MARK_FREE          = 0
	MARK_USED          = 1
	MARK_DEFRAG        = 2
	POOL_INCOHRENT     = 0
	POOL_COHERENT      = 1
	POOL_DEFAULT_PAGES = 1
	FLOAT32            = 0
	INT32              = 1
	BYTES              = 2
	MEM_REF            = 0
	FREE_REF           = 1
)

type Allocator interface {
	Allocate(size int) error
	Map(data *unsafe.Pointer, buffer vk.Buffer, handle vk.Device, key int, type_memory int) error
	Clean()
	Resize(desired_size int) error
	Stats() map[string]string
	Destroy()
	Run(x chan int)
}

type PoolStats struct {
	mean_frag     float32
	frag_variance float32
	mean_free     float32
	max_free      float32
	mean_usage    float32
}

type Block64 struct {
	offset uint64
	size   uint64
	flag   uint8
}

type MemHeap struct {
	size   uint64
	budget uint64
	usage  uint64
}

type MemPage struct {
	device_mem      vk.DeviceMemory
	mem_blocks      []Block64
	free_blocks     []Block64
	mem_block_refs  map[int]BlockRef
	free_block_refs map[int]BlockRef
	tree_mem        rbt.Tree
	tree_free       rbt.Tree
	size            uint64
	index           int
}

type MemRef struct {
	key   int
	page  int
	usage int
}

type BlockRef struct {
	page_id      int
	mem_block_id int
}

//Device Memory Pooling we use red black trees and a backing memory block reference array to manage block queries
type MemPool struct {
	pages        []MemPage
	alignment    uint32
	size         uint64
	heap         MemHeap
	status       int
	vulkan_flags vk.MemoryType
	mem_info     vk.MemoryAllocateInfo
}

//Allocator creates memory pools for each type of available vulkan memory types
type CoreAllocator struct {
	pool MemPool
}

func NewCoreAllocator(physical vk.PhysicalDevice, handle vk.Device, max_pool_size uint64, pages int) (*CoreAllocator, error) {
	core := CoreAllocator{}
	core.pool = MemPool{
		pages:  make([]MemPage, 0),
		size:   max_pool_size,
		heap:   MemHeap{},
		status: POOL_COHERENT,
	}

	mem_props := vk.PhysicalDeviceMemoryProperties{}
	dev_props := vk.PhysicalDeviceProperties{}
	vk.GetPhysicalDeviceMemoryProperties(physical, &mem_props)
	vk.GetPhysicalDeviceProperties(physical, &dev_props)
	mem_props.Deref()
	dev_props.Deref()
	mem_index := int32(0)
	core.pool.alignment = uint32(dev_props.Limits.BufferImageGranularity)

	//Get memory type index - Here we demand a default operating mode of HostVisible and HostCoherent to handle memory usage
	for i := 0; i < int(mem_props.MemoryTypeCount); i++ {
		mem_type := mem_props.MemoryTypes[i]
		mem_type.Deref()
		if match_memory_desired(int32(i), mem_type.PropertyFlags, int32(vk.MemoryPropertyHostVisibleBit|vk.MemoryPropertyHostCoherentBit)) {
			mem_index = int32(i)
		}
	}

	//Get the buffer image granularity (buffer alignment cues)
	core.pool.vulkan_flags = mem_props.MemoryTypes[mem_index]

	//Get the heap size
	heap_index := mem_props.MemoryTypes[mem_index].HeapIndex
	mem_heap := mem_props.MemoryHeaps[heap_index]
	mem_heap.Deref()
	heap_size := mem_heap.Size

	//Warning we aren't tracking
	if max_pool_size > uint64(heap_size) {
		return &core, fmt.Errorf("Error Requested Memory Pool size greater than available heap size\n")
	}

	//We internally track the heap space available while budgeting helps us determine when client side memory is full
	core.pool.heap.size = uint64(heap_size)
	core.pool.heap.budget = uint64(heap_size)
	core.pool.heap.usage = 0

	//Allocate Device Memory for each page
	page_size := core.pool.size / uint64(pages)
	mem_info := vk.MemoryAllocateInfo{}
	mem_info.SType = vk.StructureTypeMemoryAllocateInfo
	mem_info.AllocationSize = vk.DeviceSize(page_size)
	mem_info.MemoryTypeIndex = uint32(mem_index)
	mem_info.PassRef()
	core.pool.mem_info = mem_info

	for i := 0; i < pages; i++ {
		core.NewMemoryPage(page_size, handle, i)
	}

	return &core, nil
}

//Allocate allocates a new memory binding region and generates the new memory blocks
//block key is page offset in memory space
func (core *CoreAllocator) Allocate(size int, min_align int) (MemRef, error) {

	//Search pages for free memory
	var found bool
	var fnode *rbt.Node
	var mem_block_key int
	var mPage MemPage
	var page_index = 0
	alloc_size := size

	//Match Vulkan Alignment Specifications

	if size < min_align {
		alloc_size = min_align
	} else {
		div := size / min_align
		alloc_size = size + (div+1)*(min_align)
	}

	for i, page := range core.pool.pages {
		node := page.tree_free.Root
		comp_fn := page.tree_free.Comparator
		found, fnode = rbt_search(node, comp_fn, alloc_size)
		if found {
			mem_block_key = fnode.Key.(int)
			mPage = page
			page_index = i
			break
		}
	}

	//Check if found
	if found {
		key := fnode.Key.(int)
		free_block := mPage.free_block_refs[key]
		mem_block := free_block

		//Obtain Block Handles
		block := mPage.free_blocks[mem_block.mem_block_id]
		n_block := Block64{size: uint64(alloc_size), offset: block.offset, flag: MARK_USED}

		//Realign Free Block
		block.size = block.size - uint64(alloc_size)
		block.offset = block.offset + uint64(alloc_size)
		block.flag = MARK_FREE
		mPage.free_blocks[mem_block.mem_block_id] = block

		//Store memory block and block reference
		core.pool.pages[page_index].mem_blocks = append(mPage.mem_blocks, n_block)
		page := mem_block.page_id
		block_index := len(core.pool.pages[page_index].mem_blocks) - 1
		n_block_ref := BlockRef{
			page_id:      page,
			mem_block_id: block_index,
		}

		//Memory block id key is it's offset which is unique
		mem_block_key = int(n_block.offset)
		mPage.mem_block_refs[mem_block_key] = n_block_ref
		mPage.tree_mem.Put(mem_block_key, alloc_size)

		//Remove and store new free memory block
		mPage.tree_free.Remove(key)
		mPage.tree_free.Put(key, block.size)

	} else {
		return MemRef{}, NewError(vk.ErrorOutOfPoolMemory)
	}

	return MemRef{mem_block_key, page_index, MEM_REF}, nil
}

//AllocateFromBuffer(vk.Buffer)

//Map binds the keyed buffer memory and allocates and places data into the GPU visible memory. For dedicated GPU memory
//extend this implementation
func (core *CoreAllocator) Map(data *unsafe.Pointer, buffer vk.Buffer, handle vk.Device, ref MemRef, type_memory int) error {

	page := core.pool.pages[ref.page]
	block_ref := page.mem_block_refs[ref.key]

	if &block_ref == nil {
		return NewError(vk.ErrorMemoryMapFailed)
	}

	//Device memory only has 1 allocation per page
	mem_ref := page.mem_blocks[block_ref.mem_block_id]
	dev_ref := core.pool.pages[block_ref.page_id].device_mem

	res := vk.BindBufferMemory(handle, buffer, dev_ref, vk.DeviceSize(mem_ref.offset))

	if res != vk.Success {
		return NewError(res)
	}

	if type_memory == FLOAT32 {
		map_memory_float32(data, handle, dev_ref, int(mem_ref.size), int(mem_ref.offset))
	}

	return nil
}

//Get memory reference from description structure
func (core *CoreAllocator) GetMemoryRef(ref MemRef) BlockRef {
	page := core.pool.pages[ref.page]
	if ref.usage == MEM_REF {
		return page.mem_block_refs[ref.key]
	}
	return page.free_block_refs[ref.key]
}

//Cleanup checks current allocated memory blocks that have been free'd in runtime. This approach is somewhat
//Naive since we are blocking the CPU for the background task of cleaning the memory pool.
func (core *CoreAllocator) Clean() {

	//Check memory blocks for usage
	for pg_index, page := range core.pool.pages {
		for _, r_block := range page.mem_block_refs {
			ref := page.mem_blocks[r_block.mem_block_id]
			if ref.flag == MARK_FREE {
				//Return memory back to the free tree and consolidate memory if neccessary
				free_block := Block64{size: ref.size, offset: ref.offset, flag: MARK_FREE}
				free_block_ref := BlockRef{}
				page.free_blocks = append(page.free_blocks, free_block)
				free_block_ref.page_id = r_block.page_id
				free_block_ref.mem_block_id = len(page.free_blocks) - 1

				//Check for free blocks occupying relevant memory keys
				free_blocks := make([]MemRef, 0)
				link_found := true
				current_block := free_block
				fb_size := uint64(0)
				for link_found {
					offset_id := int(current_block.offset + current_block.size)
					if chain_free_block, ok := page.free_block_refs[offset_id]; ok {
						free_blocks = append(free_blocks, MemRef{offset_id, pg_index, FREE_REF})
						current_block = page.free_blocks[chain_free_block.mem_block_id]
						fb_size += (current_block.size)
					} else {
						link_found = false
					}
				}

				free_block.size += fb_size

				//Remove the free blocks by deleting all references
				for _, ref := range free_blocks {
					block_ref := core.GetMemoryRef(ref)
					page.free_blocks = append(page.free_blocks[:block_ref.mem_block_id], page.free_blocks[block_ref.mem_block_id+1:]...)
					page.tree_free.Remove(ref.key)
				}

				//Add block to pooling
				page.free_block_refs[int(free_block.offset)] = free_block_ref
				page.tree_free.Put(int(free_block.offset), free_block.size)

				//Remove memory reference
				core.pool.pages[r_block.page_id].mem_blocks = append(core.pool.pages[r_block.page_id].mem_blocks[:r_block.mem_block_id], core.pool.pages[r_block.page_id].mem_blocks[r_block.mem_block_id+1:]...)
				delete(page.mem_block_refs, int(free_block.offset))
			}
		}
	}
}

//Creates a new memory page of the desired size
func (core *CoreAllocator) NewMemoryPage(desired_size uint64, handle vk.Device, index int) error {
	page := MemPage{}
	page.free_block_refs = make(map[int]BlockRef, 1)
	page.free_blocks = make([]Block64, 1)
	page.mem_block_refs = make(map[int]BlockRef)
	page.mem_blocks = make([]Block64, 0)
	device_size := vk.DeviceSize(desired_size)
	core.pool.mem_info.AllocationSize = device_size
	core.pool.mem_info.PassRef()
	page.index = index
	page.size = uint64(device_size)

	page.tree_mem = *rbt.NewWithIntComparator()
	page.tree_free = *rbt.NewWithIntComparator()

	alloc_infos := []vk.MemoryAllocateInfo{core.pool.mem_info}
	device_mems := []vk.DeviceMemory{page.device_mem}

	res := vk.AllocateMemory(handle, &alloc_infos[0], nil, &device_mems[0])
	if res != vk.Success {
		return NewError(res)
	}

	page.device_mem = device_mems[0]

	//Store the free blocks
	page.free_blocks[0] = Block64{offset: 0, size: desired_size, flag: MARK_FREE}
	page.free_block_refs[0] = BlockRef{page_id: index, mem_block_id: 0}
	page.tree_free.Put(0, int(page.free_blocks[0].size))

	core.pool.pages = append(core.pool.pages, page)
	return nil

}

//Returns a string map of relevant statistics and can be used by other tools to visually display memory usage
func (core *CoreAllocator) Stats() map[string]string {
	return nil
}

func (core *CoreAllocator) Usage() string {
	var out string
	for i, page := range core.pool.pages {
		out += fmt.Sprintf("Page %d: (%d)bytes\nUsed Memory\n", i, page.size)
		for j, block := range page.mem_blocks {
			out += fmt.Sprintf("    m%d: offset ( %.8d ) size (%.8d )\n", j, block.offset, block.size)
		}
		out += fmt.Sprintf("Free Memory\n")
		for j, block := range page.free_blocks {
			out += fmt.Sprintf("    m%d: offset ( %.8d ) size (%.8d )\n", j, block.offset, block.size)
		}

		out += "\n----------------------\n"

	}
	return out
}

//TODO add func ShowMemoryMap() map[string]string. Outputs JSON string data structure with Memory Block structure

func (core *CoreAllocator) Free() {
	for i, page := range core.pool.pages {
		page.mem_blocks = make([]Block64, 0)
		page.free_blocks = make([]Block64, 0)
		for key := range page.mem_block_refs {
			delete(page.mem_block_refs, key)
		}
		for key := range page.free_block_refs {
			delete(page.free_block_refs, key)
		}
		page.tree_mem.Clear()
		page.tree_free.Clear()

		page.tree_mem = *rbt.NewWithIntComparator()
		page.tree_free = *rbt.NewWithIntComparator()
		page.free_blocks = append(page.free_blocks, Block64{0, page.size, MARK_FREE})
		page.free_block_refs = make(map[int]BlockRef, 1)
		page.mem_block_refs = make(map[int]BlockRef, 1)
		page.free_block_refs[0] = BlockRef{i, 0}
		page.tree_free.Put(0, page.size)

	}
}

//Destroys all memory refernces and tree and sets a free block instance to the page size
func (core *CoreAllocator) Destroy(handle vk.Device) {
	core.Free()
	for _, page := range core.pool.pages {
		vk.FreeMemory(handle, page.device_mem, nil)
	}

}

//Run is a static analyzer that can be used to identify new free blocks upon memory allocations. Pool state should
//be incohrent until the static analyzer method has identified all free memory blocks.
func (core *CoreAllocator) Run(x chan int) {
	core.Clean()
}

//Host visible memory mapping
func map_memory_float32(data *unsafe.Pointer, handle vk.Device, device_mem vk.DeviceMemory, size int, offset int) {
	pointer_mem := uintptr(0)
	p_mem := Ptr(pointer_mem)
	data_slice := unsafe.Slice((*float32)(*data), size)
	vk.MapMemory(handle, device_mem, vk.DeviceSize(offset), vk.DeviceSize(size), 0, &p_mem)
	dest_slice := unsafe.Slice((*float32)(p_mem), size)
	copy(dest_slice, data_slice)
	vk.UnmapMemory(handle, device_mem)
}

func map_memory_int32(data *unsafe.Pointer, handle vk.Device, device_mem vk.DeviceMemory, size int, offset int) {
	pointer_mem := uintptr(0)
	p_mem := Ptr(pointer_mem)
	data_slice := unsafe.Slice((*int32)(*data), size)
	vk.MapMemory(handle, device_mem, vk.DeviceSize(offset), vk.DeviceSize(size), 0, &p_mem)
	dest_slice := unsafe.Slice((*int32)(p_mem), size)
	copy(dest_slice, data_slice)
	vk.UnmapMemory(handle, device_mem)
}

func match_memory_desired(index int32, properties vk.MemoryPropertyFlags, desired int32) bool {
	pad := int32(properties) & desired
	if (pad) == desired {
		return true
	}
	return false
}

//Search Tree Min Int
func rbt_search(node *rbt.Node, comp_fn utils.Comparator, search int) (bool, *rbt.Node) {
	found := false

	//Search Red-Black Tree for appropriate memory
	for node != nil {
		if comp_fn(search, node.Value) <= 0 {
			found = true
			return found, node
		} else {
			if node.Right != nil {
				node = node.Right
			} else {
				return found, node
			}
		}
	}
	return found, node
}
