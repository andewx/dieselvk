package dieselvk

import (
	"fmt"

	vk "github.com/vulkan-go/vulkan"
)

/* Descriptor Pool & Set Manager API - common techniques are to allocate a set of descriptor pools per frame
where pool allocation failures create new pools for further memory allocation, therefore we need to be able
to handle to pool allocation regime and the set bindings*/

type CoreDescriptorPool struct {
	pool vk.DescriptorPool
}

type CoreDescriptorBuffer struct {
	binding int
	dtype   int
	buffer  *CoreBuffer
}

//Layout must have an associated pipeline to be bound against
//Buffer binding locations are mapped to the descriptor sets
type CoreDescriptor struct {
	set       []vk.DescriptorSet
	p_buffers []CoreDescriptorBuffer
	p_layout  []vk.DescriptorSetLayout
}

func NewDescriptorPool(handle vk.Device, handles_per_type int, types []int) (*CoreDescriptorPool, error) {

	core := CoreDescriptorPool{}
	sizes := make([]vk.DescriptorPoolSize, len(types))

	//Create a pool that will hold N buffer types with n buffers per buffer type ... (UNIFORM/STORED/SHADER etc..)
	for i := 0; i < len(sizes); i++ {
		sizes[i].DescriptorCount = uint32(handles_per_type)
		sizes[i].Type = vk.DescriptorType(types[i])
	}

	pool_info := vk.DescriptorPoolCreateInfo{}
	pool_info.SType = vk.StructureTypeDescriptorPoolCreateInfo
	pool_info.Flags = 0
	pool_info.MaxSets = 10
	pool_info.PoolSizeCount = uint32(len(sizes))
	pool_info.PPoolSizes = sizes

	//Create the Pool
	if res := vk.CreateDescriptorPool(handle, &pool_info, nil, &core.pool); res != vk.Success {
		return &core, NewError(res)
	}

	return &core, nil
}

func (core *CoreDescriptor) GetLayouts() []vk.DescriptorSetLayout {
	return core.p_layout
}

//Set must be allocated from a descriptor set layout. set is nil
func NewCoreDescriptor(handle vk.Device, layout []vk.DescriptorSetLayout) (*CoreDescriptor, error) {
	desc := &CoreDescriptor{}
	desc.p_layout = layout
	desc.p_buffers = make([]CoreDescriptorBuffer, len(layout))
	desc.set = make([]vk.DescriptorSet, len(layout))
	return desc, nil
}

//The passed core descriptor must have a valid associated layout
func (core *CoreDescriptorPool) AllocateSets(handle vk.Device, desc *CoreDescriptor) error {

	//Allocate Descriptor Set Bindings
	var err error

	//Pre-condition - vkDescriptorLayout is initialized
	if desc.p_layout == nil {
		err = fmt.Errorf("desc *CoreDescriptor must be have a valid and initialized layout in order to initialize the associate vkDescriptorSet vulkan cgo type")
		return err
	}

	//Descriptor Layout
	layout := desc.GetLayouts()

	//Heap allocate Set
	cset := make([]vk.DescriptorSet, len(layout))

	alloc_info := vk.DescriptorSetAllocateInfo{}
	alloc_info.DescriptorPool = core.pool
	alloc_info.SType = vk.StructureTypeDescriptorSetAllocateInfo
	alloc_info.DescriptorSetCount = uint32(len(layout))
	alloc_info.PSetLayouts = layout
	if res := vk.AllocateDescriptorSets(handle, &alloc_info, &cset[0]); res != vk.Success {
		NewError(res)
	}

	//Hold pointer
	desc.set = cset

	return nil
}

func (core *CoreDescriptorPool) Destroy(handle vk.Device) {
	vk.DestroyDescriptorPool(handle, core.pool, nil)
}

func (core *CoreDescriptor) GatherSets() ([]vk.DescriptorSet, error) {

	return core.set, nil
}

//Each Descriptor Layout points to its own binding description and has its own binding.
func NewDescriptorLayouts(handle vk.Device, location []uint32, descriptor_type []vk.DescriptorType, stage_flags vk.ShaderStageFlags) ([]vk.DescriptorSetLayout, error) {
	var layout_bindings [][]vk.DescriptorSetLayoutBinding
	var descriptor_layout []vk.DescriptorSetLayout
	var layout_infos []vk.DescriptorSetLayoutCreateInfo
	size := len(descriptor_type)

	layout_bindings = make([][]vk.DescriptorSetLayoutBinding, size)
	descriptor_layout = make([]vk.DescriptorSetLayout, size)
	layout_infos = make([]vk.DescriptorSetLayoutCreateInfo, size)

	for i := 0; i < size; i++ {
		layout_bindings[i] = make([]vk.DescriptorSetLayoutBinding, 1)
		layout_bindings[i][0].Binding = location[i]
		layout_bindings[i][0].DescriptorCount = 1
		layout_bindings[i][0].DescriptorType = descriptor_type[i]
		layout_bindings[i][0].StageFlags = stage_flags
	}

	for i := 0; i < size; i++ {
		layout_infos[i].SType = vk.StructureTypeDescriptorSetLayoutCreateInfo
		layout_infos[i].BindingCount = uint32(1)
		layout_infos[i].PBindings = layout_bindings[i]
		layout_infos[i].Flags = 0
	}

	for i := 0; i < size; i++ {
		if res := vk.CreateDescriptorSetLayout(handle, &layout_infos[i], nil, &descriptor_layout[i]); res != vk.Success {
			return descriptor_layout, NewError(res)
		}
	}

	return descriptor_layout, nil

}

//Appends new buffer and binds the buffer data type and location
func (core *CoreDescriptor) AddBuffer(handle vk.Device, binding int, data_type int, set_id int, buffer CoreBuffer) {

	core.p_buffers = append(core.p_buffers, CoreDescriptorBuffer{
		binding: binding,
		dtype:   data_type,
		buffer:  &buffer,
	})

	index := len(core.p_buffers) - 1

	binfo := make([]vk.DescriptorBufferInfo, 1)
	binfo[0].Buffer = core.p_buffers[index].buffer.buffer[0]
	binfo[0].Offset = vk.DeviceSize(0)
	binfo[0].Range = core.p_buffers[index].buffer.reqs.Size

	write := make([]vk.WriteDescriptorSet, 1)
	write[0].SType = vk.StructureTypeWriteDescriptorSet
	write[0].DstBinding = uint32(core.p_buffers[index].binding)
	write[0].DstSet = core.set[set_id]
	write[0].DescriptorCount = 1
	write[0].DescriptorType = vk.DescriptorType(vk.DescriptorTypeUniformBuffer)
	write[0].PBufferInfo = binfo
	vk.UpdateDescriptorSets(handle, 1, write, 0, nil)

}
