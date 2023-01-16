package dieselvk

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/andewx/dieselvk/json"
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

type Instance struct {
	Name     string
	Selector int
	Gpu_id   int
}

type Resource interface {
	Release()
}

//Global Constants
const (
	RENDER_INSTANCE  = 0
	DEVICE_INSTANCE  = 1
	COMPUTE_INSTANCE = 2
	MIN_ALLOC        = 10
)

//Base DieselVK Core vulkan manager with GLFW native host management
//Core structure properties are private members to enforce future interface
//compliance with outside packages. The Vulkan core manager manages the availability
//of devices and capabilities to enfore instance creation and management. Also holds
//global type information which could be useful to multiple vulkan instances in an application
//which includes buffers and textures. Light objects in Vulkan do not always warrant a Core Abstraction
type BaseCore struct {

	//Core Implementation Context Properties
	display CoreDisplay
	name    string

	//Map string id & tagging
	instance_names []string

	//List of device bidings
	logical_devices map[string]CoreDevice

	//Per Instance/Device Handles where key is the instance global id key used for accessing other held resources
	instances map[string]CoreInstance //Key: (Instance_Name) Value: Vulkan Instance

	BaseExtensionLoader *BaseInstanceExtensions
	BaseDeviceLoader    *BaseDeviceExtensions
	BaseLayerLoader     *BaseLayerExtensions
}

//Instanitates a new core context allocation sizes, default allocation prevents buffer copies but is just used to instantiate map members
func NewBaseCore(json_file string, instance_name string, window *glfw.Window) *BaseCore {
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

	//Initialize Globals
	InfoLog = log.New(info_file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(error_file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLog = log.New(warn_file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	//Derive and potentially verify the configuration file
	Dictionary = json.InitDictionary()
	if contents, file_err := ioutil.ReadFile(json_file); file_err != nil {
		ErrorLog.Fatal(file_err)
	} else {
		Vlk.UnmarshalJSON(contents)
	}

	core.instance_names = []string{instance_name}
	core.name = instance_name

	core.logical_devices = make(map[string]CoreDevice, MIN_ALLOC)
	core.instances = make(map[string]CoreInstance, MIN_ALLOC)

	if window != nil && Vlk.Config.Display == "true" {
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
			EnabledExtensionCount:   uint32(len(base.GetInstanceExtensions())),
			PpEnabledExtensionNames: safeStrings(base.GetInstanceExtensions()),
			EnabledLayerCount:       uint32(len(Vlk.Config.VulkanLayers)),
			PpEnabledLayerNames:     safeStrings(Vlk.Config.VulkanLayers),
			Flags:                   flags,
		}, nil, &instance)

		if ret != vk.Success {
			ErrorLog.Fatalf("Error creating instance with required extensions\n")
		}

		if PlatformOS == "Darwin" {
			vk.InitInstance(instance)
		}

		if ref.Selector == RENDER_INSTANCE {
			base.instances[ref.Name], err = NewCoreRenderInstance(instance, base.instance_names[0], &base.display)
		}

		if ref.Selector == DEVICE_INSTANCE {
			//	base.instances[ref.Name], err = NewCoreDeviceInstance(instance, base.instance_names[0], *inst_ext, *layer_ext, api_device)
		}

		if err != nil {
			ErrorLog.Print(err)
			return err
		}

	}

	if err != nil {
		ErrorLog.Print(err)
	}
	return nil
}

func (base *BaseCore) GetInstance(name string) CoreInstance {
	return base.instances[name]
}

func (base *BaseCore) GetValidationLayers() []string {
	return Vlk.Config.VulkanLayers
}

func (base *BaseCore) GetInstanceExtensions() []string {
	var darwin_extensions []string
	core_extensions := Vlk.Config.CoreExtensions

	if PlatformOS == "Darwin" {
		darwin_extensions = []string{"VK_MVK_macos_surface", "VK_EXT_metal_surface", "VK_KHR_portability_enumeration"}
	}

	ext := append(core_extensions, darwin_extensions...)
	return ext
}
