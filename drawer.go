package kame

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	vertexShaderSource = `
		#version 410
		layout (location = 0) in vec3 position;

		void main() {
			gl_Position = vec4(position, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

var shaderProgramId uint32

type Drawer struct {
	BackgroundColor Color
}

func newDrawer(backgroundColor Color) (*Drawer, error) {
	bgColor := backgroundColor
	if err := gl.Init(); err != nil {
		return nil, err
	}
	// version := gl.GoStr(gl.GetString(gl.VERSION))
	// fmt.Println("OpenGL initialized: version", version)
	gl.ClearColor(
		bgColor.R,
		bgColor.G,
		bgColor.B,
		bgColor.A)

	// VERTEX
	vShaderID := gl.CreateShader(gl.VERTEX_SHADER)
	cstr, free := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vShaderID, 1, cstr, nil)
	free()
	gl.CompileShader(vShaderID)
	var success int32
	gl.GetShaderiv(vShaderID, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vShaderID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vShaderID, logLength, nil, gl.Str(log))

		fmt.Printf("failed to compile vertex shader: %v\n", log)
	}

	//FRAGMENT
	fShaderID := gl.CreateShader(gl.FRAGMENT_SHADER)
	cstr, free = gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fShaderID, 1, cstr, nil)
	free()
	gl.CompileShader(fShaderID)
	gl.GetShaderiv(fShaderID, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fShaderID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fShaderID, logLength, nil, gl.Str(log))

		fmt.Printf("failed to compile fragment shader: %v\n", log)
	}

	//PROGRAM
	shaderProgramId = gl.CreateProgram()
	gl.AttachShader(shaderProgramId, vShaderID)
	gl.AttachShader(shaderProgramId, fShaderID)
	gl.LinkProgram(shaderProgramId)

	gl.GetProgramiv(shaderProgramId, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgramId, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgramId, logLength, nil, gl.Str(log))
		fmt.Printf("failed to link shader program: %v\n", log)
	}

	// delete the shaders as they're linked into our program now and no longer necessery
	gl.DeleteShader(vShaderID)
	gl.DeleteShader(fShaderID)

	return &Drawer{
		BackgroundColor: bgColor,
	}, nil
}

func (d *Drawer) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// func (d *Drawer) DrawRect(x float32, y float32, w float32, h float32) {
// }

func (d *Drawer) Draw(model DrawableModel) {
	gl.UseProgram(shaderProgramId)

	model.vao.bind()

	for attribID := range model.vao.attributes {
		gl.EnableVertexAttribArray(attribID)
	}

	gl.DrawArrays(gl.TRIANGLES, 0, model.vertexSize)

	for attribID := range model.vao.attributes {
		gl.DisableVertexAttribArray(attribID)
	}
	model.vao.unbind()
}

func (d *Drawer) dispose() {

}
