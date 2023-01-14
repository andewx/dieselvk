package dieselvk

import (
	"fmt"
	"os"
	"unsafe"

	vk "github.com/vulkan-go/vulkan"
)

const (
	SWAPCHAIN_COUNT            = 3
	MAX_UNIFORM_BUFFERS        = 10
	MAX_DESCRIPTOR_SET_BUFFERS = 10
	DESCRIPTOR_SET_HANDLES     = 3
)

type SPIRV_Constants struct {
	frame_delta float32
}

//Swapchain synchronization
type PerFrame struct {
	pool                  *CorePool
	command               []vk.CommandBuffer
	fence                 []vk.Fence
	image_acquired        []vk.Semaphore
	queue_complete        []vk.Semaphore
	local_descriptor_sets []*CoreDescriptor
}

var update_step float32

func NewPerFrame(core *CoreRenderInstance, frame int) (PerFrame, error) {
	var err error
	m_frame := PerFrame{}

	m_frame.command = make([]vk.CommandBuffer, 1)
	m_frame.fence = make([]vk.Fence, 1)
	m_frame.image_acquired = make([]vk.Semaphore, 1)
	m_frame.queue_complete = make([]vk.Semaphore, 1)
	m_frame.pool, err = NewCorePool(&core.logical_device.handle, core.render_queue_family)

	//Command buffers
	vk.AllocateCommandBuffers(core.logical_device.handle, &vk.CommandBufferAllocateInfo{
		SType:              vk.StructureTypeCommandBufferAllocateInfo,
		CommandPool:        m_frame.pool.pool,
		Level:              vk.CommandBufferLevelPrimary,
		CommandBufferCount: uint32(1),
	}, m_frame.command)

	//Create Fence
	vk.CreateFence(core.logical_device.handle, &vk.FenceCreateInfo{
		SType: vk.StructureTypeFenceCreateInfo,
		PNext: nil,
		Flags: vk.FenceCreateFlags(vk.FenceCreateSignaledBit),
	}, nil, &m_frame.fence[0])

	//Create Semaphores
	vk.CreateSemaphore(core.logical_device.handle, &vk.SemaphoreCreateInfo{
		SType: vk.StructureTypeSemaphoreCreateInfo,
		Flags: vk.SemaphoreCreateFlags(0x00000000),
	}, nil, &m_frame.image_acquired[0])

	//Create Semaphores
	vk.CreateSemaphore(core.logical_device.handle, &vk.SemaphoreCreateInfo{
		SType: vk.StructureTypeSemaphoreCreateInfo,
		Flags: vk.SemaphoreCreateFlags(0x00000000),
	}, nil, &m_frame.queue_complete[0])

	return m_frame, err

}

//Core instance API interface
type CoreInstance interface {
	AddPipeline(name string, program string, buffer CoreBuffer, renderpass string) *CorePipeline
	AddShaderPath(path string, shader_type int)
	AddRenderPass(name string) *CoreRenderPass
	AddVertexBuffer(data []float32, name string)
	CreateQueues() error
	Destroy()
	GetHandle() vk.Device
	GetPhysicalDevice() vk.PhysicalDevice
	GetVertexBuffer(name string) *CoreBuffer
	Update(ts float32)
	NewProgram(paths []string, name string) error
	NewSwapchain() *CoreSwapchain
	SetupCommands()
	AllocatorUsage()
	AddLayoutBuffer(data []float32, name string, usage vk.BufferUsageFlags)
	GetLayoutBuffers() map[string]*CoreBuffer
	BindUniforms(uniforms map[string]*CoreBuffer, binding []int) error
}

type CoreRenderInstance struct {

	//Instances
	instance            *vk.Instance
	instance_extensions BaseInstanceExtensions
	device_extensions   BaseDeviceExtensions
	validation_layers   BaseLayerExtensions
	name                string
	allocator           *CoreAllocator

	//Single Logical Device for the instance
	logical_device      *CoreDevice
	properties          *Usage
	display             *CoreDisplay
	queues              *CoreQueue
	render_queue        *vk.Queue
	render_queue_family uint32

	//Swap chain handles
	swapchain     *CoreSwapchain
	per_frame     []PerFrame
	current_frame int

	//Swapchain Synchronization
	recycled_semaphores []vk.Semaphore

	//Buffers
	uniform_buffers map[string]*CoreBuffer
	vertex_buffers  map[string]*CoreBuffer

	//Pipelines and renderpasses
	pipeline     *CorePipeline
	renderpasses map[string]*CoreRenderPass
	Builders     map[string]*PipelineBuilder

	//Maps program id's to renderpasses & pipelines
	programs map[string]string
	shaders  *CoreShader

	//Descriptor Set Globals
	global_descriptor_pool    *CoreDescriptorPool
	global_descriptor_layouts map[string][]vk.DescriptorSetLayout
	frame_descriptor_sets     *CoreDescriptor

	//Push Constant For Now
	pconstant []SPIRV_Constants
}

//Creates a new core instance from the given structure and attaches the instance to a primary graphics compatbible device
func NewCoreRenderInstance(instance vk.Instance, name string, instance_exenstions BaseInstanceExtensions, validation_extensions BaseLayerExtensions, device_extensions []string, display *CoreDisplay) (*CoreRenderInstance, error) {
	var core CoreRenderInstance
	var err error
	update_step = 0.01

	//Core Extensions
	core.instance_extensions = instance_exenstions
	core.validation_layers = validation_extensions

	core.display = display
	core.instance = &instance
	core.logical_device = &CoreDevice{}
	core.logical_device.key = name
	core.name = name
	core.renderpasses = make(map[string]*CoreRenderPass, 4)
	core.programs = make(map[string]string, 4)
	core.recycled_semaphores = make([]vk.Semaphore, 0)
	core.uniform_buffers = make(map[string]*CoreBuffer, MAX_UNIFORM_BUFFERS)
	core.vertex_buffers = make(map[string]*CoreBuffer, MAX_UNIFORM_BUFFERS)
	core.Builders = make(map[string]*PipelineBuilder, 1)
	core.global_descriptor_layouts = make(map[string][]vk.DescriptorSetLayout)

	core.pconstant = make([]SPIRV_Constants, 1)

	core.shaders = NewCoreShader()

	if display.surface == nil {
		surfPtr, err := display.window.CreateWindowSurface(instance, nil)
		if err != nil {
			fmt.Printf("Error creating window surface object")
			display.surface = vk.NullSurface
		}
		display.surface = vk.SurfaceFromPointer(surfPtr)
	}

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
		flag_bits := uint32(vk.QueueGraphicsBit)
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
		return &core, err
	}

	//New Core Allocator - Allocate 1KB Allocator Heap
	core.allocator, err = NewCoreAllocator(core.logical_device.selected_device, core.logical_device.handle, 1024, 1)

	//Pipeline and Descriptor Set Configuration - Ideally this is pre-configured and determined from SPIR-V reflection from the shaders and
	//user defined pipeline layouts and supports multiple pipeline configuration
	var descriptor_layouts []vk.DescriptorSetLayout
	pool_types := []int{int(vk.DescriptorTypeUniformBuffer), int(vk.DescriptorTypeUniformBufferDynamic), int(vk.DescriptorTypeStorageBuffer), int(vk.DescriptorTypeStorageBufferDynamic), int(vk.DescriptorTypeUniformTexelBuffer)}
	layout_types := []vk.DescriptorType{vk.DescriptorTypeUniformBuffer, vk.DescriptorTypeUniformBuffer, vk.DescriptorTypeUniformBuffer}
	core.global_descriptor_pool, err = NewDescriptorPool(core.logical_device.handle, DESCRIPTOR_SET_HANDLES, pool_types) //Make pool allocation of 10 Uniform Buffer Types with 3 Descriptor Set Handles Per
	descriptor_layouts, err = NewDescriptorLayouts(core.logical_device.handle, []uint32{0, 0, 0}, layout_types, vk.ShaderStageFlags(vk.ShaderStageVertexBit))
	core.global_descriptor_layouts["default"] = descriptor_layouts
	//Descriptor Sets set to a default Uniform buffer and Vertex Shader stage. Parameterize for engine flexibility
	core.frame_descriptor_sets, err = NewCoreDescriptor(core.logical_device.handle, core.global_descriptor_layouts["default"])
	core.global_descriptor_pool.AllocateSets(core.logical_device.handle, core.frame_descriptor_sets)
	core.pipeline = NewCorePipeline(&core, "pipe0", core.global_descriptor_layouts["default"])

	return &core, err
}

/*Creates device queue objects */
func (core *CoreRenderInstance) CreateQueues() error {
	var err error
	core.queues.CreateQueues(core.logical_device.handle)
	found, q_handle, family := core.queues.BindGraphicsQueue(core.logical_device.handle)
	core.render_queue_family = uint32(family)

	if !found {
		err = fmt.Errorf("No valid device queues found\n")
	}
	core.render_queue = q_handle
	return err
}
func (core *CoreRenderInstance) AddShaderPath(path string, shader_type int) {
	core.shaders.AddShaderPath(path, shader_type)
}

/*Adds a program to the vulkan instance*/
func (core *CoreRenderInstance) NewProgram(paths []string, name string) error {
	core.shaders.CreateProgram(name, core, paths)
	return nil
}

/*Adds vertex buffer with allocated memory to the vulkan instance*/
func (core *CoreRenderInstance) AddVertexBuffer(data []float32, name string) {
	prototype := Vertex{}
	d := Ptr(data)
	mdata := &d
	bf := vk.BufferUsageFlags(vk.BufferUsageVertexBufferBit)
	core.vertex_buffers[name] = NewCoreVertexBuffer(core.logical_device.handle, core.logical_device.selected_device, uint32(len(data)*4), int32(bf), prototype)
	mem_size := core.vertex_buffers[name].reqs.Size
	min_align := core.vertex_buffers[name].reqs.Alignment

	if mem_ref, err := core.allocator.Allocate(int(mem_size), int(min_align)); err == nil {
		if err := core.allocator.Map(mdata, core.vertex_buffers[name].buffer[0], core.logical_device.handle, mem_ref, FLOAT32); err != nil {
			fmt.Errorf("Failed to bind buffer %s\n", name)
		}
	}
}

func (core *CoreRenderInstance) AddLayoutBuffer(data []float32, name string, usage vk.BufferUsageFlags) {
	d := Ptr(data)
	mdata := &d
	bf := vk.BufferUsageFlags(usage)
	core.uniform_buffers[name] = NewLayoutBuffer(core.logical_device.handle, core.logical_device.selected_device, uint32(len(data)*4), int32(bf))
	//Allocate the min alignment
	mem_size := core.uniform_buffers[name].reqs.Size
	min_align := core.uniform_buffers[name].reqs.Alignment

	if mem_ref, err := core.allocator.Allocate(int(mem_size), int(min_align)); err == nil {
		if err := core.allocator.Map(mdata, core.uniform_buffers[name].buffer[0], core.logical_device.handle, mem_ref, FLOAT32); err != nil {
			fmt.Errorf("Failed to bind buffer %s\n", name)
		}
	}
}

func (core *CoreRenderInstance) GetLayoutBuffers() map[string]*CoreBuffer {
	return core.uniform_buffers
}

//Active uniform buffers are bound to the descriptor sets and given
func (core *CoreRenderInstance) BindUniforms(uniforms map[string]*CoreBuffer, binding []int) error {

	for index := 0; index < 3; index++ {
		for _, uniform := range uniforms {
			core.frame_descriptor_sets.AddBuffer(core.logical_device.handle, binding[index], int(vk.DescriptorTypeUniformBuffer), index, *uniform)
		}
	}
	return nil
}

func (core *CoreRenderInstance) NewSwapchain() *CoreSwapchain {
	var err error
	core.swapchain = NewCoreSwapchain(core, SWAPCHAIN_COUNT, core.display)
	core.swapchain.init(core, core.swapchain.depth, core.display)
	core.per_frame = make([]PerFrame, core.swapchain.depth)
	for index := 0; index < core.swapchain.depth; index++ {
		core.per_frame[index], err = NewPerFrame(core, index)
	}
	if err != nil {
		Fatal(fmt.Errorf("Could not initiate per frame data\n"))
	}
	return core.swapchain
}

func (core *CoreRenderInstance) SetupCommands() {
	core.setup_commands()
}

func (core *CoreRenderInstance) AddRenderPass(name string) *CoreRenderPass {
	core.renderpasses[name] = NewCoreRenderPass()
	core.renderpasses[name].CreateRenderPass(core, core.display)
	core.swapchain.create_framebuffers(core, &core.renderpasses[name].renderPass)
	return core.renderpasses[name]
}

//Adds a pipline to this existing instance and builds a pipeline based on the given program identifier and a buffer which represents the expected vertex input for the pipeline
func (core *CoreRenderInstance) AddPipeline(name string, program_name string, buffer CoreBuffer, pass string) *CorePipeline {
	core.Builders[name] = NewPipelineBuilder(core, core.shaders.shader_programs[program_name], *buffer.prototype.GetInputDescription())
	core.pipeline.pipelines[name] = core.Builders[name].BuildPipeline(core, pass, core.display, core.pipeline.layouts[name])
	return core.pipeline
}

func (core *CoreRenderInstance) GetVertexBuffer(name string) *CoreBuffer {
	return core.vertex_buffers[name]
}

func (core *CoreRenderInstance) GetHandle() vk.Device {
	return core.logical_device.handle
}

func (core *CoreRenderInstance) GetPhysicalDevice() vk.PhysicalDevice {
	return core.logical_device.selected_device
}

func (core *CoreRenderInstance) destroy_per_frame() {

	//Destroying all per frame data - Warning Vulkan validation will throw an exception
	for index := 0; index < core.swapchain.depth; index++ {
		vk.ResetFences(core.logical_device.handle, uint32(1), core.per_frame[index].fence)
		vk.ResetCommandPool(core.logical_device.handle, core.per_frame[index].pool.pool, vk.CommandPoolResetFlags(vk.CommandPoolResetReleaseResourcesBit))
		vk.DestroySemaphore(core.logical_device.handle, core.per_frame[index].image_acquired[0], nil)
		vk.DestroySemaphore(core.logical_device.handle, core.per_frame[index].queue_complete[0], nil)

		fences := core.per_frame[index].fence
		for j := 0; j < len(fences); j++ {
			vk.DestroyFence(core.logical_device.handle, fences[j], nil)
		}

	}

	for index := 0; index < len(core.recycled_semaphores); index++ {
		vk.DestroySemaphore(core.logical_device.handle, core.recycled_semaphores[index], nil)
	}

	core.recycled_semaphores = make([]vk.Semaphore, 0)

}

func (core *CoreRenderInstance) destroy_swapchain() {
	core.destroy_per_frame()
	vk.DestroySwapchain(core.logical_device.handle, core.swapchain.swapchain, nil)
}

func (core *CoreRenderInstance) submit_pipeline(image uint32) vk.Result {

	//Pipleline stage flags
	waitDstStageMask := []vk.PipelineStageFlags{
		vk.PipelineStageFlags(vk.PipelineStageColorAttachmentOutputBit),
	}

	submitInfo := vk.SubmitInfo{
		SType:                vk.StructureTypeSubmitInfo,
		WaitSemaphoreCount:   1,
		PWaitSemaphores:      core.per_frame[core.current_frame].image_acquired,
		PWaitDstStageMask:    waitDstStageMask,
		CommandBufferCount:   1,
		PCommandBuffers:      core.per_frame[core.current_frame].command,
		SignalSemaphoreCount: 1,
		PSignalSemaphores:    core.per_frame[core.current_frame].queue_complete,
	}

	queue := core.render_queue

	res_queue := vk.QueueSubmit(*queue, 1, []vk.SubmitInfo{submitInfo}, core.per_frame[core.current_frame].fence[0])

	return res_queue
}

func (core *CoreRenderInstance) Update(delta_time float32) {
	image_index := uint32(0)

	res := core.acquire_next_image(&image_index)

	if res == vk.Suboptimal || res == vk.ErrorOutOfDate {
		core.resize()
		res = core.acquire_next_image(&image_index)
	}

	if res != vk.Success {
		vk.QueueWaitIdle(*core.render_queue)
	}

	core.setup_command(int(core.current_frame), image_index)

	core.submit_pipeline(image_index)

	res = core.present_image(*core.render_queue, image_index)

	if res == vk.ErrorOutOfDate || res == vk.Suboptimal {
		core.resize()
	} else if res != vk.Success {
		Fatal(fmt.Errorf("Failed to present swapchain image\n"))
	}

	core.current_frame = (core.current_frame + 1) % core.swapchain.depth
	core.pconstant[0].frame_delta += update_step
	if core.pconstant[0].frame_delta > 1.0 {
		update_step = -0.01
	}
	if core.pconstant[0].frame_delta < 0.0 {
		update_step = 0.01
	}

	return
}

func (core *CoreRenderInstance) present_image(queue vk.Queue, image_index uint32) vk.Result {

	present_info := vk.PresentInfo{}
	present_info.SType = vk.StructureTypePresentInfo
	present_info.WaitSemaphoreCount = 1
	present_info.PWaitSemaphores = []vk.Semaphore{core.per_frame[core.current_frame].queue_complete[0]}
	swaps := []vk.Swapchain{core.swapchain.swapchain}
	present_info.PSwapchains = swaps
	present_info.SwapchainCount = 1
	present_info.PImageIndices = []uint32{image_index}

	return vk.QueuePresent(queue, &present_info)

}

func (core *CoreRenderInstance) release() {
	core.Destroy()
	for _, buffer := range core.uniform_buffers {
		buffer.Destroy(core.logical_device.handle)
	}
}

func (core *CoreRenderInstance) Destroy() {
	//TODO destroy framebuffers // destroy semaphores // Destroy Pipline // Destroy Pipline Layout //Destroy

	vk.DeviceWaitIdle(core.logical_device.handle)

	core.swapchain.teardown_framebuffers(core)

	core.destroy_per_frame()

	for _, frame := range core.per_frame {
		vk.DestroyCommandPool(core.logical_device.handle, frame.pool.pool, nil)
	}
	for _, frame := range core.recycled_semaphores {
		vk.DestroySemaphore(core.logical_device.handle, frame, nil)

	}

	core.pipeline.destroy(core.logical_device.handle)

	for _, render := range core.renderpasses {
		if render.renderPass != vk.NullRenderPass {
			vk.DestroyRenderPass(core.logical_device.handle, render.renderPass, nil)
		}
	}

	for _, shader := range core.shaders.shader_programs {
		vk.DestroyShaderModule(core.logical_device.handle, *shader.vertex_shader_modules, nil)
		vk.DestroyShaderModule(core.logical_device.handle, *shader.fragment_shader_modules, nil)
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

	for _, buffer := range core.vertex_buffers {
		buffer.Destroy(core.logical_device.handle)
	}

	for index, view := range core.swapchain.image_views {
		if view != vk.NullImageView {
			vk.DestroyImageView(core.logical_device.handle, core.swapchain.image_views[index], nil)
		}
	}

	if core.swapchain.old_swapchain != vk.NullSwapchain {
		vk.DestroySwapchain(core.logical_device.handle, core.swapchain.old_swapchain, nil)
	}

	if core.swapchain.swapchain != vk.NullSwapchain {

		vk.DestroySwapchain(core.logical_device.handle, core.swapchain.swapchain, nil)
	}

	if core.display.surface != vk.NullSurface {
		vk.DestroySurface(*core.instance, core.display.surface, nil)
	}

	core.allocator.Destroy(core.logical_device.handle)

	vk.DestroyDevice(core.logical_device.handle, nil)
}

func (core *CoreRenderInstance) acquire_next_image(image *uint32) vk.Result {

	res := vk.AcquireNextImage(core.logical_device.handle, core.swapchain.swapchain, vk.MaxUint64,
		core.per_frame[core.current_frame].image_acquired[0], nil, image)

	if res != vk.Success {
		//	core.recycled_semaphores = append(core.recycled_semaphores, acquire_semaphore)
		return res
	}

	if core.per_frame[core.current_frame].fence[0] != vk.Fence(vk.NullHandle) {
		vk.WaitForFences(core.logical_device.handle, 1, core.per_frame[core.current_frame].fence, vk.True, vk.MaxUint64)
		vk.ResetFences(core.logical_device.handle, 1, core.per_frame[core.current_frame].fence)
	}

	if core.per_frame[core.current_frame].pool.pool != vk.CommandPool(vk.NullHandle) {
		vk.QueueWaitIdle(*core.render_queue)
		vk.ResetCommandPool(core.logical_device.handle, core.per_frame[core.current_frame].pool.pool, 0)
	}

	return vk.Success

}

func (core *CoreRenderInstance) setup_command(index int, image_index uint32) {

	clearValues := []vk.ClearValue{
		vk.NewClearValue([]float32{0.15, 0.15, 0.15, 1.0}),
		vk.NewClearDepthStencil(1.0, 0.0),
	}

	viewport := vk.Viewport{}
	scissor := vk.Rect2D{}
	viewport.Width = float32(core.swapchain.extent.Width)
	viewport.Height = float32(core.swapchain.extent.Height)
	scissor.Extent.Width = core.swapchain.extent.Width
	scissor.Extent.Height = core.swapchain.extent.Height

	viewports := []vk.Viewport{
		viewport,
	}

	rects := []vk.Rect2D{
		scissor,
	}

	cmd := core.per_frame[index].command
	vk.ResetCommandBuffer(cmd[0], vk.CommandBufferResetFlags(vk.CommandPoolResetReleaseResourcesBit))
	vk.BeginCommandBuffer(cmd[0], &vk.CommandBufferBeginInfo{
		SType: vk.StructureTypeCommandBufferBeginInfo,
		Flags: vk.CommandBufferUsageFlags(vk.CommandBufferUsageOneTimeSubmitBit),
	})

	vk.CmdBeginRenderPass(cmd[0], &vk.RenderPassBeginInfo{
		SType:           vk.StructureTypeRenderPassBeginInfo,
		RenderPass:      core.renderpasses["rp0"].renderPass,
		Framebuffer:     core.swapchain.framebuffers[image_index],
		RenderArea:      core.swapchain.rect,
		ClearValueCount: uint32(len(clearValues)),
		PClearValues:    clearValues,
	}, vk.SubpassContentsInline)

	vk.CmdBindPipeline(cmd[0], vk.PipelineBindPointGraphics, core.pipeline.pipelines["pipe0"])
	gather, _ := core.frame_descriptor_sets.GatherSets()
	frame_set := make([]vk.DescriptorSet, 1)
	frame_set[0] = gather[index]
	vk.CmdBindDescriptorSets(cmd[0], vk.PipelineBindPointGraphics, core.pipeline.layouts["pipe0"], 0, 1, frame_set, 0, nil)
	vk.CmdPushConstants(cmd[0], core.pipeline.layouts["pipe0"], vk.ShaderStageFlags(vk.ShaderStageVertexBit), 0, 4, unsafe.Pointer(&core.pconstant[0]))
	vk.CmdSetViewport(cmd[0], 0, 1, viewports)
	vk.CmdSetScissor(cmd[0], 0, 1, rects)
	tri_buffer := core.vertex_buffers["triangle"]
	offsets := []vk.DeviceSize{vk.DeviceSize(0)}
	vk.CmdBindVertexBuffers(cmd[0], 0, 1, tri_buffer.buffer, offsets)
	vk.CmdDraw(cmd[0], tri_buffer.groups, 1, 0, 0)

	vk.CmdEndRenderPass(cmd[0])
	vk.EndCommandBuffer(cmd[0])

}

func (core *CoreRenderInstance) setup_commands() {
	// Command Buffer Per Render-Pass per swapchain image which means they are interchangeable
	for i := 0; i < core.swapchain.depth; i++ {
		core.setup_command(i, uint32(i))
	}
}

func (core *CoreRenderInstance) is_valid_device(device *vk.PhysicalDevice, flags uint32) bool {

	q := NewCoreQueue(*device, "Default")
	return q.IsDeviceSuitable(flags)
}

func (core *CoreRenderInstance) resize() {
	var surface_capabilities vk.SurfaceCapabilities
	vk.GetPhysicalDeviceSurfaceCapabilities(core.logical_device.selected_device, core.display.surface, &surface_capabilities)
	surface_capabilities.Deref()

	if surface_capabilities.CurrentExtent.Width == core.swapchain.extent.Width && surface_capabilities.CurrentExtent.Height == core.swapchain.extent.Height {
		return
	}
	core.swapchain.old_swapchain = core.swapchain.swapchain
	vk.DestroySwapchain(core.logical_device.handle, core.swapchain.swapchain, nil)

	if len(core.swapchain.image_views) > 0 {
		for i := 0; i < len(core.swapchain.image_views); i++ {
			vk.DestroyImageView(core.logical_device.handle, core.swapchain.image_views[i], nil)
		}
	}

	core.swapchain.teardown_framebuffers(core)
	core.swapchain.init(core, core.swapchain.depth, core.display)
	vk.DeviceWaitIdle(core.logical_device.handle)
	core.swapchain.create_framebuffers(core, &core.renderpasses["rp0"].renderPass)

}

func (core *CoreRenderInstance) AllocatorUsage() {
	output := core.allocator.Usage()
	fmt.Printf(output)
}
