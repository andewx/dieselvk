package dieselvk

import (
	"log"
	"os"

	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

const (
	RENDER_INSTANCE  = 0
	DEVICE_INSTANCE  = 1
	COMPUTE_INSTANCe = 2
)

type Instance struct {
	Name     string
	Selector int
	Gpu_id   int
}

type Resource interface {
	Release()
}

//Base DieselVK Core vulkan manager with GLFW native host management
//Core structure properties are private members to enforce future interface
//compliance with outside packages. The Vulkan core manager manages the availability
//of devices and capabilities to enfore instance creation and management. Also holds
//global type information which could be useful to multiple vulkan instances in an application
//which includes buffers and textures. Light objects in Vulkan do not always warrant a Core Abstraction
type BaseCore struct {

	//Core Implementation Context Properties
	display    CoreDisplay
	core_props map[string]string
	name       string
	info_log   *log.Logger
	error_log  *log.Logger
	warn_log   *log.Logger

	//Map string id & tagging
	instance_names []string

	//List of device bidings
	logical_devices map[string]CoreDevice

	//Per Instance/Device Handles where key is the instance global id key used for accessing other held resources
	instances map[string]CoreInstance //Key: (Instance_Name) Value: Vulkan Instance

	//Images/Buffer Data
	images         map[string]CoreImage  //Key: (Unique Image ID)
	vertex_buffers map[string]CoreBuffer //Key: Unique Buffer Key
	indice_buffers map[string]CoreBuffer //Key: Unique Buffer Key
	uv_buffers     map[string]CoreBuffer //Key: Unique Buffer Key
	color_buffers  map[string]CoreBuffer //Key: Unique Buffer Key
	attr_buffers   map[string]CoreBuffer //Key: Unique Buffer Key

	//Shaders
	shaders  *CoreShader
	uniforms map[string]int //Uniform location mapping

}

//Instanitates a new core context allocation sizes, default allocation prevents buffer copies but is just used to instantiate map members
func NewBaseCore(config map[string]string, instance_name string, map_allocate_size int, buffer_instance_allocate_size int, window *glfw.Window) *BaseCore {
	var core BaseCore

	info_file, err := os.OpenFile("info_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	error_file, err := os.OpenFile("error_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	warn_file, err := os.OpenFile("warn_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	core.core_props = config
	core.info_log = log.New(info_file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	core.error_log = log.New(error_file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	core.warn_log = log.New(warn_file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	core.instance_names = []string{instance_name}
	core.name = instance_name

	core.logical_devices = make(map[string]CoreDevice, map_allocate_size)
	core.instances = make(map[string]CoreInstance, map_allocate_size)

	core.images = make(map[string]CoreImage, buffer_instance_allocate_size)
	core.vertex_buffers = make(map[string]CoreBuffer, buffer_instance_allocate_size)
	core.indice_buffers = make(map[string]CoreBuffer, buffer_instance_allocate_size)
	core.uv_buffers = make(map[string]CoreBuffer, buffer_instance_allocate_size)
	core.color_buffers = make(map[string]CoreBuffer, buffer_instance_allocate_size)
	core.attr_buffers = make(map[string]CoreBuffer, buffer_instance_allocate_size)
	core.uniforms = make(map[string]int, buffer_instance_allocate_size)

	if window != nil && core.core_props["display"] == "true" {
		core.display = CoreDisplay{
			window: window,
		}
	}

	return &core
}

func (base *BaseCore) Release() {
	for _, inst := range base.instances {
		inst.Destroy()
	}
}

func (base *BaseCore) CreateInstance(instances []Instance) error {
	var err error

	//Get Core API Defined Wanted Layers and Extensions
	api_validation := base.GetValidationLayers()
	api_device := base.GetDeviceExtensions()
	api_instance := base.GetInstanceExtensions()
	api_required := base.display.window.GetRequiredInstanceExtensions()

	//Extension objects
	inst_ext := NewBaseInstanceExtensions(api_instance, api_required)
	layer_ext := NewBaseLayerExtensions(api_validation)

	//Create instance
	var instance vk.Instance
	var flags vk.InstanceCreateFlags
	if PlatformOS == "Darwin" {
		flags = vk.InstanceCreateFlags(0x00000001) //VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT
	} else {
		flags = vk.InstanceCreateFlags(0)
	}

	for _, ref := range instances {
		ret := vk.CreateInstance(&vk.InstanceCreateInfo{
			SType: vk.StructureTypeInstanceCreateInfo,
			PApplicationInfo: &vk.ApplicationInfo{
				SType:              vk.StructureTypeApplicationInfo,
				ApiVersion:         uint32(vk.MakeVersion(1, 1, 0)),
				ApplicationVersion: uint32(vk.MakeVersion(1, 1, 0)),
				PApplicationName:   safeString(ref.Name),
				PEngineName:        base.name + "\x00",
			},
			EnabledExtensionCount:   uint32(len(inst_ext.GetExtensions())),
			PpEnabledExtensionNames: safeStrings(inst_ext.GetExtensions()),
			EnabledLayerCount:       uint32(len(layer_ext.GetExtensions())),
			PpEnabledLayerNames:     safeStrings(layer_ext.GetExtensions()),
			Flags:                   flags,
		}, nil, &instance)

		if ret != vk.Success {
			base.error_log.Fatalf("Error creating instance with required extensions\n")
		}

		if PlatformOS == "Darwin" {
			vk.InitInstance(instance)
		}

		if ref.Selector == RENDER_INSTANCE {
			base.instances[ref.Name], err = NewCoreRenderInstance(instance, base.instance_names[0], *inst_ext, *layer_ext, api_device, &base.display)
		}

		if ref.Selector == DEVICE_INSTANCE {
			base.instances[ref.Name], err = NewCoreDeviceInstance(instance, base.instance_names[0], *inst_ext, *layer_ext, api_device)
		}

		if err != nil {
			base.error_log.Print(err)
			return err
		}

	}

	if err != nil {
		base.error_log.Print(err)
	}
	return nil
}

func (base *BaseCore) GetInstance(name string) CoreInstance {
	return base.instances[name]
}

func (base *BaseCore) GetValidationLayers() []string {

	if base.core_props["validation"] != "" {
		return []string{
			base.core_props["validation"],
		}
	} else {
		return []string{}
	}

}
func (base *BaseCore) GetDeviceExtensions() []string {
	return []string{"VK_KHR_swapchain", "VK_KHR_portability_subset", "VK_KHR_device_group"}
}

func (base *BaseCore) GetInstanceExtensions() []string {
	var darwin_extensions []string
	var other_extensions []string
	core_extensions := []string{"VK_KHR_surface", "VK_KHR_device_group_creation"}
	if PlatformOS == "Darwin" {
		darwin_extensions = []string{"VK_MVK_macos_surface", "VK_EXT_metal_surface", "VK_KHR_portability_enumeration"}
	}

	if usage := base.core_props["external"]; usage == "default" {
		other_extensions = []string{"VK_KHR_external_fence_capabilities", "VK_KHR_external_semaphore_capabilities", "VK_KHR_external_memory_capabilities"}
	}
	if debug := base.core_props["debug"]; debug == "true" {
		other_extensions = append(other_extensions, "VK_EXT_debug_report", "VK_EXT_debug_utils")
	}
	ext := append(darwin_extensions, other_extensions...)
	return append(ext, core_extensions...)
}
