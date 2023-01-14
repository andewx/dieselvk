//we will be using glsl version 4.5 syntax
#version 450

layout(location = 0) in vec3 inPosition;
layout(location = 1) out vec3 fragColor;

layout(std140, binding = 0) uniform UniformBufferObject{
 	mat4 model;
    mat4 view;
    mat4 proj;
} ubo;

layout (push_constant) uniform constants{
	float delta;
}PushConstants;

void main()
{

	mat4 mvp =  ubo.proj * ubo.view * ubo.model;

	vec3 colors[3] = vec3[3](
		vec3(1.0,0.0,PushConstants.delta),
		vec3(PushConstants.delta,1.0,0.0),
		vec3(0.0,PushConstants.delta,1.0 )
	);


	gl_Position =   mvp * vec4(inPosition,1.0);
    fragColor = colors[gl_VertexIndex];
}