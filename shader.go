package kame

const (
	default2DVertexShader = `
	#version 410
	layout (location = 0) in vec4 vPos;
	layout (location = 1) in vec2 vUV;
	layout (location = 2) in mat4 iModelMatrix;

	out vec2 UV;

	uniform mat4 v;
	uniform mat4 p;

	void main() {
		UV = vUV;
		gl_Position = p * v * iModelMatrix * vPos;
		// gl_Position = iModelMatrix[0];
		// gl_Position = vPos;
	}`

	default2DFragmentShader = `
	#version 410
	in vec2 UV;
	
	uniform sampler2D texture0;
	uniform vec3 tintColor;
	
	out vec4 color;
	
	void main()
	{    
		// color = vec4(1,0,0,1);
		color = vec4(tintColor, 1.0) * texture(texture0, UV);
		if (color.a == 0.0){
			discard;
		}
	}`

	default3DVertexShader = `
	#version 410
	layout (location = 0) in vec4 vPos;
	layout (location = 1) in vec2 vUV;
	layout (location = 2) in vec3 vNormal;
	layout (location = 3) in mat4 iModelMatrix;

	out vec2 UV;

	uniform mat4 v;
	uniform mat4 p;

	void main() {
		UV = vUV;
		gl_Position = p * v * iModelMatrix * vPos;
		// gl_Position = iModelMatrix[0];
		// gl_Position = vPos;
	}`

	default3DFragmentShader = `
	#version 410

	in vec2 UV;
	
	uniform sampler2D texture0;
	uniform vec3 tintColor;

	out vec4 color;

	void main() {
		color = texture(texture0, UV);
		if (color.a == 0.0){
			discard;
		}
	}`
)
