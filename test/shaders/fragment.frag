#version 450

//output write
layout (location = 0) out vec4 outColor;
layout (location = 1) in vec3 fragColor;

 layout(std140, binding = 0) uniform UniformBufferObject{
  mat4 model;
    mat4 view;
    mat4 proj;
} ubo;


void main()
{
	outColor = vec4(fragColor,1.0);
}
