#version 410

in vec2 vOutTexCoord;
in vec3 vOutSurfaceNormal;
in vec3 vOutToLightVector;

out vec4 fOutColor;

uniform sampler2D texture1;
uniform vec3 lightColor;

void main() {

	vec3 unitNormal = normalize(vOutSurfaceNormal);
	vec3 unitToLightVector = normalize(vOutToLightVector);

	float nDotl = dot(unitToLightVector, unitNormal);
	float brightness = max(nDotl, 0.0);
	vec3 diffuse = brightness * lightColor;

	fOutColor = vec4(diffuse, 1.0) * texture(texture1, vOutTexCoord);
}