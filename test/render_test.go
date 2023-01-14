package test

import (
	"log"
	"math"
	"os"
	"runtime"
	"testing"

	"github.com/andewx/dieselvk"
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

const (
	WIDTH  = 500
	HEIGHT = 500
)

func TestRender(t *testing.T) {

	runtime.LockOSThread()
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.True)
	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

	if err := vk.Init(); err != nil {
		t.Errorf("Unable to initialize application %v", err)
		return
	}

	log.Printf("Creating Vulkan Instance...\n")

	window, errW := glfw.CreateWindow(WIDTH, HEIGHT, "Vulkan", nil, nil)

	if errW != nil {
		panic(errW)
	}

	//Dir
	dirs, derr := os.Getwd()
	if derr != nil {
		return
	}

	//Shaders
	shaders := []string{
		dirs + "/shaders/frag.spv",
		dirs + "/shaders/vert.spv",
	}

	//Vulkan Core Config
	config := make(map[string]string, 10)
	config["extensions"] = "default"
	config["display"] = "true"
	config["debug"] = "false"
	config["gpu_exclusive_instance"] = "true"
	config["validation"] = "VK_LAYER_KHRONOS_validation"

	//Vulkan Desired Instances
	instances := make([]dieselvk.Instance, 1)
	instances[0] = dieselvk.Instance{Name: "default", Selector: dieselvk.RENDER_INSTANCE, Gpu_id: 0}
	vulkan_core := dieselvk.NewBaseCore(config, "default", 5, 5, window)
	if err := vulkan_core.CreateInstance(instances); err != nil {
		t.Errorf("Error could not create instance")
	}
	render := vulkan_core.GetInstance("default").(*dieselvk.CoreRenderInstance)

	//Create shaders
	render.AddShaderPath(shaders[0], dieselvk.FRAG)
	render.AddShaderPath(shaders[1], dieselvk.VERTEX)
	render.NewProgram(shaders, "default")

	//Add in uniform buffers for matrices
	fov := 45.0
	f := 100.0
	n := 0.1
	f1 := -float32(fov/f - n)
	f2 := -float32(f*n/f - n)

	S := float32(math.Tan(fov / 2 * math.Pi / 180))

	//Uniform Layout Buffers - With struct2bytes() type function we could combine this into a single layout
	model := []float32{1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	view := []float32{1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	proj := []float32{S, 0.0, 0.0, 0.0,
		0.0, S, 0.0, 0.0,
		0.0, 0.0, f1, -1.0,
		0.0, 0.0, f2, 0.0,
	}

	mvp := []float32{}
	mvp = append(mvp, model...)
	mvp = append(mvp, view...)
	mvp = append(mvp, proj...)

	//Adds Layout buffers to "uniform bufffer objects"
	bindings := []int{0, 0, 0}
	render.AddLayoutBuffer(mvp, "mvp", vk.BufferUsageFlags(vk.BufferUsageUniformBufferBit))
	render.BindUniforms(render.GetLayoutBuffers(), bindings)

	//Add in vertex buffers
	vertices := []float32{-1.0, 1.0, -1.0, 0.0, -1.0, -1.0, 1.0, 1.0, -1.0}
	render.AddVertexBuffer(vertices, "triangle")

	//Configure renderpasses + pipelines
	render.NewSwapchain()
	render.AddRenderPass("rp0")
	render.AddPipeline("pipe0", "default", *render.GetVertexBuffer("triangle"), "rp0")
	render.SetupCommands()

	for !window.ShouldClose() {
		render.Update(0.0)
		glfw.PollEvents()
	}

	render.AllocatorUsage()

	vulkan_core.Release()

}

func Mat4Mul(a []float32, b []float32) []float32 {
	c := make([]float32, 16)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ij := i*j + j
			mul := float32(0.0)
			for k := 0; k < 4; k++ {
				row_entry := i*j + k
				col_entry := j + i*k
				mul += a[row_entry] * b[col_entry]
			}
			c[ij] = mul
		}
	}
	return c
}

func RotateX(angle float64, b []float32) []float32 {

	cs := float32(math.Cos(angle))
	ss := float32(math.Sin(angle))

	rot := []float32{
		1.0, 0.0, 0.0, 0.0,
		0.0, cs, -ss, 0,
		0.0, ss, cs, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	return Mat4Mul(rot, b)
}

func RotateY(angle float64, b []float32) []float32 {

	cs := float32(math.Cos(angle))
	ss := float32(math.Sin(angle))

	rot := []float32{
		cs, 0.0, ss, 0.0,
		0.0, 1.0, 0.0, 0.0,
		-ss, 0.0, cs, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	return Mat4Mul(rot, b)
}

func Transpose4x4(vec []float32) []float32 {
	nVec := make([]float32, len(vec))

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a := 4*i + j
			b := 4*j + i
			nVec[b] = vec[a]
		}
	}
	return nVec
}
