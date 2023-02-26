package dieselvk

import (
	"github.com/andewx/dieselvk/json"
	vk "github.com/vulkan-go/vulkan"
)

//Just make extensions checker a global thing because this is confusing for the config
//If extension is not available it is returned in the missing return string
func ValidateExtensions(vlk json.Vlk) []string {
	missing := []string{}

	instance_extensions, err := InstanceExtensions()

	if err != nil {
		return []string{"Error Loading Instance or Device Extensions"}
	}

	user_extensions := append(vlk.Config.CoreExtensions, vlk.Config.UserExtensions...)

	for _, ext := range user_extensions {
		found := false
		for _, gext := range instance_extensions {
			if ext == gext {
				found = true
			}
		}
		if !found {
			missing = append(missing, ext)
		}
	}

	return missing

}

func ValidateDeviceExtensions(vlk json.Vlk, gpu vk.PhysicalDevice) []string {
	missing := []string{}

	device_extensions, err_dev := DeviceExtensions(gpu)

	if err_dev != nil {
		return []string{"Error Loading Device Extensions"}
	}

	user_extensions := append(vlk.Config.CoreExtensions, vlk.Config.UserExtensions...)

	for _, ext := range user_extensions {
		found := false
		for _, gext := range device_extensions {
			if ext == gext {
				found = true
			}
		}
		if !found {
			missing = append(missing, ext)
		}
	}

	return missing
}

func ValidateLayers(vlk json.Vlk) []string {
	missing := []string{}

	platform_layers, err := ValidationLayers()

	if err != nil {
		return []string{"Error getting platform layers"}
	}

	for _, ext := range vlk.Config.VulkanLayers {
		found := false
		for _, gext := range platform_layers {
			if ext == gext {
				found = true
			}
		}
		if !found {
			missing = append(missing, ext)
		}
	}

	return missing
}
