package json

import (
	"bytes"
	"encoding/json"
	"errors"
)

// AttachmentBlend
type AttachmentBlend struct {
	BlendMode string `json:"blend_mode"`
	WriteMode string `json:"write_mode"`
}

// Buffer
type Buffer struct {
	ElementPaddedSize int    `json:"element_padded_size,omitempty"`
	ElementSize       int    `json:"element_size,omitempty"`
	Elements          int    `json:"elements,omitempty"`
	Format            string `json:"format,,omitempty"`
	GpuAllocationSize int    `json:"gpu_allocation_size,omitempty"`
	Name              string `json:"name"`
	Offset            int    `json:"offset"`
	Size              int    `json:"size"`
}

// CommandBuffer
type CommandBuffer struct {
	Count int    `json:"count"`
	Level string `json:"level"`
	Name  string `json:"name"`
	Queue string `json:"queue"`
}

// Config
type Config struct {
	CoreExtensions  []string    `json:"core_extensions"`
	DeviceMode      string      `json:"device_mode"`
	InstanceName    string      `json:"instance_name"`
	InstanceVersion string      `json:"instance_version"`
	SwapchainSize   int         `json:"swapchain_size"`
	Swapchains      []Swapchain `json:"swapchains"`
	UserExtensions  []string    `json:"user_extensions"`
	VulkanLayers    []string    `json:"vulkan_layers"`
	Display         string      `json:"display"`
	Window          *Window     `json:"window"`
}

// Depth
type Depth struct {
	BiasClamp    float64 `json:"bias_clamp"`
	BiasConstant float64 `json:"bias_constant"`
	BiasSlope    float64 `json:"bias_slope"`
}

// DescriptorLayout
type DescriptorLayout struct {
	Binding       int      `json:"binding"`
	Name          string   `json:"name"`
	Sets          int      `json:"sets"`
	ShaderStages  []string `json:"shader_stages"`
	Type          string   `json:"type"`
	UniformBuffer string   `json:"uniform_buffer"`
}

// Extent
type Extent struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

// Framebuffer
type Framebuffer struct {
	Attachments []ImageAttachment `json:"attachments"`
	ImageViews  []Image           `json:"image_views"`
}

// Image
type Image struct {
	AspectFlags string  `json:"aspect_flags"`
	BaseArray   int     `json:"base_array"`
	BaseMip     int     `json:"base_mip"`
	Extent      *Extent `json:"extent"`
	Format      string  `json:"format"`
	IsPerFrame  bool    `json:"is_per_frame"`
	LayerCount  int     `json:"layer_count"`
	LevelCount  int     `json:"level_count"`
	Name        string  `json:"name"`
	Sampling    string  `json:"sampling"`
	Tiling      string  `json:"tiling"`
	ViewType    string  `json:"view_type"`
}

// ImageAttachment
type ImageAttachment struct {
	FinalLayout    string `json:"final_layout"`
	Format         string `json:"format"`
	ImageRef       string `json:"image_ref"`
	InitialLayout  string `json:"initial_layout"`
	Layout         string `json:"layout"`
	LoadOp         string `json:"load_op"`
	Name           string `json:"name"`
	SampleCount    string `json:"sample_count"`
	StencilLoadOp  string `json:"stencil_load_op"`
	StencilStoreOp string `json:"stencil_store_op"`
	StoreOp        string `json:"store_op"`
}

// MeshDescriptor
type MeshDescriptor struct {
	Description   string         `json:"description"`
	IndicesBuffer bool           `json:"indices_buffer"`
	Name          string         `json:"name"`
	VertexBinding *VertexBinding `json:"vertex_binding"`
}

// Mssa
type Mssa struct {
	AlphaEnable      bool    `json:"alpha_enable"`
	AlphaOne         bool    `json:"alpha_one"`
	BitSample        string  `json:"bit_sample"`
	Enable           bool    `json:"enable"`
	MinSampleShading float64 `json:"min_sample_shading"`
}

// Pipeline
type Pipeline struct {
	AttachmentBlend    *AttachmentBlend    `json:"attachment_blend"`
	Constants          []PushConstant      `json:"constants"`
	CullMode           string              `json:"cull_mode"`
	Depth              *Depth              `json:"depth"`
	FrontFace          string              `json:"front_face"`
	Layouts            []DescriptorLayout  `json:"layouts"`
	LineWidth          float64             `json:"line_width"`
	MeshDescriptor     *MeshDescriptor     `json:"mesh_descriptor"`
	Mssa               *Mssa               `json:"mssa"`
	Name               string              `json:"name"`
	PipelineAttributes []PipelineAttribute `json:"pipeline_attributes"`
	Program            string              `json:"program"`
	Renderpass         string              `json:"renderpass"`
	Subpass            int                 `json:"subpass"`
	Topology           string              `json:"topology"`
}

// PipelineAttribute
type PipelineAttribute struct {
	Binding  int    `json:"binding"`
	Format   string `json:"format"`
	Location int    `json:"location"`
	Name     string `json:"name"`
	Offset   int    `json:"offset"`
}

// PushConstant
type PushConstant struct {
	Name         string   `json:"name"`
	Offset       int      `json:"offset"`
	ShaderStages []string `json:"shader_stages"`
	Size         int      `json:"size"`
}

// Queue
type Queue struct {
	Family string `json:"family"`
	Index  int    `json:"index"`
	Name   string `json:"name"`
}

// Renderpass
type Renderpass struct {
	CommandBuffer string    `json:"command_buffer"`
	Name          string    `json:"name"`
	Subpasses     []Subpass `json:"subpasses"`
}

// Shader
type Shader struct {
	Name    string   `json:"name"`
	Sources []Source `json:"sources"`
}

// Source
type Source struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// Subpass
type Subpass struct {
	Bindpoint           string              `json:"bindpoint"`
	ColorAttachments    []string            `json:"color_attachments"`
	Dependencies        []SubpassDependency `json:"dependencies"`
	DepthAttachments    []string            `json:"depth_attachments"`
	InputAttachments    []string            `json:"input_attachments"`
	PreserveAttachments []string            `json:"preserve_attachments"`
}

// SubpassDependency
type SubpassDependency struct {
	DstAccessMask string `json:"dst_access_mask"`
	DstMask       string `json:"dst_mask"`
	DstSubpass    int    `json:"dst_subpass"`
	Name          string `json:"name"`
	SrcAccessMask int    `json:"src_access_mask"`
	SrcMask       string `json:"src_mask"`
	SrcSubpass    string `json:"src_subpass"`
}

// Swapchain
type Swapchain struct {
	ColorAttachment string `json:"color_attachment"`
	DepthAttachment string `json:"depth_attachment"`
}

// VertexBinding
type VertexBinding struct {
	Binding     int    `json:"binding"`
	Rate        string `json:"rate"`
	StrideBytes int    `json:"stride_bytes"`
}

// Vlk
type Vlk struct {
	Buffer            *Buffer            `json:"buffer,omitempty"`
	CommandBuffer     *CommandBuffer     `json:"command_buffer,omitempty"`
	CommandBuffers    []CommandBuffer    `json:"command_buffers"`
	Config            *Config            `json:"config"`
	DescriptorLayout  *DescriptorLayout  `json:"descriptor_layout,omitempty"`
	Framebuffer       *Framebuffer       `json:"framebuffer,omitempty"`
	Framebuffers      []Framebuffer      `json:"framebuffers"`
	Image             *Image             `json:"image,omitempty"`
	ImageAttachment   *ImageAttachment   `json:"image_attachment,omitempty"`
	IndiceBuffers     []Buffer           `json:"indice_buffers"`
	Pipeline          *Pipeline          `json:"pipeline,omitempty"`
	PipelineAttribute *PipelineAttribute `json:"pipeline_attribute,omitempty"`
	Pipelines         []Pipeline         `json:"pipelines"`
	PushConstant      *PushConstant      `json:"push_constant,omitempty"`
	Queue             *Queue             `json:"queue,omitempty"`
	Queues            []Queue            `json:"queues"`
	Renderpass        *Renderpass        `json:"renderpass,omitempty"`
	Renderpasses      []Renderpass       `json:"renderpasses"`
	SamplerBuffers    []Buffer           `json:"sampler_buffers"`
	Schema            string             `json:"schema"`
	Shader            *Shader            `json:"shader,omitempty"`
	ShaderPrograms    []Shader           `json:"shader_programs"`
	Source            *Source            `json:"source,omitempty"`
	Subpass           *Subpass           `json:"subpass,omitempty"`
	SubpassDependency *SubpassDependency `json:"subpass_dependency,omitempty"`
	Swapchain         *Swapchain         `json:"swapchain,omitempty"`
	UniformBuffers    []Buffer           `json:"uniform_buffers"`
	VertexBuffers     []Buffer           `json:"vertex_buffers"`
}

// Window
type Window struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

func (strct *AttachmentBlend) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "BlendMode" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "blend_mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"blend_mode\": ")
	if tmp, err := json.Marshal(strct.BlendMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "WriteMode" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "write_mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"write_mode\": ")
	if tmp, err := json.Marshal(strct.WriteMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AttachmentBlend) UnmarshalJSON(b []byte) error {
	blend_modeReceived := false
	write_modeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "blend_mode":
			if err := json.Unmarshal([]byte(v), &strct.BlendMode); err != nil {
				return err
			}
			blend_modeReceived = true
		case "write_mode":
			if err := json.Unmarshal([]byte(v), &strct.WriteMode); err != nil {
				return err
			}
			write_modeReceived = true
		}
	}
	// check if blend_mode (a required property) was received
	if !blend_modeReceived {
		return errors.New("\"blend_mode\" is required but was not present")
	}
	// check if write_mode (a required property) was received
	if !write_modeReceived {
		return errors.New("\"write_mode\" is required but was not present")
	}
	return nil
}

func (strct *Buffer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Offset" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "offset" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"offset\": ")
	if tmp, err := json.Marshal(strct.Offset); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Size" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "size" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"size\": ")
	if tmp, err := json.Marshal(strct.Size); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Buffer) UnmarshalJSON(b []byte) error {

	nameReceived := false
	offsetReceived := false
	sizeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "element_padded_size":
			if err := json.Unmarshal([]byte(v), &strct.ElementPaddedSize); err != nil {
				return err
			}

		case "element_size":
			if err := json.Unmarshal([]byte(v), &strct.ElementSize); err != nil {
				return err
			}

		case "elements":
			if err := json.Unmarshal([]byte(v), &strct.Elements); err != nil {
				return err
			}

		case "format":
			if err := json.Unmarshal([]byte(v), &strct.Format); err != nil {
				return err
			}

		case "gpu_allocation_size":
			if err := json.Unmarshal([]byte(v), &strct.GpuAllocationSize); err != nil {
				return err
			}

		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "offset":
			if err := json.Unmarshal([]byte(v), &strct.Offset); err != nil {
				return err
			}
			offsetReceived = true
		case "size":
			if err := json.Unmarshal([]byte(v), &strct.Size); err != nil {
				return err
			}
			sizeReceived = true
		}
	}

	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("Buffer \"name\" is required but was not present")
	}
	// check if offset (a required property) was received
	if !offsetReceived {
		return errors.New("Buffer \"offset\" is required but was not present")
	}
	// check if size (a required property) was received
	if !sizeReceived {
		return errors.New("Buffer \"size\" is required but was not present")
	}
	return nil
}

func (strct *CommandBuffer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Count" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "count" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"count\": ")
	if tmp, err := json.Marshal(strct.Count); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Level" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "level" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"level\": ")
	if tmp, err := json.Marshal(strct.Level); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Queue" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "queue" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"queue\": ")
	if tmp, err := json.Marshal(strct.Queue); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CommandBuffer) UnmarshalJSON(b []byte) error {
	countReceived := false
	levelReceived := false
	nameReceived := false
	queueReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "count":
			if err := json.Unmarshal([]byte(v), &strct.Count); err != nil {
				return err
			}
			countReceived = true
		case "level":
			if err := json.Unmarshal([]byte(v), &strct.Level); err != nil {
				return err
			}
			levelReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "queue":
			if err := json.Unmarshal([]byte(v), &strct.Queue); err != nil {
				return err
			}
			queueReceived = true
		}
	}
	// check if count (a required property) was received
	if !countReceived {
		return errors.New("CommandBuffer \"count\" is required but was not present")
	}
	// check if level (a required property) was received
	if !levelReceived {
		return errors.New("CommandBuffer  \"level\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("CommandBuffer  \"name\" is required but was not present")
	}
	// check if queue (a required property) was received
	if !queueReceived {
		return errors.New("CommandBuffer  \"queue\" is required but was not present")
	}
	return nil
}

func (strct *Config) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "CoreExtensions" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "core_extensions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"core_extensions\": ")
	if tmp, err := json.Marshal(strct.CoreExtensions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "DeviceMode" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "device_mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"device_mode\": ")
	if tmp, err := json.Marshal(strct.DeviceMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "InstanceName" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "instance_name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"instance_name\": ")
	if tmp, err := json.Marshal(strct.InstanceName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "InstanceVersion" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "instance_version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"instance_version\": ")
	if tmp, err := json.Marshal(strct.InstanceVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "SwapchainSize" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "swapchain_size" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"swapchain_size\": ")
	if tmp, err := json.Marshal(strct.SwapchainSize); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Swapchains" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "swapchains" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"swapchains\": ")
	if tmp, err := json.Marshal(strct.Swapchains); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "UserExtensions" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "user_extensions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"user_extensions\": ")
	if tmp, err := json.Marshal(strct.UserExtensions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "VulkanLayers" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "vulkan_layers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"vulkan_layers\": ")
	if tmp, err := json.Marshal(strct.VulkanLayers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Window" field is required
	if strct.Window == nil {
		return nil, errors.New("window is a required field")
	}
	// Marshal the "window" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"window\": ")
	if tmp, err := json.Marshal(strct.Window); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Config) UnmarshalJSON(b []byte) error {
	core_extensionsReceived := false
	device_modeReceived := false
	instance_nameReceived := false
	instance_versionReceived := false
	swapchain_sizeReceived := false
	swapchainsReceived := false
	user_extensionsReceived := false
	vulkan_layersReceived := false
	windowReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "core_extensions":
			if err := json.Unmarshal([]byte(v), &strct.CoreExtensions); err != nil {
				return err
			}
			core_extensionsReceived = true
		case "device_mode":
			if err := json.Unmarshal([]byte(v), &strct.DeviceMode); err != nil {
				return err
			}
			device_modeReceived = true
		case "instance_name":
			if err := json.Unmarshal([]byte(v), &strct.InstanceName); err != nil {
				return err
			}
			instance_nameReceived = true
		case "instance_version":
			if err := json.Unmarshal([]byte(v), &strct.InstanceVersion); err != nil {
				return err
			}
			instance_versionReceived = true
		case "swapchain_size":
			if err := json.Unmarshal([]byte(v), &strct.SwapchainSize); err != nil {
				return err
			}
			swapchain_sizeReceived = true
		case "swapchains":
			if err := json.Unmarshal([]byte(v), &strct.Swapchains); err != nil {
				return err
			}
			swapchainsReceived = true
		case "user_extensions":
			if err := json.Unmarshal([]byte(v), &strct.UserExtensions); err != nil {
				return err
			}
			user_extensionsReceived = true
		case "vulkan_layers":
			if err := json.Unmarshal([]byte(v), &strct.VulkanLayers); err != nil {
				return err
			}
			vulkan_layersReceived = true
		case "window":
			if err := json.Unmarshal([]byte(v), &strct.Window); err != nil {
				return err
			}
		case "display":
			if err := json.Unmarshal([]byte(v), &strct.Display); err != nil {
				return err
			}
			windowReceived = true
		}
	}
	// check if core_extensions (a required property) was received
	if !core_extensionsReceived {
		return errors.New("Config \"core_extensions\" is required but was not present")
	}
	// check if device_mode (a required property) was received
	if !device_modeReceived {
		return errors.New("Config  \"device_mode\" is required but was not present")
	}
	// check if instance_name (a required property) was received
	if !instance_nameReceived {
		return errors.New("Config  \"instance_name\" is required but was not present")
	}
	// check if instance_version (a required property) was received
	if !instance_versionReceived {
		return errors.New("Config  \"instance_version\" is required but was not present")
	}
	// check if swapchain_size (a required property) was received
	if !swapchain_sizeReceived {
		return errors.New("Config  \"swapchain_size\" is required but was not present")
	}
	// check if swapchains (a required property) was received
	if !swapchainsReceived {
		return errors.New("Config \"swapchains\" is required but was not present")
	}
	// check if user_extensions (a required property) was received
	if !user_extensionsReceived {
		return errors.New("Config \"user_extensions\" is required but was not present")
	}
	// check if vulkan_layers (a required property) was received
	if !vulkan_layersReceived {
		return errors.New("Config \"vulkan_layers\" is required but was not present")
	}
	// check if window (a required property) was received
	if !windowReceived {
		return errors.New("Config \"window\" is required but was not present")
	}
	return nil
}

func (strct *Depth) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "BiasClamp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "bias_clamp" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"bias_clamp\": ")
	if tmp, err := json.Marshal(strct.BiasClamp); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "BiasConstant" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "bias_constant" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"bias_constant\": ")
	if tmp, err := json.Marshal(strct.BiasConstant); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "BiasSlope" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "bias_slope" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"bias_slope\": ")
	if tmp, err := json.Marshal(strct.BiasSlope); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Depth) UnmarshalJSON(b []byte) error {
	bias_clampReceived := false
	bias_constantReceived := false
	bias_slopeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "bias_clamp":
			if err := json.Unmarshal([]byte(v), &strct.BiasClamp); err != nil {
				return err
			}
			bias_clampReceived = true
		case "bias_constant":
			if err := json.Unmarshal([]byte(v), &strct.BiasConstant); err != nil {
				return err
			}
			bias_constantReceived = true
		case "bias_slope":
			if err := json.Unmarshal([]byte(v), &strct.BiasSlope); err != nil {
				return err
			}
			bias_slopeReceived = true
		}
	}
	// check if bias_clamp (a required property) was received
	if !bias_clampReceived {
		return errors.New("Depth \"bias_clamp\" is required but was not present")
	}
	// check if bias_constant (a required property) was received
	if !bias_constantReceived {
		return errors.New("Depth \"bias_constant\" is required but was not present")
	}
	// check if bias_slope (a required property) was received
	if !bias_slopeReceived {
		return errors.New("Depth \"bias_slope\" is required but was not present")
	}
	return nil
}

func (strct *DescriptorLayout) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Binding" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "binding" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("DescriptorLayout \"binding\": ")
	if tmp, err := json.Marshal(strct.Binding); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("DescriptorLayout \"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Sets" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sets" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("DescriptorLayout \"sets\": ")
	if tmp, err := json.Marshal(strct.Sets); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ShaderStages" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "shader_stages" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("DescriptorLayout \"shader_stages\": ")
	if tmp, err := json.Marshal(strct.ShaderStages); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Type" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "type" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("DescriptorLayout \"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "UniformBuffer" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "uniform_buffer" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("DescriptorLayout \"uniform_buffer\": ")
	if tmp, err := json.Marshal(strct.UniformBuffer); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DescriptorLayout) UnmarshalJSON(b []byte) error {
	bindingReceived := false
	nameReceived := false
	setsReceived := false
	shader_stagesReceived := false
	typeReceived := false
	uniform_bufferReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "binding":
			if err := json.Unmarshal([]byte(v), &strct.Binding); err != nil {
				return err
			}
			bindingReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "sets":
			if err := json.Unmarshal([]byte(v), &strct.Sets); err != nil {
				return err
			}
			setsReceived = true
		case "shader_stages":
			if err := json.Unmarshal([]byte(v), &strct.ShaderStages); err != nil {
				return err
			}
			shader_stagesReceived = true
		case "type":
			if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
				return err
			}
			typeReceived = true
		case "uniform_buffer":
			if err := json.Unmarshal([]byte(v), &strct.UniformBuffer); err != nil {
				return err
			}
			uniform_bufferReceived = true
		}
	}
	// check if binding (a required property) was received
	if !bindingReceived {
		return errors.New("DescriptorLayout \"binding\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("DescriptorLayout \"name\" is required but was not present")
	}
	// check if sets (a required property) was received
	if !setsReceived {
		return errors.New("DescriptorLayout \"sets\" is required but was not present")
	}
	// check if shader_stages (a required property) was received
	if !shader_stagesReceived {
		return errors.New("DescriptorLayout \"shader_stages\" is required but was not present")
	}
	// check if type (a required property) was received
	if !typeReceived {
		return errors.New("DescriptorLayout \"type\" is required but was not present")
	}
	// check if uniform_buffer (a required property) was received
	if !uniform_bufferReceived {
		return errors.New("DescriptorLayout \"uniform_buffer\" is required but was not present")
	}
	return nil
}

func (strct *Extent) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Height" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "height" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Extent \"height\": ")
	if tmp, err := json.Marshal(strct.Height); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Width" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "width" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Extent \"width\": ")
	if tmp, err := json.Marshal(strct.Width); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Extent) UnmarshalJSON(b []byte) error {
	heightReceived := false
	widthReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "height":
			if err := json.Unmarshal([]byte(v), &strct.Height); err != nil {
				return err
			}
			heightReceived = true
		case "width":
			if err := json.Unmarshal([]byte(v), &strct.Width); err != nil {
				return err
			}
			widthReceived = true
		}
	}
	// check if height (a required property) was received
	if !heightReceived {
		return errors.New("Extent \"height\" is required but was not present")
	}
	// check if width (a required property) was received
	if !widthReceived {
		return errors.New("Extent \"width\" is required but was not present")
	}
	return nil
}

func (strct *Framebuffer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Attachments" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "attachments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Framebuffer \"attachments\": ")
	if tmp, err := json.Marshal(strct.Attachments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ImageViews" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "image_views" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Framebuffer \"image_views\": ")
	if tmp, err := json.Marshal(strct.ImageViews); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Framebuffer) UnmarshalJSON(b []byte) error {
	attachmentsReceived := false
	image_viewsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "attachments":
			if err := json.Unmarshal([]byte(v), &strct.Attachments); err != nil {
				return err
			}
			attachmentsReceived = true
		case "image_views":
			if err := json.Unmarshal([]byte(v), &strct.ImageViews); err != nil {
				return err
			}
			image_viewsReceived = true
		}
	}
	// check if attachments (a required property) was received
	if !attachmentsReceived {
		return errors.New("Framebuffer \"attachments\" is required but was not present")
	}
	// check if image_views (a required property) was received
	if !image_viewsReceived {
		return errors.New("Framebuffer \"image_views\" is required but was not present")
	}
	return nil
}

func (strct *Image) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "AspectFlags" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "aspect_flags" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"aspect_flags\": ")
	if tmp, err := json.Marshal(strct.AspectFlags); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "BaseArray" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "base_array" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"base_array\": ")
	if tmp, err := json.Marshal(strct.BaseArray); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "BaseMip" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "base_mip" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"base_mip\": ")
	if tmp, err := json.Marshal(strct.BaseMip); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Extent" field is required
	if strct.Extent == nil {
		return nil, errors.New("extent is a required field")
	}
	// Marshal the "extent" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"extent\": ")
	if tmp, err := json.Marshal(strct.Extent); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Format" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "format" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"format\": ")
	if tmp, err := json.Marshal(strct.Format); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IsPerFrame" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "is_per_frame" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"is_per_frame\": ")
	if tmp, err := json.Marshal(strct.IsPerFrame); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "LayerCount" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "layer_count" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"layer_count\": ")
	if tmp, err := json.Marshal(strct.LayerCount); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "LevelCount" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "level_count" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"level_count\": ")
	if tmp, err := json.Marshal(strct.LevelCount); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Sampling" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sampling" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sampling\": ")
	if tmp, err := json.Marshal(strct.Sampling); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Tiling" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "tiling" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tiling\": ")
	if tmp, err := json.Marshal(strct.Tiling); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ViewType" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "view_type" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"view_type\": ")
	if tmp, err := json.Marshal(strct.ViewType); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Image) UnmarshalJSON(b []byte) error {
	aspect_flagsReceived := false
	base_arrayReceived := false
	base_mipReceived := false
	extentReceived := false
	formatReceived := false
	is_per_frameReceived := false
	layer_countReceived := false
	level_countReceived := false
	nameReceived := false
	samplingReceived := false
	tilingReceived := false
	view_typeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "aspect_flags":
			if err := json.Unmarshal([]byte(v), &strct.AspectFlags); err != nil {
				return err
			}
			aspect_flagsReceived = true
		case "base_array":
			if err := json.Unmarshal([]byte(v), &strct.BaseArray); err != nil {
				return err
			}
			base_arrayReceived = true
		case "base_mip":
			if err := json.Unmarshal([]byte(v), &strct.BaseMip); err != nil {
				return err
			}
			base_mipReceived = true
		case "extent":
			if err := json.Unmarshal([]byte(v), &strct.Extent); err != nil {
				return err
			}
			extentReceived = true
		case "format":
			if err := json.Unmarshal([]byte(v), &strct.Format); err != nil {
				return err
			}
			formatReceived = true
		case "is_per_frame":
			if err := json.Unmarshal([]byte(v), &strct.IsPerFrame); err != nil {
				return err
			}
			is_per_frameReceived = true
		case "layer_count":
			if err := json.Unmarshal([]byte(v), &strct.LayerCount); err != nil {
				return err
			}
			layer_countReceived = true
		case "level_count":
			if err := json.Unmarshal([]byte(v), &strct.LevelCount); err != nil {
				return err
			}
			level_countReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "sampling":
			if err := json.Unmarshal([]byte(v), &strct.Sampling); err != nil {
				return err
			}
			samplingReceived = true
		case "tiling":
			if err := json.Unmarshal([]byte(v), &strct.Tiling); err != nil {
				return err
			}
			tilingReceived = true
		case "view_type":
			if err := json.Unmarshal([]byte(v), &strct.ViewType); err != nil {
				return err
			}
			view_typeReceived = true
		}
	}
	// check if aspect_flags (a required property) was received
	if !aspect_flagsReceived {
		return errors.New("Image \"aspect_flags\" is required but was not present")
	}
	// check if base_array (a required property) was received
	if !base_arrayReceived {
		return errors.New("Image \"base_array\" is required but was not present")
	}
	// check if base_mip (a required property) was received
	if !base_mipReceived {
		return errors.New("Image \"base_mip\" is required but was not present")
	}
	// check if extent (a required property) was received
	if !extentReceived {
		return errors.New("Image \"extent\" is required but was not present")
	}
	// check if format (a required property) was received
	if !formatReceived {
		return errors.New("Image \"format\" is required but was not present")
	}
	// check if is_per_frame (a required property) was received
	if !is_per_frameReceived {
		return errors.New("Image \"is_per_frame\" is required but was not present")
	}
	// check if layer_count (a required property) was received
	if !layer_countReceived {
		return errors.New("Image \"layer_count\" is required but was not present")
	}
	// check if level_count (a required property) was received
	if !level_countReceived {
		return errors.New("Image \"level_count\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("Image \"name\" is required but was not present")
	}
	// check if sampling (a required property) was received
	if !samplingReceived {
		return errors.New("Image \"sampling\" is required but was not present")
	}
	// check if tiling (a required property) was received
	if !tilingReceived {
		return errors.New("Image \"tiling\" is required but was not present")
	}
	// check if view_type (a required property) was received
	if !view_typeReceived {
		return errors.New("Image \"view_type\" is required but was not present")
	}
	return nil
}

func (strct *ImageAttachment) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "FinalLayout" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "final_layout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"final_layout\": ")
	if tmp, err := json.Marshal(strct.FinalLayout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Format" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "format" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"format\": ")
	if tmp, err := json.Marshal(strct.Format); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ImageRef" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "image_ref" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"image_ref\": ")
	if tmp, err := json.Marshal(strct.ImageRef); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "InitialLayout" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "initial_layout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"initial_layout\": ")
	if tmp, err := json.Marshal(strct.InitialLayout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Layout" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "layout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"layout\": ")
	if tmp, err := json.Marshal(strct.Layout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "LoadOp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "load_op" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"load_op\": ")
	if tmp, err := json.Marshal(strct.LoadOp); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "SampleCount" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sample_count" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"sample_count\": ")
	if tmp, err := json.Marshal(strct.SampleCount); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StencilLoadOp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "stencil_load_op" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"stencil_load_op\": ")
	if tmp, err := json.Marshal(strct.StencilLoadOp); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StencilStoreOp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "stencil_store_op" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"stencil_store_op\": ")
	if tmp, err := json.Marshal(strct.StencilStoreOp); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StoreOp" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "store_op" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("ImageAttachment \"store_op\": ")
	if tmp, err := json.Marshal(strct.StoreOp); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ImageAttachment) UnmarshalJSON(b []byte) error {
	final_layoutReceived := false
	formatReceived := false
	image_refReceived := false
	initial_layoutReceived := false
	layoutReceived := false
	load_opReceived := false
	nameReceived := false
	sample_countReceived := false
	stencil_load_opReceived := false
	stencil_store_opReceived := false
	store_opReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "final_layout":
			if err := json.Unmarshal([]byte(v), &strct.FinalLayout); err != nil {
				return err
			}
			final_layoutReceived = true
		case "format":
			if err := json.Unmarshal([]byte(v), &strct.Format); err != nil {
				return err
			}
			formatReceived = true
		case "image_ref":
			if err := json.Unmarshal([]byte(v), &strct.ImageRef); err != nil {
				return err
			}
			image_refReceived = true
		case "initial_layout":
			if err := json.Unmarshal([]byte(v), &strct.InitialLayout); err != nil {
				return err
			}
			initial_layoutReceived = true
		case "layout":
			if err := json.Unmarshal([]byte(v), &strct.Layout); err != nil {
				return err
			}
			layoutReceived = true
		case "load_op":
			if err := json.Unmarshal([]byte(v), &strct.LoadOp); err != nil {
				return err
			}
			load_opReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "sample_count":
			if err := json.Unmarshal([]byte(v), &strct.SampleCount); err != nil {
				return err
			}
			sample_countReceived = true
		case "stencil_load_op":
			if err := json.Unmarshal([]byte(v), &strct.StencilLoadOp); err != nil {
				return err
			}
			stencil_load_opReceived = true
		case "stencil_store_op":
			if err := json.Unmarshal([]byte(v), &strct.StencilStoreOp); err != nil {
				return err
			}
			stencil_store_opReceived = true
		case "store_op":
			if err := json.Unmarshal([]byte(v), &strct.StoreOp); err != nil {
				return err
			}
			store_opReceived = true
		}
	}
	// check if final_layout (a required property) was received
	if !final_layoutReceived {
		return errors.New("ImageAttachment \"final_layout\" is required but was not present")
	}
	// check if format (a required property) was received
	if !formatReceived {
		return errors.New("ImageAttachment \"format\" is required but was not present")
	}
	// check if image_ref (a required property) was received
	if !image_refReceived {
		return errors.New("ImageAttachment \"image_ref\" is required but was not present")
	}
	// check if initial_layout (a required property) was received
	if !initial_layoutReceived {
		return errors.New("ImageAttachment \"initial_layout\" is required but was not present")
	}
	// check if layout (a required property) was received
	if !layoutReceived {
		return errors.New("ImageAttachment \"layout\" is required but was not present")
	}
	// check if load_op (a required property) was received
	if !load_opReceived {
		return errors.New("ImageAttachment \"load_op\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("ImageAttachment \"name\" is required but was not present")
	}
	// check if sample_count (a required property) was received
	if !sample_countReceived {
		return errors.New("ImageAttachment \"sample_count\" is required but was not present")
	}
	// check if stencil_load_op (a required property) was received
	if !stencil_load_opReceived {
		return errors.New("ImageAttachment \"stencil_load_op\" is required but was not present")
	}
	// check if stencil_store_op (a required property) was received
	if !stencil_store_opReceived {
		return errors.New("ImageAttachment \"stencil_store_op\" is required but was not present")
	}
	// check if store_op (a required property) was received
	if !store_opReceived {
		return errors.New("ImageAttachment \"store_op\" is required but was not present")
	}
	return nil
}

func (strct *MeshDescriptor) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Description" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("MeshDescriptor \"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IndicesBuffer" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "indices_buffer" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("MeshDescriptor \"indices_buffer\": ")
	if tmp, err := json.Marshal(strct.IndicesBuffer); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("MeshDescriptor \"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "VertexBinding" field is required
	if strct.VertexBinding == nil {
		return nil, errors.New("vertex_binding is a required field")
	}
	// Marshal the "vertex_binding" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("MeshDescriptor \"vertex_binding\": ")
	if tmp, err := json.Marshal(strct.VertexBinding); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MeshDescriptor) UnmarshalJSON(b []byte) error {
	descriptionReceived := false
	indices_bufferReceived := false
	nameReceived := false
	vertex_bindingReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
			descriptionReceived = true
		case "indices_buffer":
			if err := json.Unmarshal([]byte(v), &strct.IndicesBuffer); err != nil {
				return err
			}
			indices_bufferReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "vertex_binding":
			if err := json.Unmarshal([]byte(v), &strct.VertexBinding); err != nil {
				return err
			}
			vertex_bindingReceived = true
		}
	}
	// check if description (a required property) was received
	if !descriptionReceived {
		return errors.New("MeshDescriptor \"description\" is required but was not present")
	}
	// check if indices_buffer (a required property) was received
	if !indices_bufferReceived {
		return errors.New("MeshDescriptor \"indices_buffer\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("MeshDescriptor \"name\" is required but was not present")
	}
	// check if vertex_binding (a required property) was received
	if !vertex_bindingReceived {
		return errors.New("MeshDescriptor \"vertex_binding\" is required but was not present")
	}
	return nil
}

func (strct *Mssa) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "AlphaEnable" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "alpha_enable" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Msaa \"alpha_enable\": ")
	if tmp, err := json.Marshal(strct.AlphaEnable); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "AlphaOne" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "alpha_one" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Msaa \"alpha_one\": ")
	if tmp, err := json.Marshal(strct.AlphaOne); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "BitSample" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "bit_sample" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Msaa \"bit_sample\": ")
	if tmp, err := json.Marshal(strct.BitSample); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Enable" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "enable" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Msaa \"enable\": ")
	if tmp, err := json.Marshal(strct.Enable); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "MinSampleShading" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "min_sample_shading" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("Msaa \"min_sample_shading\": ")
	if tmp, err := json.Marshal(strct.MinSampleShading); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Mssa) UnmarshalJSON(b []byte) error {
	alpha_enableReceived := false
	alpha_oneReceived := false
	bit_sampleReceived := false
	enableReceived := false
	min_sample_shadingReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "alpha_enable":
			if err := json.Unmarshal([]byte(v), &strct.AlphaEnable); err != nil {
				return err
			}
			alpha_enableReceived = true
		case "alpha_one":
			if err := json.Unmarshal([]byte(v), &strct.AlphaOne); err != nil {
				return err
			}
			alpha_oneReceived = true
		case "bit_sample":
			if err := json.Unmarshal([]byte(v), &strct.BitSample); err != nil {
				return err
			}
			bit_sampleReceived = true
		case "enable":
			if err := json.Unmarshal([]byte(v), &strct.Enable); err != nil {
				return err
			}
			enableReceived = true
		case "min_sample_shading":
			if err := json.Unmarshal([]byte(v), &strct.MinSampleShading); err != nil {
				return err
			}
			min_sample_shadingReceived = true
		}
	}
	// check if alpha_enable (a required property) was received
	if !alpha_enableReceived {
		return errors.New("Msaa \"alpha_enable\" is required but was not present")
	}
	// check if alpha_one (a required property) was received
	if !alpha_oneReceived {
		return errors.New("Msaa \"alpha_one\" is required but was not present")
	}
	// check if bit_sample (a required property) was received
	if !bit_sampleReceived {
		return errors.New("Msaa \"bit_sample\" is required but was not present")
	}
	// check if enable (a required property) was received
	if !enableReceived {
		return errors.New("Msaa \"enable\" is required but was not present")
	}
	// check if min_sample_shading (a required property) was received
	if !min_sample_shadingReceived {
		return errors.New("Msaa \"min_sample_shading\" is required but was not present")
	}
	return nil
}

func (strct *Pipeline) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "AttachmentBlend" field is required
	if strct.AttachmentBlend == nil {
		return nil, errors.New("attachment_blend is a required field")
	}
	// Marshal the "attachment_blend" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"attachment_blend\": ")
	if tmp, err := json.Marshal(strct.AttachmentBlend); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Constants" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "constants" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"constants\": ")
	if tmp, err := json.Marshal(strct.Constants); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "CullMode" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "cull_mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"cull_mode\": ")
	if tmp, err := json.Marshal(strct.CullMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Depth" field is required
	if strct.Depth == nil {
		return nil, errors.New("depth is a required field")
	}
	// Marshal the "depth" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"depth\": ")
	if tmp, err := json.Marshal(strct.Depth); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "FrontFace" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "front_face" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"front_face\": ")
	if tmp, err := json.Marshal(strct.FrontFace); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Layouts" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "layouts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"layouts\": ")
	if tmp, err := json.Marshal(strct.Layouts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "LineWidth" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "line_width" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"line_width\": ")
	if tmp, err := json.Marshal(strct.LineWidth); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "MeshDescriptor" field is required
	if strct.MeshDescriptor == nil {
		return nil, errors.New("mesh_descriptor is a required field")
	}
	// Marshal the "mesh_descriptor" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"mesh_descriptor\": ")
	if tmp, err := json.Marshal(strct.MeshDescriptor); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Mssa" field is required
	if strct.Mssa == nil {
		return nil, errors.New("mssa is a required field")
	}
	// Marshal the "mssa" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"mssa\": ")
	if tmp, err := json.Marshal(strct.Mssa); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "PipelineAttributes" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "pipeline_attributes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"pipeline_attributes\": ")
	if tmp, err := json.Marshal(strct.PipelineAttributes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Program" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "program" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"program\": ")
	if tmp, err := json.Marshal(strct.Program); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Renderpass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "renderpass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"renderpass\": ")
	if tmp, err := json.Marshal(strct.Renderpass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Subpass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "subpass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"subpass\": ")
	if tmp, err := json.Marshal(strct.Subpass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Topology" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "topology" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"topology\": ")
	if tmp, err := json.Marshal(strct.Topology); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Pipeline) UnmarshalJSON(b []byte) error {
	attachment_blendReceived := false
	constantsReceived := false
	cull_modeReceived := false
	depthReceived := false
	front_faceReceived := false
	layoutsReceived := false
	line_widthReceived := false
	mesh_descriptorReceived := false
	mssaReceived := false
	nameReceived := false
	pipeline_attributesReceived := false
	programReceived := false
	renderpassReceived := false
	subpassReceived := false
	topologyReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "attachment_blend":
			if err := json.Unmarshal([]byte(v), &strct.AttachmentBlend); err != nil {
				return err
			}
			attachment_blendReceived = true
		case "constants":
			if err := json.Unmarshal([]byte(v), &strct.Constants); err != nil {
				return err
			}
			constantsReceived = true
		case "cull_mode":
			if err := json.Unmarshal([]byte(v), &strct.CullMode); err != nil {
				return err
			}
			cull_modeReceived = true
		case "depth":
			if err := json.Unmarshal([]byte(v), &strct.Depth); err != nil {
				return err
			}
			depthReceived = true
		case "front_face":
			if err := json.Unmarshal([]byte(v), &strct.FrontFace); err != nil {
				return err
			}
			front_faceReceived = true
		case "layouts":
			if err := json.Unmarshal([]byte(v), &strct.Layouts); err != nil {
				return err
			}
			layoutsReceived = true
		case "line_width":
			if err := json.Unmarshal([]byte(v), &strct.LineWidth); err != nil {
				return err
			}
			line_widthReceived = true
		case "mesh_descriptor":
			if err := json.Unmarshal([]byte(v), &strct.MeshDescriptor); err != nil {
				return err
			}
			mesh_descriptorReceived = true
		case "mssa":
			if err := json.Unmarshal([]byte(v), &strct.Mssa); err != nil {
				return err
			}
			mssaReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "pipeline_attributes":
			if err := json.Unmarshal([]byte(v), &strct.PipelineAttributes); err != nil {
				return err
			}
			pipeline_attributesReceived = true
		case "program":
			if err := json.Unmarshal([]byte(v), &strct.Program); err != nil {
				return err
			}
			programReceived = true
		case "renderpass":
			if err := json.Unmarshal([]byte(v), &strct.Renderpass); err != nil {
				return err
			}
			renderpassReceived = true
		case "subpass":
			if err := json.Unmarshal([]byte(v), &strct.Subpass); err != nil {
				return err
			}
			subpassReceived = true
		case "topology":
			if err := json.Unmarshal([]byte(v), &strct.Topology); err != nil {
				return err
			}
			topologyReceived = true
		}
	}
	// check if attachment_blend (a required property) was received
	if !attachment_blendReceived {
		return errors.New("Pipeline \"attachment_blend\" is required but was not present")
	}
	// check if constants (a required property) was received
	if !constantsReceived {
		return errors.New("Pipeline \"constants\" is required but was not present")
	}
	// check if cull_mode (a required property) was received
	if !cull_modeReceived {
		return errors.New("Pipeline \"cull_mode\" is required but was not present")
	}
	// check if depth (a required property) was received
	if !depthReceived {
		return errors.New("Pipeline \"depth\" is required but was not present")
	}
	// check if front_face (a required property) was received
	if !front_faceReceived {
		return errors.New("Pipeline \"front_face\" is required but was not present")
	}
	// check if layouts (a required property) was received
	if !layoutsReceived {
		return errors.New("Pipeline \"layouts\" is required but was not present")
	}
	// check if line_width (a required property) was received
	if !line_widthReceived {
		return errors.New("Pipeline \"line_width\" is required but was not present")
	}
	// check if mesh_descriptor (a required property) was received
	if !mesh_descriptorReceived {
		return errors.New("Pipeline \"mesh_descriptor\" is required but was not present")
	}
	// check if mssa (a required property) was received
	if !mssaReceived {
		return errors.New("Pipeline \"mssa\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("Pipeline \"name\" is required but was not present")
	}
	// check if pipeline_attributes (a required property) was received
	if !pipeline_attributesReceived {
		return errors.New("Pipeline \"pipeline_attributes\" is required but was not present")
	}
	// check if program (a required property) was received
	if !programReceived {
		return errors.New("Pipeline \"program\" is required but was not present")
	}
	// check if renderpass (a required property) was received
	if !renderpassReceived {
		return errors.New("Pipeline \"renderpass\" is required but was not present")
	}
	// check if subpass (a required property) was received
	if !subpassReceived {
		return errors.New("Pipeline \"subpass\" is required but was not present")
	}
	// check if topology (a required property) was received
	if !topologyReceived {
		return errors.New("Pipeline \"topology\" is required but was not present")
	}
	return nil
}

func (strct *PipelineAttribute) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Binding" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "binding" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"binding\": ")
	if tmp, err := json.Marshal(strct.Binding); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Format" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "format" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"format\": ")
	if tmp, err := json.Marshal(strct.Format); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Location" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Offset" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "offset" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"offset\": ")
	if tmp, err := json.Marshal(strct.Offset); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PipelineAttribute) UnmarshalJSON(b []byte) error {
	bindingReceived := false
	formatReceived := false
	locationReceived := false
	nameReceived := false
	offsetReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "binding":
			if err := json.Unmarshal([]byte(v), &strct.Binding); err != nil {
				return err
			}
			bindingReceived = true
		case "format":
			if err := json.Unmarshal([]byte(v), &strct.Format); err != nil {
				return err
			}
			formatReceived = true
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
			locationReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "offset":
			if err := json.Unmarshal([]byte(v), &strct.Offset); err != nil {
				return err
			}
			offsetReceived = true
		}
	}
	// check if binding (a required property) was received
	if !bindingReceived {
		return errors.New("PipelineAttribute \"binding\" is required but was not present")
	}
	// check if format (a required property) was received
	if !formatReceived {
		return errors.New("PipelineAttribute \"format\" is required but was not present")
	}
	// check if location (a required property) was received
	if !locationReceived {
		return errors.New("PipelineAttribute \"location\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("PipelineAttribute \"name\" is required but was not present")
	}
	// check if offset (a required property) was received
	if !offsetReceived {
		return errors.New("PipelineAttribute \"offset\" is required but was not present")
	}
	return nil
}

func (strct *PushConstant) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Offset" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "offset" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"offset\": ")
	if tmp, err := json.Marshal(strct.Offset); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ShaderStages" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "shader_stages" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shader_stages\": ")
	if tmp, err := json.Marshal(strct.ShaderStages); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Size" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "size" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"size\": ")
	if tmp, err := json.Marshal(strct.Size); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PushConstant) UnmarshalJSON(b []byte) error {
	nameReceived := false
	offsetReceived := false
	shader_stagesReceived := false
	sizeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "offset":
			if err := json.Unmarshal([]byte(v), &strct.Offset); err != nil {
				return err
			}
			offsetReceived = true
		case "shader_stages":
			if err := json.Unmarshal([]byte(v), &strct.ShaderStages); err != nil {
				return err
			}
			shader_stagesReceived = true
		case "size":
			if err := json.Unmarshal([]byte(v), &strct.Size); err != nil {
				return err
			}
			sizeReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("PushConstant \"name\" is required but was not present")
	}
	// check if offset (a required property) was received
	if !offsetReceived {
		return errors.New("PushConstant \"offset\" is required but was not present")
	}
	// check if shader_stages (a required property) was received
	if !shader_stagesReceived {
		return errors.New("PushConstant \"shader_stages\" is required but was not present")
	}
	// check if size (a required property) was received
	if !sizeReceived {
		return errors.New("PushConstant \"size\" is required but was not present")
	}
	return nil
}

func (strct *Queue) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Family" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "family" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"family\": ")
	if tmp, err := json.Marshal(strct.Family); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Index" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Queue) UnmarshalJSON(b []byte) error {
	familyReceived := false
	indexReceived := false
	nameReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "family":
			if err := json.Unmarshal([]byte(v), &strct.Family); err != nil {
				return err
			}
			familyReceived = true
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
			indexReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		}
	}
	// check if family (a required property) was received
	if !familyReceived {
		return errors.New("Queue \"family\" is required but was not present")
	}
	// check if index (a required property) was received
	if !indexReceived {
		return errors.New("Queue \"index\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("Queue \"name\" is required but was not present")
	}
	return nil
}

func (strct *Renderpass) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "CommandBuffer" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "command_buffer" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"command_buffer\": ")
	if tmp, err := json.Marshal(strct.CommandBuffer); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Subpasses" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "subpasses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"subpasses\": ")
	if tmp, err := json.Marshal(strct.Subpasses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Renderpass) UnmarshalJSON(b []byte) error {
	command_bufferReceived := false
	nameReceived := false
	subpassesReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "command_buffer":
			if err := json.Unmarshal([]byte(v), &strct.CommandBuffer); err != nil {
				return err
			}
			command_bufferReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "subpasses":
			if err := json.Unmarshal([]byte(v), &strct.Subpasses); err != nil {
				return err
			}
			subpassesReceived = true
		}
	}
	// check if command_buffer (a required property) was received
	if !command_bufferReceived {
		return errors.New("Renderpass \"command_buffer\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("Renderpass \"name\" is required but was not present")
	}
	// check if subpasses (a required property) was received
	if !subpassesReceived {
		return errors.New("Renderpass \"subpasses\" is required but was not present")
	}
	return nil
}

func (strct *Shader) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Sources" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sources" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sources\": ")
	if tmp, err := json.Marshal(strct.Sources); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Shader) UnmarshalJSON(b []byte) error {
	nameReceived := false
	sourcesReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "sources":
			if err := json.Unmarshal([]byte(v), &strct.Sources); err != nil {
				return err
			}
			sourcesReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("Shader \"name\" is required but was not present")
	}
	// check if sources (a required property) was received
	if !sourcesReceived {
		return errors.New("Shader \"sources\" is required but was not present")
	}
	return nil
}

func (strct *Source) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Type" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "type" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Url" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "url" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Source) UnmarshalJSON(b []byte) error {
	typeReceived := false
	urlReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "type":
			if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
				return err
			}
			typeReceived = true
		case "url":
			if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
				return err
			}
			urlReceived = true
		}
	}
	// check if type (a required property) was received
	if !typeReceived {
		return errors.New("Source \"type\" is required but was not present")
	}
	// check if url (a required property) was received
	if !urlReceived {
		return errors.New("Source \"url\" is required but was not present")
	}
	return nil
}

func (strct *Subpass) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Bindpoint" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "bindpoint" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"bindpoint\": ")
	if tmp, err := json.Marshal(strct.Bindpoint); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ColorAttachments" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "color_attachments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"color_attachments\": ")
	if tmp, err := json.Marshal(strct.ColorAttachments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Dependencies" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "dependencies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"dependencies\": ")
	if tmp, err := json.Marshal(strct.Dependencies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "DepthAttachments" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "depth_attachments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"depth_attachments\": ")
	if tmp, err := json.Marshal(strct.DepthAttachments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "InputAttachments" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "input_attachments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"input_attachments\": ")
	if tmp, err := json.Marshal(strct.InputAttachments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "PreserveAttachments" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "preserve_attachments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"preserve_attachments\": ")
	if tmp, err := json.Marshal(strct.PreserveAttachments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Subpass) UnmarshalJSON(b []byte) error {
	bindpointReceived := false
	color_attachmentsReceived := false
	dependenciesReceived := false
	depth_attachmentsReceived := false
	input_attachmentsReceived := false
	preserve_attachmentsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "bindpoint":
			if err := json.Unmarshal([]byte(v), &strct.Bindpoint); err != nil {
				return err
			}
			bindpointReceived = true
		case "color_attachments":
			if err := json.Unmarshal([]byte(v), &strct.ColorAttachments); err != nil {
				return err
			}
			color_attachmentsReceived = true
		case "dependencies":
			if err := json.Unmarshal([]byte(v), &strct.Dependencies); err != nil {
				return err
			}
			dependenciesReceived = true
		case "depth_attachments":
			if err := json.Unmarshal([]byte(v), &strct.DepthAttachments); err != nil {
				return err
			}
			depth_attachmentsReceived = true
		case "input_attachments":
			if err := json.Unmarshal([]byte(v), &strct.InputAttachments); err != nil {
				return err
			}
			input_attachmentsReceived = true
		case "preserve_attachments":
			if err := json.Unmarshal([]byte(v), &strct.PreserveAttachments); err != nil {
				return err
			}
			preserve_attachmentsReceived = true
		}
	}
	// check if bindpoint (a required property) was received
	if !bindpointReceived {
		return errors.New("Subpass \"bindpoint\" is required but was not present")
	}
	// check if color_attachments (a required property) was received
	if !color_attachmentsReceived {
		return errors.New("Subpass \"color_attachments\" is required but was not present")
	}
	// check if dependencies (a required property) was received
	if !dependenciesReceived {
		return errors.New("Subpass \"dependencies\" is required but was not present")
	}
	// check if depth_attachments (a required property) was received
	if !depth_attachmentsReceived {
		return errors.New("Subpass \"depth_attachments\" is required but was not present")
	}
	// check if input_attachments (a required property) was received
	if !input_attachmentsReceived {
		return errors.New("Subpass \"input_attachments\" is required but was not present")
	}
	// check if preserve_attachments (a required property) was received
	if !preserve_attachmentsReceived {
		return errors.New("Subpass \"preserve_attachments\" is required but was not present")
	}
	return nil
}

func (strct *SubpassDependency) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "DstAccessMask" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "dst_access_mask" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"dst_access_mask\": ")
	if tmp, err := json.Marshal(strct.DstAccessMask); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "DstMask" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "dst_mask" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"dst_mask\": ")
	if tmp, err := json.Marshal(strct.DstMask); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "DstSubpass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "dst_subpass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"dst_subpass\": ")
	if tmp, err := json.Marshal(strct.DstSubpass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "SrcAccessMask" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "src_access_mask" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"src_access_mask\": ")
	if tmp, err := json.Marshal(strct.SrcAccessMask); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "SrcMask" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "src_mask" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"src_mask\": ")
	if tmp, err := json.Marshal(strct.SrcMask); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "SrcSubpass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "src_subpass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"src_subpass\": ")
	if tmp, err := json.Marshal(strct.SrcSubpass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SubpassDependency) UnmarshalJSON(b []byte) error {
	dst_access_maskReceived := false
	dst_maskReceived := false
	dst_subpassReceived := false
	nameReceived := false
	src_access_maskReceived := false
	src_maskReceived := false
	src_subpassReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "dst_access_mask":
			if err := json.Unmarshal([]byte(v), &strct.DstAccessMask); err != nil {
				return err
			}
			dst_access_maskReceived = true
		case "dst_mask":
			if err := json.Unmarshal([]byte(v), &strct.DstMask); err != nil {
				return err
			}
			dst_maskReceived = true
		case "dst_subpass":
			if err := json.Unmarshal([]byte(v), &strct.DstSubpass); err != nil {
				return err
			}
			dst_subpassReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "src_access_mask":
			if err := json.Unmarshal([]byte(v), &strct.SrcAccessMask); err != nil {
				return err
			}
			src_access_maskReceived = true
		case "src_mask":
			if err := json.Unmarshal([]byte(v), &strct.SrcMask); err != nil {
				return err
			}
			src_maskReceived = true
		case "src_subpass":
			if err := json.Unmarshal([]byte(v), &strct.SrcSubpass); err != nil {
				return err
			}
			src_subpassReceived = true
		}
	}
	// check if dst_access_mask (a required property) was received
	if !dst_access_maskReceived {
		return errors.New("SubpassDependency \"dst_access_mask\" is required but was not present")
	}
	// check if dst_mask (a required property) was received
	if !dst_maskReceived {
		return errors.New("SubpassDependency \"dst_mask\" is required but was not present")
	}
	// check if dst_subpass (a required property) was received
	if !dst_subpassReceived {
		return errors.New("SubpassDependency \"dst_subpass\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("SubpassDependency \"name\" is required but was not present")
	}
	// check if src_access_mask (a required property) was received
	if !src_access_maskReceived {
		return errors.New("SubpassDependency \"src_access_mask\" is required but was not present")
	}
	// check if src_mask (a required property) was received
	if !src_maskReceived {
		return errors.New("SubpassDependency \"src_mask\" is required but was not present")
	}
	// check if src_subpass (a required property) was received
	if !src_subpassReceived {
		return errors.New("SubpassDependency \"src_subpass\" is required but was not present")
	}
	return nil
}

func (strct *Swapchain) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ColorAttachment" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "color_attachment" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"color_attachment\": ")
	if tmp, err := json.Marshal(strct.ColorAttachment); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "DepthAttachment" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "depth_attachment" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"depth_attachment\": ")
	if tmp, err := json.Marshal(strct.DepthAttachment); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Swapchain) UnmarshalJSON(b []byte) error {
	color_attachmentReceived := false
	depth_attachmentReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "color_attachment":
			if err := json.Unmarshal([]byte(v), &strct.ColorAttachment); err != nil {
				return err
			}
			color_attachmentReceived = true
		case "depth_attachment":
			if err := json.Unmarshal([]byte(v), &strct.DepthAttachment); err != nil {
				return err
			}
			depth_attachmentReceived = true
		}
	}
	// check if color_attachment (a required property) was received
	if !color_attachmentReceived {
		return errors.New("Swapchain \"color_attachment\" is required but was not present")
	}
	// check if depth_attachment (a required property) was received
	if !depth_attachmentReceived {
		return errors.New("Swapchain \"depth_attachment\" is required but was not present")
	}
	return nil
}

func (strct *VertexBinding) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Binding" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "binding" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"binding\": ")
	if tmp, err := json.Marshal(strct.Binding); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Rate" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "rate" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rate\": ")
	if tmp, err := json.Marshal(strct.Rate); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StrideBytes" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "stride_bytes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stride_bytes\": ")
	if tmp, err := json.Marshal(strct.StrideBytes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *VertexBinding) UnmarshalJSON(b []byte) error {
	bindingReceived := false
	rateReceived := false
	stride_bytesReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "binding":
			if err := json.Unmarshal([]byte(v), &strct.Binding); err != nil {
				return err
			}
			bindingReceived = true
		case "rate":
			if err := json.Unmarshal([]byte(v), &strct.Rate); err != nil {
				return err
			}
			rateReceived = true
		case "stride_bytes":
			if err := json.Unmarshal([]byte(v), &strct.StrideBytes); err != nil {
				return err
			}
			stride_bytesReceived = true
		}
	}
	// check if binding (a required property) was received
	if !bindingReceived {
		return errors.New("VertexBinding \"binding\" is required but was not present")
	}
	// check if rate (a required property) was received
	if !rateReceived {
		return errors.New("VertexBinding \"rate\" is required but was not present")
	}
	// check if stride_bytes (a required property) was received
	if !stride_bytesReceived {
		return errors.New("VertexBinding \"stride_bytes\" is required but was not present")
	}
	return nil
}

func (strct *Vlk) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "buffer" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"buffer\": ")
	if tmp, err := json.Marshal(strct.Buffer); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "command_buffer" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"command_buffer\": ")
	if tmp, err := json.Marshal(strct.CommandBuffer); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "CommandBuffers" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "command_buffers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"command_buffers\": ")
	if tmp, err := json.Marshal(strct.CommandBuffers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Config" field is required
	if strct.Config == nil {
		return nil, errors.New("config is a required field")
	}
	// Marshal the "config" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"config\": ")
	if tmp, err := json.Marshal(strct.Config); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "descriptor_layout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"descriptor_layout\": ")
	if tmp, err := json.Marshal(strct.DescriptorLayout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "framebuffer" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"framebuffer\": ")
	if tmp, err := json.Marshal(strct.Framebuffer); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Framebuffers" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "framebuffers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"framebuffers\": ")
	if tmp, err := json.Marshal(strct.Framebuffers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "image" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"image\": ")
	if tmp, err := json.Marshal(strct.Image); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "image_attachment" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"image_attachment\": ")
	if tmp, err := json.Marshal(strct.ImageAttachment); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IndiceBuffers" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "indice_buffers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"indice_buffers\": ")
	if tmp, err := json.Marshal(strct.IndiceBuffers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "pipeline" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"pipeline\": ")
	if tmp, err := json.Marshal(strct.Pipeline); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "pipeline_attribute" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"pipeline_attribute\": ")
	if tmp, err := json.Marshal(strct.PipelineAttribute); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Pipelines" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "pipelines" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"pipelines\": ")
	if tmp, err := json.Marshal(strct.Pipelines); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "push_constant" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"push_constant\": ")
	if tmp, err := json.Marshal(strct.PushConstant); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "queue" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"queue\": ")
	if tmp, err := json.Marshal(strct.Queue); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Queues" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "queues" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"queues\": ")
	if tmp, err := json.Marshal(strct.Queues); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "renderpass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"renderpass\": ")
	if tmp, err := json.Marshal(strct.Renderpass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Renderpasses" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "renderpasses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"renderpasses\": ")
	if tmp, err := json.Marshal(strct.Renderpasses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "SamplerBuffers" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sampler_buffers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sampler_buffers\": ")
	if tmp, err := json.Marshal(strct.SamplerBuffers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Schema" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "schema" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"schema\": ")
	if tmp, err := json.Marshal(strct.Schema); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "shader" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shader\": ")
	if tmp, err := json.Marshal(strct.Shader); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ShaderPrograms" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "shader_programs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shader_programs\": ")
	if tmp, err := json.Marshal(strct.ShaderPrograms); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "source" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"source\": ")
	if tmp, err := json.Marshal(strct.Source); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "subpass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"subpass\": ")
	if tmp, err := json.Marshal(strct.Subpass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "subpass_dependency" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"subpass_dependency\": ")
	if tmp, err := json.Marshal(strct.SubpassDependency); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "swapchain" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"swapchain\": ")
	if tmp, err := json.Marshal(strct.Swapchain); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "UniformBuffers" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "uniform_buffers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"uniform_buffers\": ")
	if tmp, err := json.Marshal(strct.UniformBuffers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "VertexBuffers" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "vertex_buffers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"vertex_buffers\": ")
	if tmp, err := json.Marshal(strct.VertexBuffers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Vlk) UnmarshalJSON(b []byte) error {
	command_buffersReceived := false
	configReceived := false
	framebuffersReceived := false
	indice_buffersReceived := false
	pipelinesReceived := false
	queuesReceived := false
	renderpassesReceived := false
	sampler_buffersReceived := false
	schemaReceived := false
	shader_programsReceived := false
	uniform_buffersReceived := false
	vertex_buffersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "buffer":
			if err := json.Unmarshal([]byte(v), &strct.Buffer); err != nil {
				return err
			}
		case "command_buffer":
			if err := json.Unmarshal([]byte(v), &strct.CommandBuffer); err != nil {
				return err
			}
		case "command_buffers":
			if err := json.Unmarshal([]byte(v), &strct.CommandBuffers); err != nil {
				return err
			}
			command_buffersReceived = true
		case "config":
			if err := json.Unmarshal([]byte(v), &strct.Config); err != nil {
				return err
			}
			configReceived = true
		case "descriptor_layout":
			if err := json.Unmarshal([]byte(v), &strct.DescriptorLayout); err != nil {
				return err
			}
		case "framebuffer":
			if err := json.Unmarshal([]byte(v), &strct.Framebuffer); err != nil {
				return err
			}
		case "framebuffers":
			if err := json.Unmarshal([]byte(v), &strct.Framebuffers); err != nil {
				return err
			}
			framebuffersReceived = true
		case "image":
			if err := json.Unmarshal([]byte(v), &strct.Image); err != nil {
				return err
			}
		case "image_attachment":
			if err := json.Unmarshal([]byte(v), &strct.ImageAttachment); err != nil {
				return err
			}
		case "indice_buffers":
			if err := json.Unmarshal([]byte(v), &strct.IndiceBuffers); err != nil {
				return err
			}
			indice_buffersReceived = true
		case "pipeline":
			if err := json.Unmarshal([]byte(v), &strct.Pipeline); err != nil {
				return err
			}
		case "pipeline_attribute":
			if err := json.Unmarshal([]byte(v), &strct.PipelineAttribute); err != nil {
				return err
			}
		case "pipelines":
			if err := json.Unmarshal([]byte(v), &strct.Pipelines); err != nil {
				return err
			}
			pipelinesReceived = true
		case "push_constant":
			if err := json.Unmarshal([]byte(v), &strct.PushConstant); err != nil {
				return err
			}
		case "queue":
			if err := json.Unmarshal([]byte(v), &strct.Queue); err != nil {
				return err
			}
		case "queues":
			if err := json.Unmarshal([]byte(v), &strct.Queues); err != nil {
				return err
			}
			queuesReceived = true
		case "renderpass":
			if err := json.Unmarshal([]byte(v), &strct.Renderpass); err != nil {
				return err
			}
		case "renderpasses":
			if err := json.Unmarshal([]byte(v), &strct.Renderpasses); err != nil {
				return err
			}
			renderpassesReceived = true
		case "sampler_buffers":
			if err := json.Unmarshal([]byte(v), &strct.SamplerBuffers); err != nil {
				return err
			}
			sampler_buffersReceived = true
		case "schema":
			if err := json.Unmarshal([]byte(v), &strct.Schema); err != nil {
				return err
			}
			schemaReceived = true
		case "shader":
			if err := json.Unmarshal([]byte(v), &strct.Shader); err != nil {
				return err
			}
		case "shader_programs":
			if err := json.Unmarshal([]byte(v), &strct.ShaderPrograms); err != nil {
				return err
			}
			shader_programsReceived = true
		case "source":
			if err := json.Unmarshal([]byte(v), &strct.Source); err != nil {
				return err
			}
		case "subpass":
			if err := json.Unmarshal([]byte(v), &strct.Subpass); err != nil {
				return err
			}
		case "subpass_dependency":
			if err := json.Unmarshal([]byte(v), &strct.SubpassDependency); err != nil {
				return err
			}
		case "swapchain":
			if err := json.Unmarshal([]byte(v), &strct.Swapchain); err != nil {
				return err
			}
		case "uniform_buffers":
			if err := json.Unmarshal([]byte(v), &strct.UniformBuffers); err != nil {
				return err
			}
			uniform_buffersReceived = true
		case "vertex_buffers":
			if err := json.Unmarshal([]byte(v), &strct.VertexBuffers); err != nil {
				return err
			}
			vertex_buffersReceived = true
		}
	}
	// check if command_buffers (a required property) was received
	if !command_buffersReceived {
		return errors.New("Vlk \"command_buffers\" is required but was not present")
	}
	// check if config (a required property) was received
	if !configReceived {
		return errors.New("Vlk \"config\" is required but was not present")
	}
	// check if framebuffers (a required property) was received
	if !framebuffersReceived {
		return errors.New("Vlk \"framebuffers\" is required but was not present")
	}
	// check if indice_buffers (a required property) was received
	if !indice_buffersReceived {
		return errors.New("Vlk \"indice_buffers\" is required but was not present")
	}
	// check if pipelines (a required property) was received
	if !pipelinesReceived {
		return errors.New("Vlk \"pipelines\" is required but was not present")
	}
	// check if queues (a required property) was received
	if !queuesReceived {
		return errors.New("Vlk \"queues\" is required but was not present")
	}
	// check if renderpasses (a required property) was received
	if !renderpassesReceived {
		return errors.New("Vlk \"renderpasses\" is required but was not present")
	}
	// check if sampler_buffers (a required property) was received
	if !sampler_buffersReceived {
		return errors.New("Vlk \"sampler_buffers\" is required but was not present")
	}
	// check if schema (a required property) was received
	if !schemaReceived {
		return errors.New("Vlk \"schema\" is required but was not present")
	}
	// check if shader_programs (a required property) was received
	if !shader_programsReceived {
		return errors.New("Vlk \"shader_programs\" is required but was not present")
	}
	// check if uniform_buffers (a required property) was received
	if !uniform_buffersReceived {
		return errors.New("Vlk \"uniform_buffers\" is required but was not present")
	}
	// check if vertex_buffers (a required property) was received
	if !vertex_buffersReceived {
		return errors.New("Vlk \"vertex_buffers\" is required but was not present")
	}
	return nil
}

func (strct *Window) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Height" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "height" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"height\": ")
	if tmp, err := json.Marshal(strct.Height); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Width" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "width" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"width\": ")
	if tmp, err := json.Marshal(strct.Width); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Window) UnmarshalJSON(b []byte) error {
	heightReceived := false
	widthReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "height":
			if err := json.Unmarshal([]byte(v), &strct.Height); err != nil {
				return err
			}
			heightReceived = true
		case "width":
			if err := json.Unmarshal([]byte(v), &strct.Width); err != nil {
				return err
			}
			widthReceived = true
		}
	}
	// check if height (a required property) was received
	if !heightReceived {
		return errors.New("Window \"height\" is required but was not present")
	}
	// check if width (a required property) was received
	if !widthReceived {
		return errors.New("Window \"width\" is required but was not present")
	}
	return nil
}
