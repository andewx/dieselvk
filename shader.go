package dieselvk

import (
	"fmt"
	"io/ioutil"
	"os"

	vk "github.com/vulkan-go/vulkan"
)

const (
	VERTEX  = 0
	FRAG    = 1
	COMPUTE = 2
	GEOM    = 3
	TESS    = 4
)

type CoreShader struct {
	shader_descriptors     vk.DescriptorSet //Key: (Shader Program ID Key) Value: vkDescriptor Set
	compute_shader_modules vk.ShaderModule  //Key: (Shader Program ID Key) Value: Vulkan Shader Module
	shader_paths           map[string]int   //Key: Shader path, Value : Shader type
	shader_programs        map[string]*ShaderProgram
}

func NewCoreShader() *CoreShader {
	var core CoreShader
	core.shader_programs = make(map[string]*ShaderProgram, 1)
	core.shader_paths = make(map[string]int, 2)
	return &core
}

func (core *CoreShader) AddShaderPath(path string, shader_type int) {
	core.shader_paths[path] = shader_type
}

func (core *CoreShader) CreateProgram(name string, instance CoreInstance, paths []string) {

	var pg ShaderProgram

	for _, path := range paths {

		path_id := core.shader_paths[path]
		var bindingModule vk.ShaderModule
		core.LoadShaderModule(instance, path, &bindingModule)

		if path_id == VERTEX {
			pg.vertex_shader_modules = &bindingModule
		}

		if path_id == FRAG {
			pg.fragment_shader_modules = &bindingModule
		}

	}
	core.shader_programs[name] = &pg

}

type ShaderProgram struct {
	vertex_shader_modules   *vk.ShaderModule //Key: (Shader Program ID Key) Value: Vulkan Shader Module
	fragment_shader_modules *vk.ShaderModule //Key: (Shader PRogram ID Key) Value: Vulkan Shader Module
}

func (core *CoreShader) LoadShaderModule(instance CoreInstance, path string, out_shader *vk.ShaderModule) {
	buffer, err := ioutil.ReadFile(path)

	if err != nil {
		return
	}
	//Vulkan expects to recieve type uint32 data
	convertBytes := sliceUint32(buffer)
	module := vk.ShaderModuleCreateInfo{}
	module.SType = vk.StructureTypeShaderModuleCreateInfo
	module.PNext = nil
	module.CodeSize = uint(len(buffer))
	module.PCode = convertBytes
	handle := instance.GetHandle()

	//Create module
	var shaderModule vk.ShaderModule

	res := vk.CreateShaderModule(handle, &module, nil, &shaderModule)

	if res != vk.Success {
		fmt.Printf("Error unable to create shader module in LoadShaderModule()\nExiting...")
		os.Exit(1)
	}

	*out_shader = shaderModule

	return

}
