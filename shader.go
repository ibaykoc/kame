package kame

const (
	defaultVertexShader = `
	#version 410
	layout (location = 0) in vec3 aPos;
	layout (location = 1) in vec2 aUV;

	out vec2 UV;

	uniform mat4 m;
	uniform mat4 v;
	uniform mat4 p;

	void main() {
		UV = aUV;
		gl_Position = p * v * m * vec4(aPos, 1.0);
	}`

	defaultFragmentShader = `
	#version 410
	out vec4 color;

	in vec2 UV;

	uniform sampler2D defaultTexture;
	uniform sampler2D userDefinedTexture0;

	uniform float hasTexture;

	void main() {
		vec4 defaultTextureColor = texture(defaultTexture, UV);
		vec4 userDefinedTexture0Color = texture(userDefinedTexture0, UV);
		color = mix(defaultTextureColor, userDefinedTexture0Color, hasTexture);
		// color = vec4(hasTexture,0,0,1);
	}`
)
