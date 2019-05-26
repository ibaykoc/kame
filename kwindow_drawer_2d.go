package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type KwindowDrawer2DController struct {
	KwindowDrawerController
	kwindowDrawer2D *kwindowDrawer2D
}

type KdrawerCamera2DController struct {
	camera *kdrawerCamera2D
}

func (wdCon KwindowDrawer2DController) StoreMesh(positions []float32, uvs []float32, elements []uint32) (kmeshID, error) {
	mesh, err := newkmeshPosUV(positions, uvs, elements)
	if err != nil {
		return kmeshID(0), err
	}
	wdCon.kwindowDrawer.kmeshes[mesh.id] = mesh
	return mesh.id, nil
}

func (wdCon KwindowDrawer2DController) StoreTintColor(color Kcolor) kcolorID {
	colID := kcolorID(len(wdCon.kwindowDrawer2D.tintColors))
	wdCon.kwindowDrawer2D.tintColors[colID] = color
	return colID
}

func (wdCon KwindowDrawer2DController) Camera() KdrawerCamera2DController {
	return KdrawerCamera2DController{
		camera: wdCon.kwindowDrawer.kdrawerCamera.(*kdrawerCamera2D),
	}
}

func (camCon KdrawerCamera2DController) Frustum() Kfrustum {
	return camCon.camera.frustum()
}

type kwindowDrawer2D struct {
	kwindowDrawer
	tintColors map[kcolorID]Kcolor
	batch      map[kshaderID]map[kmeshID]map[KtextureID]map[kcolorID][]mgl32.Mat4
}

func newKwindowDrawer2D(config kwindowDrawer2DBuilder) (kwindowDrawer2D, error) {
	if err := gl.Init(); err != nil {
		return kwindowDrawer2D{}, err
	}
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	debug.pf("2D OpenGL initialized: version %s\n", gl.GoStr(gl.GetString(gl.VERSION)))

	kshaders := make(map[kshaderID]*kshader)

	defaultShader := createkshader(
		default2DVertexShader,
		default2DFragmentShader,
		[]string{
			"v", "p",
			"texture0",
			"tintColor",
		},
	)

	camera := newkdrawerCamera2D(config.windowWidth, config.windowHeight, config.ppu)

	defaultShader.use()
	defaultShader.setUniformMat4F("v", camera.viewMatrix())
	defaultShader.setUniformMat4F("p", camera.projectionMatrix())
	defaultShader.setUniform1i("texture0", 0)
	defaultShader.unuse()

	kshaders[kshaderID(defaultShader.id)] = &defaultShader
	// gl.Viewport(0, 0, int32(config.windowWidth), int32(config.windowHeight))
	gl.ClearColor(config.backgroundColor.RGBA())

	kwd := kwindowDrawer{
		backgroundColor:  config.backgroundColor,
		defaultkshaderID: kshaderID(defaultShader.id),
		kdrawerCamera:    &camera,
		kshaders:         kshaders,
		kmeshes:          make(map[kmeshID]kmesh),
		ktextures:        make(map[KtextureID]ktexture),
	}

	return kwindowDrawer2D{
		kwindowDrawer: kwd,
		tintColors:    make(map[kcolorID]Kcolor),
		batch:         make(map[kshaderID]map[kmeshID]map[KtextureID]map[kcolorID][]mgl32.Mat4),
	}, nil
}

func (d *kwindowDrawer2D) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (d *kwindowDrawer2D) AppendDrawable(kdrawable Kdrawable, translation mgl32.Mat4) {
	dw := kdrawable.(Kdrawable2d)
	if _, shaderIDHasAdded := d.batch[dw.ShaderID]; !shaderIDHasAdded {
		d.batch[dw.ShaderID] = make(map[kmeshID]map[KtextureID]map[kcolorID][]mgl32.Mat4)
	}
	if _, meshIDHasAdded := d.batch[dw.ShaderID][dw.MeshID]; !meshIDHasAdded {
		d.batch[dw.ShaderID][dw.MeshID] = make(map[KtextureID]map[kcolorID][]mgl32.Mat4)
	}
	if _, textureIDHasAdded := d.batch[dw.ShaderID][dw.MeshID][dw.TextureID]; !textureIDHasAdded {
		d.batch[dw.ShaderID][dw.MeshID][dw.TextureID] = make(map[kcolorID][]mgl32.Mat4)
	}
	if _, tintColorIDHasAdded := d.batch[dw.ShaderID][dw.MeshID][dw.TextureID][dw.TintColorID]; !tintColorIDHasAdded {
		d.batch[dw.ShaderID][dw.MeshID][dw.TextureID][dw.TintColorID] = []mgl32.Mat4{}
	}
	d.batch[dw.ShaderID][dw.MeshID][dw.TextureID][dw.TintColorID] = append(d.batch[dw.ShaderID][dw.MeshID][dw.TextureID][dw.TintColorID], translation)
}

func (d *kwindowDrawer2D) draw() {
	for kshaderID, meshIDtoTextureIDtoTintColorIDtoTrans := range d.batch {
		d.kshaders[kshaderID].use()
		d.kshaders[kshaderID].setUniformMat4F("p", d.kdrawerCamera.projectionMatrix())
		d.kshaders[kshaderID].setUniformMat4F("v", d.kdrawerCamera.viewMatrix())
		for kmeshID, textureIDtoTintColorIDtoTrans := range meshIDtoTextureIDtoTintColorIDtoTrans {
			kmesh := d.kmeshes[kmeshID]
			for ktextureID, colorIDtoTrans := range textureIDtoTintColorIDtoTrans {
				var modelMat4Datas = []float32{}
				totalMeshToDraw := int32(0)
				ktexture := d.ktextures[ktextureID]
				ktexture.startDraw()
				for tintColorID, trans := range colorIDtoTrans {
					tintColor := d.tintColors[tintColorID]
					r, g, b, _ := tintColor.RGBA()
					d.kshaders[kshaderID].setUniform3F("tintColor", r, g, b)
					for _, t := range trans {
						totalMeshToDraw++
						for _, tData := range [16]float32(t) {
							modelMat4Datas = append(modelMat4Datas, tData)
						}
					}
					delete(d.batch[kshaderID][kmeshID][ktextureID], tintColorID)
				}
				kmesh.vao.updateModelMat4VBO(modelMat4Datas)
				kmesh.startDraw()
				gl.DrawElementsInstanced(gl.TRIANGLES, kmesh.elementSize, gl.UNSIGNED_INT, gl.PtrOffset(0), totalMeshToDraw)
				ktexture.stopDraw()
				delete(d.batch[kshaderID][kmeshID], ktextureID)
			}
			kmesh.stopDraw()
			delete(d.batch[kshaderID], kmeshID)
		}
		d.kshaders[kshaderID].unuse()
		delete(d.batch, kshaderID)
	}
}

func (kwd2d *kwindowDrawer2D) GetCamera() *kdrawerCamera {
	return &kwd2d.kdrawerCamera
}

func (d *kwindowDrawer2D) DefaultShaderID() kshaderID {
	return d.defaultkshaderID
}

func (d *kwindowDrawer2D) onWindowSizeChange(newWidth, newHeight float32) {
	d.kdrawerCamera.onWindowSizeChange(newWidth, newHeight)
}
