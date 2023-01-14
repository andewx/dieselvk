package dieselvk

import (
	"fmt"
	"os"

	vk "github.com/vulkan-go/vulkan"
)

/* Device Instance is minimized Vulkan instance with no display output renderpasses or render-targets*/
type CoreDeviceInstance struct {

	//Instances
	instance            *vk.Instance
	instance_extensions BaseInstanceExtensions
	device_extensions   BaseDeviceExtensions
	validation_layers   BaseLayerExtensions
	name                string

	//Core Allocatore
	allocator *CoreAllocator

	//Single Logical Device for the instance
	logical_device      *CoreDevice
	properties          *Usage
	queues              *CoreQueue
	device_queue        *vk.Queue
	device_queue_family uint32

	//Swapchain Synchronization
	recycled_semaphores []vk.Semaphore
	cmds                []vk.CommandBuffer

	//Buffers
	uniform_buffers map[string]*CoreBuffer
	vertex_buffers  map[string]*CoreBuffer

	//Maps program id's to renderpasses & pipelines
	programs map[string]string
	shaders  *CoreShader

	//Descriptor Set Globals
	global_descriptor_pool    *CoreDescriptorPool
	global_descriptor_layouts map[string][]vk.DescriptorSetLayout
	frame_descriptor_sets     *CoreDescriptor
}

//Creates a new core instance from the given structure and attaches the instance to a primary graphics compatbible device
func NewCoreDeviceInstance(instance vk.Instance, name string, instance_exenstions BaseInstanceExtensions, validation_extensions BaseLayerExtensions, device_extensions []string) (*CoreDeviceInstance, error) {
	var core CoreDeviceInstance
	var err error

	//Core Extensions
	core.instance_extensions = instance_exenstions
	core.validation_layers = validation_extensions

	core.instance = &instance
	core.logical_device = &CoreDevice{}
	core.logical_device.key = name
	core.name = name
	core.programs = make(map[string]string, 4)
	core.recycled_semaphores = make([]vk.Semaphore, 0)
	core.uniform_buffers = make(map[string]*CoreBuffer, MAX_UNIFORM_BUFFERS)
	core.vertex_buffers = make(map[string]*CoreBuffer, MAX_UNIFORM_BUFFERS)
	core.global_descriptor_layouts = make(map[string][]vk.DescriptorSetLayout)
	core.cmds = make([]vk.CommandBuffer, 0)
	core.shaders = NewCoreShader()

	var gpu_count uint32
	var gpus []vk.PhysicalDevice

	ret := vk.EnumeratePhysicalDevices(*core.instance, &gpu_count, nil)

	if gpu_count == 0 {
		Fatal(fmt.Errorf("func (core *CoreRenderInstance)Init() -- No valid physical devices found, count is 0\n"))
	}

	gpus = make([]vk.PhysicalDevice, gpu_count)

	ret = vk.EnumeratePhysicalDevices(*core.instance, &gpu_count, gpus)

	if ret != vk.Success {
		Fatal(fmt.Errorf("func (core *CoreRenderInstance)Ini() -- Unable to query physical devices\n"))
	}

	core.logical_device.physical_devices = append(core.logical_device.physical_devices, gpus...)

	//Select Valid Device By Desired Queue Properties
	has_device := false
	for index := 0; index < int(gpu_count); index++ {
		mGPU := gpus[index]
		flag_bits := uint32(vk.QueueComputeBit)
		if core.is_valid_device(&mGPU, flag_bits) {
			core.logical_device.selected_device = mGPU
			core.logical_device.selected_device_properties = &vk.PhysicalDeviceProperties{}
			core.logical_device.selected_device_memory_properties = &vk.PhysicalDeviceMemoryProperties{}
			has_device = true
			break
		}
	}

	if !has_device {
		fmt.Errorf("Could not find suitable GPU device for graphics and presentation\n")
		os.Exit(1)
	}

	//Load in extensions
	core.device_extensions = *NewBaseDeviceExtensions(device_extensions, []string{}, core.logical_device.selected_device)

	//Gather device properties
	vk.GetPhysicalDeviceProperties(core.logical_device.selected_device, core.logical_device.selected_device_properties)
	core.logical_device.selected_device_properties.Deref()
	vk.GetPhysicalDeviceMemoryProperties(core.logical_device.selected_device, core.logical_device.selected_device_memory_properties)
	core.logical_device.selected_device_memory_properties.Deref()

	// Select device extensions
	core.device_extensions = *NewBaseDeviceExtensions(core.device_extensions.wanted, []string{}, core.logical_device.selected_device)
	has_extensions, ext_string := core.device_extensions.HasWanted()

	if !has_extensions {
		fmt.Printf("Vulkan Missing Device Extensions %s", ext_string)
	} else {
		fmt.Printf("Vulkan Device Extensions loaded...\n")
	}

	//Bind the suitable device with assigned queues
	core.queues = NewCoreQueue(core.logical_device.selected_device, core.name)
	queue_infos := core.queues.GetCreateInfos()
	dev_extensions := core.device_extensions.GetExtensions()

	//Create Device
	var device vk.Device
	ret = vk.CreateDevice(core.logical_device.selected_device, &vk.DeviceCreateInfo{
		SType:                   vk.StructureTypeDeviceCreateInfo,
		QueueCreateInfoCount:    uint32(len(queue_infos)),
		PQueueCreateInfos:       queue_infos,
		EnabledExtensionCount:   uint32(len(dev_extensions)),
		PpEnabledExtensionNames: safeStrings(dev_extensions),
		EnabledLayerCount:       uint32(len(core.validation_layers.GetExtensions())),
		PpEnabledLayerNames:     safeStrings(core.validation_layers.GetExtensions()),
	}, nil, &device)

	if ret != vk.Success {
		if ret == vk.ErrorFeatureNotPresent || ret == vk.ErrorExtensionNotPresent {
			fmt.Printf("Error certain device features may not be available on the requested GPU device\n%s\nExiting...", dev_extensions)
			os.Exit(1)
		} else {
			fmt.Printf("Fatal error creating device device not found or device state invalid\nExiting...")
			os.Exit(1)
		}
	}

	core.logical_device.handle = device

	if err = core.CreateQueues(); err != nil {
		Fatal(err)
	}

	//Create Pipeline and Descriptor Pools - Create Descriptor Pool with Types and Handles Per Type
	var layout []vk.DescriptorSetLayout
	types := make([]int, MAX_UNIFORM_BUFFERS)
	for i := 0; i < MAX_UNIFORM_BUFFERS; i++ {
		types[i] = int(vk.DescriptorTypeUniformBuffer)
	}

	layout_types := []vk.DescriptorType{vk.DescriptorTypeUniformBuffer, vk.DescriptorTypeUniformBuffer, vk.DescriptorTypeUniformBuffer}

	core.global_descriptor_pool, err = NewDescriptorPool(core.logical_device.handle, DESCRIPTOR_SET_HANDLES, types) //Make pool allocation of 10 Uniform Buffer Types with 3 Descriptor Set Handles Per
	layout, err = NewDescriptorLayouts(core.logical_device.handle, []uint32{0, 0, 0}, layout_types, vk.ShaderStageFlags(vk.PipelineStageVertexShaderBit|vk.PipelineStageFragmentShaderBit))
	core.global_descriptor_layouts["default"] = layout
	//Descriptor Sets set to a default Uniform buffer and Vertex Shader stage. Parameterize for engine flexibility
	core.frame_descriptor_sets, err = NewCoreDescriptor(core.logical_device.handle, core.global_descriptor_layouts["default"])
	core.global_descriptor_pool.AllocateSets(core.logical_device.handle, core.frame_descriptor_sets)

	core.allocator, err = NewCoreAllocator(core.logical_device.selected_device, device, 1024, 1)
	return &core, err
}

/*Creates device queue objects */
func (core *CoreDeviceInstance) CreateQueues() error {
	var err error
	core.queues.CreateQueues(core.logical_device.handle)
	found, q_handle, family := core.queues.BindGraphicsQueue(core.logical_device.handle)
	core.device_queue_family = uint32(family)

	if !found {
		err = fmt.Errorf("No valid device queues found\n")
	}
	core.device_queue = q_handle
	return err
}

func (core CoreDeviceInstance) SetupCommands() {
	return
}

func (core *CoreDeviceInstance) AddLayoutBuffer(data []float32, name string, usage vk.BufferUsageFlags) {
	d := Ptr(data)
	mdata := &d
	bf := vk.BufferUsageFlags(usage)
	core.uniform_buffers[name] = NewLayoutBuffer(core.logical_device.handle, core.logical_device.selected_device, uint32(len(data)*4), int32(bf))
	mem_size := core.uniform_buffers[name].reqs.Size
	min_align := core.uniform_buffers[name].reqs.Alignment

	if mem_ref, err := core.allocator.Allocate(int(mem_size), int(min_align)); err == nil {
		if err := core.allocator.Map(mdata, core.uniform_buffers[name].buffer[0], core.logical_device.handle, mem_ref, FLOAT32); err != nil {
			fmt.Errorf("Failed to bind buffer %s\n", name)
		}
	}
}

func (core *CoreDeviceInstance) GetLayoutBuffers() map[string]*CoreBuffer {
	return core.uniform_buffers
}

//Active uniform buffers are bound to the descriptor sets and given
func (core *CoreDeviceInstance) BindUniforms(uniforms map[string]*CoreBuffer, binding []int) error {

	for index := 0; index < 3; index++ {
		for _, uniform := range uniforms {
			core.frame_descriptor_sets.AddBuffer(core.logical_device.handle, binding[index], int(vk.DescriptorTypeUniformBuffer), index, *uniform)
		}
	}
	return nil
}

func (core CoreDeviceInstance) AddShaderPath(path string, shader_type int) {
	core.shaders.AddShaderPath(path, shader_type)
}

/*Adds a program to the vulkan instance*/
func (core *CoreDeviceInstance) NewProgram(paths []string, name string) error {
	core.shaders.CreateProgram(name, core, paths)
	return nil
}

/*Adds vertex buffer with allocated memory to the vulkan instance*/
func (core *CoreDeviceInstance) AddVertexBuffer(data []float32, name string) {
	prototype := Vertex{}
	d := Ptr(data)
	mdata := &d
	bf := vk.BufferUsageFlags(vk.BufferUsageVertexBufferBit)
	core.vertex_buffers[name] = NewCoreVertexBuffer(core.logical_device.handle, core.logical_device.selected_device, uint32(len(data)*4), int32(bf), prototype)
	mem_size := core.vertex_buffers[name].reqs.Size
	min_align := core.vertex_buffers[name].reqs.Alignment

	if mem_ref, err := core.allocator.Allocate(int(mem_size), int(min_align)); err == nil {
		core.allocator.Map(mdata, core.vertex_buffers[name].buffer[0], core.logical_device.handle, mem_ref, int(core.vertex_buffers[name].reqs.MemoryTypeBits))
	}
}

func (core *CoreDeviceInstance) AddPipeline(name string, program_name string, buffer CoreBuffer, pass string) *CorePipeline {
	return nil
}

func (core CoreDeviceInstance) AddRenderPass(name string) *CoreRenderPass {
	return nil
}

func (core CoreDeviceInstance) NewSwapchain() *CoreSwapchain {
	return nil
}

func (core CoreDeviceInstance) GetVertexBuffer(name string) *CoreBuffer {
	return core.vertex_buffers[name]
}

func (core *CoreDeviceInstance) GetHandle() vk.Device {
	return core.logical_device.handle
}

func (core *CoreDeviceInstance) GetPhysicalDevice() vk.PhysicalDevice {
	return core.logical_device.selected_device
}

func (core *CoreDeviceInstance) Update(delta_time float32) {

	vk.QueueWaitIdle(*core.device_queue)

	return
}

func (core *CoreDeviceInstance) release() {
	core.Destroy()
	for _, buffer := range core.uniform_buffers {
		buffer.Destroy(core.logical_device.handle)
	}
}

func (core *CoreDeviceInstance) Destroy() {

	vk.DeviceWaitIdle(core.logical_device.handle)

	for _, shader := range core.shaders.shader_programs {
		vk.DestroyShaderModule(core.logical_device.handle, *shader.vertex_shader_modules, nil)
		vk.DestroyShaderModule(core.logical_device.handle, *shader.fragment_shader_modules, nil)
	}

	for _, vertex_buffer := range core.vertex_buffers {
		vertex_buffer.Destroy(core.logical_device.handle)
	}

	core.global_descriptor_pool.Destroy(core.logical_device.handle)

	for _, layouts := range core.global_descriptor_layouts {
		for _, layout := range layouts {
			vk.DestroyDescriptorSetLayout(core.logical_device.handle, layout, nil)
		}
	}

	for _, buffer := range core.uniform_buffers {
		buffer.Destroy(core.logical_device.handle)
	}

	core.allocator.Destroy(core.logical_device.handle)

	vk.DestroyDevice(core.logical_device.handle, nil)
}

func (core *CoreDeviceInstance) is_valid_device(device *vk.PhysicalDevice, flags uint32) bool {

	q := NewCoreQueue(*device, "Default")
	return q.IsDeviceSuitable(flags)
}

func (core *CoreDeviceInstance) AllocatorUsage() {
	fmt.Printf(core.allocator.Usage())
}
