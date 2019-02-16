package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/go-gl/mathgl/mgl32"
)

// Drawer to draw something onto the screen
type Drawer struct {
	BackgroundColor      Color
	camera               Camera
	defaultShaderProgram ShaderProgram
	loadedTextureFile    map[string]uint32
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
	gl.Enable(gl.CULL_FACE)
	// version := gl.GoStr(gl.GetString(gl.VERSION))
	// fmt.Println("OpenGL initialized: version", version)

	defaultShaderProgram := createShaderProgram(
		defaultVertexShader,
		defaultFragmentShader,
		[]string{
			"m",
			"v",
			"p",
			"defaultTexture",
			"userDefinedTexture0",
			"hasTexture",
		},
	)
	camera := createCamera(mgl32.DegToRad(90))

	defaultShaderProgram.Start()
	defaultShaderProgram.SetUniformMat4F("v", camera.viewMatrix())
	defaultShaderProgram.SetUniformMat4F("p", mgl32.Perspective(camera.fov, float32(window.width)/float32(window.height), 0.1, 100))
	defaultShaderProgram.SetUniform1i("defaultTexture", 0)
	defaultShaderProgram.SetUniform1i("userDefinedTexture0", 1)
	defaultShaderProgram.Stop()

	gl.ClearColor(
		bgColor.R,
		bgColor.G,
		bgColor.B,
		bgColor.A)

	return &Drawer{
		BackgroundColor:      bgColor,
		defaultShaderProgram: defaultShaderProgram,
		camera:               camera,
		loadedTextureFile:    make(map[string]uint32),
	}, nil
}

func (d *Drawer) MoveCameraRelative(x, y, z float32) {
	d.camera.Move(x, y, z)
	d.defaultShaderProgram.Start()
	d.defaultShaderProgram.SetUniformMat4F("v", d.camera.viewMatrix())
	d.defaultShaderProgram.Stop()
}

func (d *Drawer) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Draw draw model at default position
func (d *Drawer) Draw(dm DrawableModel) {
	d.DrawAt(dm, mgl32.Translate3D(0, 0, 0))
}

// Draw0 draw model at specified position
func (d *Drawer) DrawAtPosition(dm DrawableModel, position mgl32.Vec3) {
	d.DrawAt(dm, mgl32.Translate3D(position.Elem()))
}

func (d *Drawer) DrawAtRotation(dm DrawableModel, rotation mgl32.Vec3) {
	rValue := rotation.Len()
	rAxis := rotation.Normalize()
	d.DrawAt(dm, mgl32.Translate3D(0, 0, 0).Mul4(mgl32.HomogRotate3D(rValue, rAxis)))
}
func (d *Drawer) DrawAt(dm DrawableModel, translation mgl32.Mat4) {
	d.defaultShaderProgram.Start()
	d.defaultShaderProgram.SetUniformMat4F("v", d.camera.viewMatrix())
	d.defaultShaderProgram.SetUniformMat4F("m", translation)
	dm.startDraw()
	gl.DrawElements(gl.TRIANGLES, dm.elementSize, gl.UNSIGNED_INT, gl.PtrOffset(0))
	dm.stopDraw()
	d.defaultShaderProgram.Stop()
}
func (d *Drawer) changeSize(width int32, height int32) {
	d.defaultShaderProgram.Start()
	d.defaultShaderProgram.SetUniformMat4F("p", mgl32.Perspective(d.camera.fov, float32(width)/float32(height), 0.1, 100))
	d.defaultShaderProgram.Stop()
	gl.Viewport(0, 0, width, height)
}

func (d *Drawer) dispose() {
	d.defaultShaderProgram.Dispose()
	for _, textureID := range d.loadedTextureFile {
		gl.DeleteTextures(1, &textureID)
	}
}
