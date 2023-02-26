package dieselvk

import (
	"io/ioutil"
	"log"
	"os"

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

// Global Constants
const (
	RENDER_INSTANCE  = 0
	DEVICE_INSTANCE  = 1
	COMPUTE_INSTANCE = 2
	MIN_ALLOC        = 10
)

// Base DieselVK Core vulkan manager with GLFW native host management
// Core structure properties are private members to enforce future interface
// compliance with outside packages. The Vulkan core manager manages the availability
// of devices and capabilities to enfore instance creation and management. Also holds
// global type information which could be useful to multiple vulkan instances in an application
// which includes buffers and textures. Light objects in Vulkan do not always warrant a Core Abstraction
type BaseCore struct {

	//Core Implementation Context Properties
	display CoreDisplay
	Name    string

	//List of device bindings
	logical_devices map[string]CoreDevice

	//Single instance handle
	instance CoreInstance
}

// Instanitates a new core context allocation sizes, default allocation prevents buffer copies but is just used to instantiate map members
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

	//Open VLK Json Configuration file into global vars {Vlk, Dictionary}
	if contents, file_err := ioutil.ReadFile(json_file); file_err != nil {
		ErrorLog.Fatal(file_err)
	} else {
		json_err := Vlk.UnmarshalJSON(contents)
		if json_err != nil {
			ErrorLog.Fatal(json_err)
		}
	}

	core.Name = instance_name
	core.logical_devices = make(map[string]CoreDevice, MIN_ALLOC)

	//Setup the window structs
	if window != nil && Vlk.Config.Display == "true" {
		core.display = CoreDisplay{
			window: window,
		}
	}

	return &core
}

func (base *BaseCore) GetExtensions() []string {
	return append(Vlk.Config.CoreExtensions, Vlk.Config.UserExtensions...)
}

func (base *BaseCore) Release() {
	base.instance.Destroy()
}

func (base *BaseCore) CreateInstance() error {
	var err error

	//Create instance
	var instance vk.Instance
	var flags vk.InstanceCreateFlags
	if PlatformOS == "Darwin" {
		flags = vk.InstanceCreateFlags(0x00000001) //VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT
	} else {
		flags = vk.InstanceCreateFlags(0)
	}

	//Validate Extensions/Devices/Layers
	var missing []string
	missing = ValidateExtensions(Vlk)
	missing = append(missing, ValidateLayers(Vlk)...)

	//Handle missing Extensions Case
	if len(missing) > 0 {
		var accum string
		for _, str := range missing {
			accum += str + "\n"
		}
		ErrorLog.Fatalf("Could not instantiate a device the following extensions or errors occured\n%s", accum)
	}

	ret := vk.CreateInstance(&vk.InstanceCreateInfo{
		SType: vk.StructureTypeInstanceCreateInfo,
		PApplicationInfo: &vk.ApplicationInfo{
			SType:              vk.StructureTypeApplicationInfo,
			ApiVersion:         uint32(vk.MakeVersion(1, 1, 0)),
			ApplicationVersion: uint32(vk.MakeVersion(1, 1, 0)),
			PApplicationName:   safeString(base.Name),
			PEngineName:        base.Name + "\x00",
		},
		EnabledExtensionCount:   uint32(len(base.GetExtensions())),
		PpEnabledExtensionNames: safeStrings(base.GetExtensions()),
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

	base.instance, err = NewCoreRenderInstance(instance, base.Name, &base.display)

	if err != nil {
		ErrorLog.Fatal(err)
		return err
	}

	if err != nil {
		ErrorLog.Fatal(err)
	}
	return nil
}

func (base *BaseCore) GetInstance() CoreInstance {
	return base.instance
}
