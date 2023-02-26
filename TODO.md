#TODO Alpha Milestone

> Lists developer todo categories to maintain continuity with specific tasks

1. Integrate JSON Scehma
 - ~~Create and Generate JSON Schema~~
 - ~~Marshall and Unmarshall JSON Schema in GO~~
 - ~~Provide Generative Method to Capture Vulkan String Enumerations from JSON files into integers for driver support~~
 - Driver Properties Configuration
 - Swapchain Configuration
 - Framebuffer Configuration General
 - Pipleine Configuration
 - Renderpass Configuration
 - Queue Configuration 
 - Buffer Setup Configuration
 - Allocation Configuration
 - Shader Configuration
 - Shader Stage Layout Configuration
 - Uniform Configuration
 - Sampler Configuration


 2. Texture Support and Texture Mapping
 - Add in texture support, image loading, mip-mapping and integration into uniform sampler shaders



 **Comments**
> These are the last items before an alpha release could be considered working and overall the implementation could take a while before all the configuration features are implemented. One thing though is that making our entire "pipeline" configurable makes the process of starting up a vulkan rendering pipeline user-friendly as hell and will provide the interface to work with other rendering libraries as well.

> JSON file schemas should also help with identifying potential bugs or interelated issues between parts. 
