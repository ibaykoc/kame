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

func newDrawer(backgroundColor Color) (*Drawer, error) {
	bgColor := backgroundColor
	if err := gl.Init(); err != nil {
		return nil, err
	}
	// Enable alpha blending
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

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

	gl.ClearColor(
		bgColor.R,
		bgColor.G,
		bgColor.B,
		bgColor.A)

	return &Drawer{
		BackgroundColor:    bgColor,
		Camera:             CreateCamera(0, 0, -3),
		basicShaderProgram: basicShaderProgram,
	}, nil
}

func (d *Drawer) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// func (d *Drawer) DrawRect(x float32, y float32, w float32, h float32) {
// }

func (d *Drawer) Draw(model DrawableModel) {
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE
	d.basicShaderProgram.Start()
	mMat := mgl.Translate3D(0, 0, 0)
	mMat = mMat.Mul4(mgl.HomogRotate3DZ(0))
	mMat = mMat.Mul4(mgl.Scale3D(1, 1, 1))
	d.basicShaderProgram.SetUniformMat4F("model", mMat)
	d.basicShaderProgram.SetUniformMat4F("view", d.Camera.viewMatrix())
	pMat := mgl.Perspective(mgl.DegToRad(45), 540/480, 0.1, 100)
	d.basicShaderProgram.SetUniformMat4F("projection", pMat)
	model.vao.bind()

	for attrID := uint32(0); attrID < model.vao.attributeSize; attrID++ {
		gl.EnableVertexAttribArray(attrID)
	}
	gl.BindTexture(gl.TEXTURE_2D, model.textureID)
	gl.DrawElements(gl.TRIANGLES, model.vertexSize, gl.UNSIGNED_INT, gl.PtrOffset(0))

	for attrID := uint32(0); attrID < model.vao.attributeSize; attrID++ {
		gl.DisableVertexAttribArray(attrID)
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
