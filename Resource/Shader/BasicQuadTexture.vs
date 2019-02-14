#version 410

layout (location = 0) in vec3 inPosition;
layout (location = 1) in vec2 inTexCoord;
layout (location = 2) in vec3 inNormal;

out vec2 vOutTexCoord;
out vec3 vOutSurfaceNormal;
out vec3 vOutToLightVector;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;
uniform vec3 lightPosition;

void main() {
    vec4 worldPosition = model * vec4(inPosition, 1.0);
    vOutTexCoord = inTexCoord;
    vOutSurfaceNormal = (model * vec4(inNormal, 0.0)).xyz;
    vOutToLightVector = worldPosition.xyz -lightPosition ;
	gl_Position = projection * view * worldPosition;
}