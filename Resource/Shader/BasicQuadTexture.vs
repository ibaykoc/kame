#version 410

layout (location = 0) in vec3 inPosition;
layout (location = 1) in vec2 inTexCoord;

out vec2 vOutTexCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {
    vOutTexCoord = inTexCoord;
	gl_Position = projection * view * model * vec4(inPosition, 1.0);
}