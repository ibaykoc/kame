package kame

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type kshaderID uint32

type kshader struct {
	id       kshaderID
	uniforms map[string]int32
}

func createkshader(vertexShaderSource string, fragmentShaderSource string, uniforms []string) kshader {
	vsID := loadShader(vertexShaderSource, gl.VERTEX_SHADER)
	fsID := loadShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	shaderID := gl.CreateProgram()
	gl.AttachShader(shaderID, vsID)
	gl.AttachShader(shaderID, fsID)
	gl.LinkProgram(shaderID)

	var success int32
	gl.GetProgramiv(shaderID, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderID, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderID, logLength, nil, gl.Str(log))
		fmt.Printf("failed to link shader program: %v\n", log)
	}

	// delete the shaders as they're linked into our program now and no longer necessery
	gl.DeleteShader(fsID)
	gl.DeleteShader(vsID)

	// Get uniforms
	_uniforms := make(map[string]int32)
	for _, uniform := range uniforms {
		uniCstr := gl.Str(uniform + "\x00")
		_uniforms[uniform] = gl.GetUniformLocation(shaderID, uniCstr)

	}
	return kshader{
		id:       kshaderID(shaderID),
		uniforms: _uniforms,
	}
}

func (p *kshader) setUniform1i(name string, value int32) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	gl.Uniform1i(uniLocation, value)
}

func (p *kshader) setUniform1F(name string, value float32) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	gl.Uniform1f(uniLocation, value)
}

func (p *kshader) setUniform3F(name string, v0 float32, v1 float32, v2 float32) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	gl.Uniform3f(uniLocation, v0, v1, v2)
}
func (p *kshader) setUniformMat4F(name string, value mgl.Mat4) {
	uniLocation, found := p.uniforms[name]
	if !found {
		panic(fmt.Errorf("Uniform (%v) not found", name))
	}
	m4 := [16]float32(value)
	gl.UniformMatrix4fv(uniLocation, 1, false, &m4[0])
}

func (p *kshader) use() {
	gl.UseProgram(uint32(p.id))
}
func (p *kshader) unuse() {
	gl.UseProgram(0)
}
func (p *kshader) dispose() {
	gl.DeleteProgram(uint32(p.id))
}
func loadShader(source string, shaderType uint32) uint32 {
	shaderSource := source + "\x00"
	shaderID := gl.CreateShader(shaderType)
	cstr, free := gl.Strs(shaderSource)
	gl.ShaderSource(shaderID, 1, cstr, nil)
	free()
	gl.CompileShader(shaderID)
	var success int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderID, logLength, nil, gl.Str(log))

		fmt.Printf("failed to compile shader\nTYPE: %v\nLOG: %v\nSOURCE: %v\n", shaderType, log, shaderSource)
	}
	return shaderID
}
