package kame

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type ShaderProgram struct {
	id uint32
}

func createShaderProgram(vertexFilePath string, fragmetFilePath string) ShaderProgram {
	vsID := loadShader(vertexFilePath, gl.VERTEX_SHADER)
	fsID := loadShader(fragmetFilePath, gl.FRAGMENT_SHADER)
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
	return ShaderProgram{
		id: shaderProgramID,
	}
}

func (p *ShaderProgram) Start() {
	gl.UseProgram(p.id)
}
func (p *ShaderProgram) Stop() {
	gl.UseProgram(0)
}
func (p *ShaderProgram) Dispose() {
	gl.DeleteProgram(p.id)
}
func loadShader(filePath string, shaderType uint32) uint32 {
	shaderSourceByte, err := Resource.Find(filePath)
	if err != nil {
		panic(err)
	}
	shaderSource := string(shaderSourceByte) + "\x00"
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

		fmt.Printf("failed to compile shader: \n%v\n%v", log, shaderSource)
	}
	return shaderID
}
