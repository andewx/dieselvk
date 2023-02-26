# DieselVK

![diselvk logo](logo/dslvk.jpg?raw=true "dieslvk")

-`dieselvk.Core`

> **Attention** currently conducting JSON configuration refactor in *alpha* branch. 

- Diesel Vulkan initializes entire Vulkan Instance from a Vlk Json Schema. See the [schema](json/vlk_schema.json) document and the [example](json/vlk_example.json) for examples on how to set up a configured instance. A global `Vlk` variable holds the configuration values.

- Global `Dictionary` variable also is used to convert the Json configuration file stored strings which represent Vulkan Enumerations into ints. To use this value for example we can call

```
my_int := Dictionary.Get(Vlk.Piplelines.Topology)
```

- For now we haven't yet implemented general struct encoding/decoding into byte arrays in a way that can be used by the cgo backend so all data passed to Vulkan in buffers should be in the form of golang arrays or single scalars via `Ptr()` utility function. For example `[]float32, []int32...` Once everything is more or less in place we will revisit these issues.

>Please be aware that DieselVK is still in a pre-alpha build and therefore may be considered highly unstable. It is recommended that users only **fork** this repository at the moment.

> Additionally documentation is very limited but hopefully is taken care of after the first Alpha release.

---------------------

## API Description

`dieselvk.BaseCore`

- Application entry point and manages the creation of Vulkan instances and any neccessary setup or configuration.

-`dieselvk.CoreRenderInstance`

  - Core render instance hosts the entire GPU rendering driver context and is the lifetime instance for a Vulkan program. The instance host all other peer modules as resources as they are needed and is responsible for the integration of those resources into rendering operations. 

  - Much of what the CoreRenderInstance does will be dependent on the Configuration values in the supplied Json Schema file. *Be prepared that certain configurations may be invalid without a specific error or warning unless validaiton layers exist*.

-`dieselvk.CoreQueue`- Provides high level queue operation control and manages queue operations. Provides reliable access to reproducible queues, queries their states, and list their properties for the underlying `Core/Instance` implementations

-`dieselvk.CorePool`- Provides high level application pool state tracking and attached command buffers to their own pools. Tracks lifetime of pools, links them to their physical devices and maintains command pools as thread specific objects for submitting work. Core Pools can maintain their own semaphores and tracks the command buffers as submit blocking or not for that specifc instance. Internal buffers can be retrieved and created as needed.

-`dieselvk.CoreDevice`- Maintians linkage of physical device to logical device. List properites and request and tracks device features/memory/capabilities. Delivers core pools and provides device linkage for multi-gpu usage.

-`dieselvk.CoreProgram`- Maintains list of descriptor sets for single shaders and provides high level buffer linkage data to individual shader programs.

-`dieselvk.CoreShader`- Provides a shader set manager for the core instance. List individual shaders and linkages to global buffer data

-`dieselvk.CorePipeline`- Assembles pipeline state info context holding and links with core program objects.

-`dieselvk.CoreSwapchain`- Creates swapchain instance + associated swap chain images. Holds swapchain image view refences. Holds and request Displace surfaces if requested. Manages swap chain fencing and semaphores for frame requests. Provides framebuffer references with default depth + color attachments. Holds desired VSYNC rates if desired.

-`dieselvk.CoreAllocator` - This provides a sub-allocator and resource management object to the instance. Allocations can be done with `Allocate()` and `Map()` functions which sub allocate Vulkan memory, has the benefit of allowing for GPU memory tracking and memory management techniques. 

-`dieselvk.CoreBuffers` - Provides High level buffer allocation routines and allows gathering references for shader/pipelines.

-`dieselvk.CoreImage` - Provides high level image allocation, attachment, image view creation, and GPU formatting issues as well as host/gpu 
communication.

-`dieselvk.CoreDisplay` - Manages screen device pixel format and other rendering format issues.


> **Attention**, I am looking for community volunteers and team-mates who can take this project to the next level and also help with identifying and resolving any bugs and issues. If you would like to join the team on this project please let me know and we will start including all the licenses etc.

# Future

> As I'm nearing the completion of the Alpha build here, which is the JSON component configuration refactor I wanted to look ahead and summarize the milestones we have in store. JSON configuration does constistute a major refactor as it affects all modules

**Bravo**
 - Dynamic Buffers / Sampler Support

**Charlie**
 - Multiple Renderpass / Compute Shader integration Testing 

**Delta**
 - Multi-Threaded Instance Support / API Documentation / Device Testing / CI Builds / API Tutorial / Mult-GPU Support Testing

**Echo**
 - Native GUI Integration Support / Raycasting Support / Electron Based Utility Workflow Applications / Casting components into submodules / Shader reflection

 If these milestones are met I will likely close the repository to be fork only. 