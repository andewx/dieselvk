package dieselvk


//Descriptor sets and descriptor pools file

/*
	ubo_layout := vk.DescriptorSetLayoutBinding{}
	ubo_layout.Binding = core.location
	ubo_layout.DescriptorCount = 1
	ubo_layout.DescriptorType = vk.DescriptorTypeUniformBuffer
	ubo_layout.StageFlags = core.stage_flags
	ubo_layout.PImmutableSamplers = nil

	bindings := []vk.DescriptorSetLayoutBinding{ubo_layout}

	ubo_create := vk.DescriptorSetLayoutCreateInfo{}
	ubo_create.SType = vk.StructureTypeDescriptorSetLayoutCreateInfo
	ubo_create.BindingCount = 1
	ubo_create.PBindings = bindings

	if vk.CreateDescriptorSetLayout(handle, &ubo_create, nil, &core.layout) != vk.Success {
		Fatal(fmt.Errorf("Failed to create uniform buffer object"))
	}
	*/