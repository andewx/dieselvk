package test

import (
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/andewx/dieselvk"
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

func TestAllocate(t *testing.T) {

	runtime.LockOSThread()
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

	if err := vk.Init(); err != nil {
		t.Errorf("Unable to initialize application %v", err)
		return
	}

	log.Printf("Creating Vulkan Instance...\n")

	dirs, _ := os.Getwd()

	//Shaders
	shaders := []string{
		dirs + "/shaders/frag.spv",
		dirs + "/shaders/vert.spv",
	}

	//Vulkan Core Config
	config := make(map[string]string, 10)
	config["extensions"] = "default"
	config["display"] = "false"
	config["debug"] = "false"
	config["gpu_exclusive_instance"] = "true"
	config["validation"] = "VK_LAYER_KHRONOS_validation"

	//Vulkan Desired Instances
	instances := make([]dieselvk.Instance, 1)
	instances[0] = dieselvk.Instance{Name: "default", Selector: dieselvk.DEVICE_INSTANCE, Gpu_id: 0}
	vulkan_core := dieselvk.NewBaseCore("json/vlk_example.json", "default", nil)
	vulkan_core.CreateInstance()
	render := vulkan_core.GetInstance()

	//Configue instance shaders
	render.AddShaderPath(shaders[0], dieselvk.FRAG)
	render.AddShaderPath(shaders[1], dieselvk.VERTEX)
	render.NewProgram(shaders, "default")

	//Add in vertex buffers
	vertices := []float32{-1.0, 0.0, 0.0, 0.0, -1.0, 0.0, 1.0, 0.0, 0.0}
	render.AddVertexBuffer(vertices, "triangle")

	//Uniform Layout Buffers - With struct2bytes() type function we could combine this into a single layout
	model := []float32{1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
	}

	mvp := []float32{}
	mvp = append(mvp, model...)

	//Adds Layout buffers to "uniform bufffer objects"
	bindings := []int{2, 2, 2}
	render.AddLayoutBuffer(mvp, "mvp", vk.BufferUsageFlags(vk.BufferUsageUniformBufferBit))
	render.BindUniforms(render.GetLayoutBuffers(), bindings)

	render.AllocatorUsage()

	vulkan_core.Release()

}
