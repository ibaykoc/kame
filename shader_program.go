package kame

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type ShaderProgram struct {
	id               uint32
	defaultTextureID uint32
	uniforms         map[string]int32
}

func createShaderProgram(vertexShaderSource string, fragmentShaderSource string, uniforms []string) ShaderProgram {
	vsID := loadShader(vertexShaderSource, gl.VERTEX_SHADER)
	fsID := loadShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	shaderProgramID := gl.CreateProgram()
	gl.AttachShader(shaderProgramID, vsID)
	gl.AttachShader(shaderProgramID, fsID)
	gl.LinkProgram(shaderProgramID)

	var success int32
	gl.GetProgramiv(shaderProgramID, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgramID, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgramID, logLength, nil, gl.Str(log))
		fmt.Printf("failed to link shader program: %v\n", log)
	}

	// delete the shaders as they're linked into our program now and no longer necessery
	gl.DeleteShader(fsID)
	gl.DeleteShader(vsID)

	// Get uniforms
	_uniforms := make(map[string]int32)
	for _, uniform := range uniforms {
		uniCstr := gl.Str(uniform + "\x00")
		_uniforms[uniform] = gl.GetUniformLocation(shaderProgramID, uniCstr)

	}
	return ShaderProgram{
		id:               shaderProgramID,
		defaultTextureID: loadDefaultTexture(),
		uniforms:         _uniforms,
	}
}

func (p *ShaderProgram) setUniform1i(name string, value int32) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	gl.Uniform1i(uniLocation, value)
}

func (p *ShaderProgram) setUniform1F(name string, value float32) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	gl.Uniform1f(uniLocation, value)
}

func (p *ShaderProgram) setUniform3F(name string, v0 float32, v1 float32, v2 float32) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	gl.Uniform3f(uniLocation, v0, v1, v2)
}
func (p *ShaderProgram) setUniformMat4F(name string, value mgl.Mat4) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	m4 := [16]float32(value)
	gl.UniformMatrix4fv(uniLocation, 1, false, &m4[0])
}

func (p *ShaderProgram) start() {
	gl.UseProgram(p.id)
}
func (p *ShaderProgram) stop() {
	gl.UseProgram(0)
}
func (p *ShaderProgram) dispose() {
	gl.DeleteProgram(p.id)
}

// func loadShader(source string, shaderType uint32) uint32 {
// 	shaderSource := source + "\x00"
// 	shaderID := gl.CreateShader(shaderType)
// 	cstr, free := gl.Strs(shaderSource)
// 	gl.ShaderSource(shaderID, 1, cstr, nil)
// 	free()
// 	gl.CompileShader(shaderID)
// 	var success int32
// 	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &success)
// 	if success == gl.FALSE {
// 		var logLength int32
// 		gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLength)

// 		log := strings.Repeat("\x00", int(logLength+1))
// 		gl.GetShaderInfoLog(shaderID, logLength, nil, gl.Str(log))

// 		fmt.Printf("failed to compile shader\nTYPE: %v\nLOG: %v\nSOURCE: %v\n", shaderType, log, shaderSource)
// 	}
// 	return shaderID
// }
