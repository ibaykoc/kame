package kame

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Drawer struct {
	basicShaderProgram ShaderProgram
	Camera             Camera
	BackgroundColor    Color
}

func newDrawer(window *Window, backgroundColor Color) (*Drawer, error) {
	bgColor := backgroundColor
	if err := gl.Init(); err != nil {
		return nil, err
	}
	// Enable alpha blending
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Enable(gl.DEPTH_TEST)
	// gl.Enable(gl.CULL_FACE)
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL initialized: version", version)

	basicShaderProgram := createShaderProgram(
		"Shader/BasicQuadTexture.vs",
		"Shader/BasicQuadTexture.fs",
		[]string{
			"model",
			"view",
			"projection",
		},
	)
	basicShaderProgram.Start()

	pMat := mgl.Perspective(mgl.DegToRad(45), float32(window.width)/float32(window.height), 0.1, 1000)
	basicShaderProgram.SetUniformMat4F("projection", pMat)
	basicShaderProgram.Stop()

	gl.ClearColor(
		bgColor.R,
		bgColor.G,
		bgColor.B,
		bgColor.A)

	return &Drawer{
		BackgroundColor:    bgColor,
		Camera:             CreateCamera(),
		basicShaderProgram: basicShaderProgram,
	}, nil
}

func (d *Drawer) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// func (d *Drawer) DrawRect(x float32, y float32, w float32, h float32) {
// }

func (d *Drawer) Draw(e Entity) {
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE
	d.basicShaderProgram.Start()
	d.basicShaderProgram.SetUniformMat4F("model", e.modelMatrix())
	d.basicShaderProgram.SetUniformMat4F("view", d.Camera.viewMatrix())
	e.DrawableModel.vao.bind()

	for attrID := uint32(0); attrID < e.DrawableModel.vao.attributeSize; attrID++ {
		gl.EnableVertexAttribArray(attrID)
	}
	gl.BindTexture(gl.TEXTURE_2D, e.DrawableModel.textureID)
	gl.DrawElements(gl.TRIANGLES, e.DrawableModel.vertexSize, gl.UNSIGNED_INT, gl.PtrOffset(0))

	for attrID := uint32(0); attrID < e.DrawableModel.vao.attributeSize; attrID++ {
		gl.DisableVertexAttribArray(attrID)
	}
	e.DrawableModel.vao.unbind()
	d.basicShaderProgram.Stop()
}

func (d *Drawer) changeSize(width int32, height int32) {
	gl.Viewport(0, 0, width, height)
}

func (d *Drawer) dispose() {
	d.basicShaderProgram.Dispose()
}
