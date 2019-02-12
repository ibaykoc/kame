package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Drawer struct {
	basicShaderProgram ShaderProgram
	BackgroundColor    Color
}

func newDrawer(backgroundColor Color) (*Drawer, error) {
	bgColor := backgroundColor
	if err := gl.Init(); err != nil {
		return nil, err
	}
	// version := gl.GoStr(gl.GetString(gl.VERSION))
	// fmt.Println("OpenGL initialized: version", version)

	basicShaderProgram := createShaderProgram(
		"Shader/BasicVertexShader.glsl",
		"Shader/BasicFragmentShader.glsl")

	gl.ClearColor(
		bgColor.R,
		bgColor.G,
		bgColor.B,
		bgColor.A)

	return &Drawer{
		BackgroundColor:    bgColor,
		basicShaderProgram: basicShaderProgram,
	}, nil
}

func (d *Drawer) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// func (d *Drawer) DrawRect(x float32, y float32, w float32, h float32) {
// }

func (d *Drawer) Draw(model DrawableModel) {
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	d.basicShaderProgram.Start()

	model.vao.bind()

	for attribID := range model.vao.attributes {
		gl.EnableVertexAttribArray(attribID)
	}

	gl.DrawElements(gl.TRIANGLES, model.vertexSize, gl.UNSIGNED_INT, nil)

	for attribID := range model.vao.attributes {
		gl.DisableVertexAttribArray(attribID)
	}
	model.vao.unbind()
	d.basicShaderProgram.Stop()
}

func (d *Drawer) changeSize(width int32, height int32) {
	gl.Viewport(0, 0, width, height)
}

func (d *Drawer) dispose() {
	d.basicShaderProgram.Dispose()
}
