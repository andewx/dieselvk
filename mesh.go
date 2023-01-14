package dieselvk

import (
	"unsafe"

	vk "github.com/vulkan-go/vulkan"
)

/*
 Per Vertex Attribute Shader Layout Attribute Locations are fixed given the layouts
 in this document. Vertex Attribute layouts are given by the VertexAttribute interface
 function GetInputDescription from a prototype (or filled) Vertex object
*/
type VertexInputDescription struct {
	bindings   []vk.VertexInputBindingDescription
	attributes []vk.VertexInputAttributeDescription
	flags      vk.PipelineVertexInputStateCreateFlags
}

type VertexAttribute interface {
	GetInputDescription() *VertexInputDescription
}

type Vertex struct {
	position [3]float32
}

func (v Vertex) GetInputDescription() *VertexInputDescription {

	vertex := VertexInputDescription{}
	vertex.bindings = make([]vk.VertexInputBindingDescription, 1)
	vertex.attributes = make([]vk.VertexInputAttributeDescription, 1)
	vertex.flags = vk.PipelineVertexInputStateCreateFlags(0)

	//Position binding
	binding := vk.VertexInputBindingDescription{}
	binding.Binding = 0
	binding.Stride = uint32(unsafe.Sizeof(v))
	binding.InputRate = vk.VertexInputRateVertex
	binding.InputRate = vk.VertexInputRateInstance

	//Attribute Location 0
	p_attr := vk.VertexInputAttributeDescription{}
	p_attr.Binding = 0
	p_attr.Location = 0
	p_attr.Format = vk.FormatR32g32b32Sfloat

	//Set Object
	vertex.bindings[0] = binding
	vertex.attributes[0] = p_attr
	return &vertex
}
