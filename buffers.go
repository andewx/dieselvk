package dieselvk

import (
	vk "github.com/vulkan-go/vulkan"
)

type CoreBuffer struct {
	buffer    []vk.Buffer
	mode      vk.SharingMode
	usage     int32
	reqs      vk.MemoryRequirements
	mem_index uint32
	elements  uint32
	groups    uint32
	prototype VertexAttribute
}

//Specifies new buffer memory allocation with a vertex attribute description attachment
func NewCoreVertexBuffer(handle vk.Device, physical vk.PhysicalDevice, bytes_size uint32, buffer_type int32, vertex VertexAttribute) *CoreBuffer {
	core := CoreBuffer{}
	core.buffer = make([]vk.Buffer, 1)
	core.usage = buffer_type
	dev_size := vk.DeviceSize(bytes_size)
	core.mode = vk.SharingMode(vk.SharingModeExclusive)
	core.prototype = vertex

	buffer_create := vk.BufferCreateInfo{}
	buffer_create.SType = vk.StructureTypeBufferCreateInfo
	buffer_create.Flags = vk.BufferCreateFlags(0)
	buffer_create.Usage = vk.BufferUsageFlags(buffer_type)
	buffer_create.SharingMode = core.mode

	buffer_create.Size = dev_size
	core.elements = uint32(dev_size / 4)
	core.groups = core.elements / 3

	res := vk.CreateBuffer(handle, &buffer_create, nil, &core.buffer[0])

	if res != vk.Success {
		Fatal(NewError(res))
	}

	vk.GetBufferMemoryRequirements(handle, core.buffer[0], &core.reqs)
	core.reqs.Deref()

	return &core

}

//Specifies new buffer memory allocation with a vertex attribute description attachment
func NewLayoutBuffer(handle vk.Device, physical vk.PhysicalDevice, bytes_size uint32, buffer_type int32) *CoreBuffer {
	core := CoreBuffer{}
	core.buffer = make([]vk.Buffer, 1)
	core.usage = buffer_type
	dev_size := vk.DeviceSize(bytes_size)
	core.mode = vk.SharingMode(vk.SharingModeExclusive)
	core.prototype = nil

	buffer_create := vk.BufferCreateInfo{}
	buffer_create.SType = vk.StructureTypeBufferCreateInfo
	buffer_create.Flags = vk.BufferCreateFlags(0)
	buffer_create.Usage = vk.BufferUsageFlags(buffer_type)
	buffer_create.SharingMode = core.mode
	buffer_create.Size = dev_size
	core.elements = uint32(dev_size / 4)
	core.groups = core.elements

	res := vk.CreateBuffer(handle, &buffer_create, nil, &core.buffer[0])

	if res != vk.Success {
		Fatal(NewError(res))
	}

	vk.GetBufferMemoryRequirements(handle, core.buffer[0], &core.reqs)
	core.reqs.Deref()

	return &core

}

func (core *CoreBuffer) Destroy(handle vk.Device) {
	vk.DestroyBuffer(handle, core.buffer[0], nil)

}
