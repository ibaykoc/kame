#version 410

in vec2 vOutTexCoord;
uniform sampler2D texture1;

out vec4 fOutColor;

void main() {
	fOutColor = texture(texture1, vOutTexCoord);
}