package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/go-gl/mathgl/mgl32"
)

// KDrawer to draw something onto the screen
type KDrawer struct {
	BackgroundColor      Color
	camera               Camera
	defaultShaderProgram ShaderProgram
	loadedTextureFile    map[string]uint32
}

func newDrawer2D(backgroundColor Color) (*KDrawer, error) {
	return newDrawer(Orthographic, backgroundColor)
}

func newDrawer3D(backgroundColor Color) (*KDrawer, error) {
	return newDrawer(Perspective, backgroundColor)
}

func newDrawer(cameraType ProjectionType, backgroundColor Color) (*KDrawer, error) {
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
	var camera Camera
	if cameraType == Orthographic {
		camera = createCamera2D(50)
	} else {
		camera = createCamera3D(mgl32.DegToRad(90))
	}

	defaultShaderProgram.Start()
	defaultShaderProgram.SetUniformMat4F("v", camera.viewMatrix())
	defaultShaderProgram.SetUniformMat4F("p", camera.projectionMatrix())
	defaultShaderProgram.SetUniform1i("defaultTexture", 0)
	defaultShaderProgram.SetUniform1i("userDefinedTexture0", 1)
	defaultShaderProgram.Stop()

	gl.ClearColor(
		bgColor.R,
		bgColor.G,
		bgColor.B,
		bgColor.A)

	return &KDrawer{
		BackgroundColor:      bgColor,
		defaultShaderProgram: defaultShaderProgram,
		camera:               camera,
		loadedTextureFile:    make(map[string]uint32),
	}, nil
}

func (d *KDrawer) MoveCameraRelative(x, y, z float32) {
	d.camera.Move(x, y, z)
	d.defaultShaderProgram.Start()
	d.defaultShaderProgram.SetUniformMat4F("v", d.camera.viewMatrix())
	d.defaultShaderProgram.Stop()
}

func (d *KDrawer) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Draw draw model at default position
func (d *KDrawer) Draw(dm DrawableModel) {
	d.DrawAt(dm, mgl32.Translate3D(0, 0, 0))
}

// Draw0 draw model at specified position
func (d *KDrawer) DrawAtPosition(dm DrawableModel, position mgl32.Vec3) {
	d.DrawAt(dm, mgl32.Translate3D(position.Elem()))
}

func (d *KDrawer) DrawAtRotation(dm DrawableModel, rotation mgl32.Vec3) {
	rValue := rotation.Len()
	rAxis := rotation.Normalize()
	d.DrawAt(dm, mgl32.Translate3D(0, 0, 0).Mul4(mgl32.HomogRotate3D(rValue, rAxis)))
}
func (d *KDrawer) DrawAt(dm DrawableModel, translation mgl32.Mat4) {
	d.defaultShaderProgram.Start()
	d.defaultShaderProgram.SetUniformMat4F("v", d.camera.viewMatrix())
	d.defaultShaderProgram.SetUniformMat4F("m", translation)
	dm.startDraw()
	gl.DrawElements(gl.TRIANGLES, dm.elementSize, gl.UNSIGNED_INT, gl.PtrOffset(0))
	dm.stopDraw()
	d.defaultShaderProgram.Stop()
}
func (d *KDrawer) changeSize(width int32, height int32) {
	d.defaultShaderProgram.Start()
	d.defaultShaderProgram.SetUniformMat4F("p", d.camera.projectionMatrix())
	d.defaultShaderProgram.Stop()
	gl.Viewport(0, 0, width, height)
}

func (d *KDrawer) dispose() {
	d.defaultShaderProgram.Dispose()
	for _, textureID := range d.loadedTextureFile {
		gl.DeleteTextures(1, &textureID)
	}
}
