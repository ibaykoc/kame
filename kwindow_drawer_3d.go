package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type KwindowDrawer3DController struct {
	KwindowDrawerController
	kwindowDrawer3D *kwindowDrawer3D
}

func (wdCon KwindowDrawer3DController) StoreMesh(positions []float32, uvs []float32, normals []float32, elements []uint32) (kmeshID, error) {
	mesh, err := newkmeshPosUVNormals(positions, uvs, normals, elements)
	if err != nil {
		return kmeshID(0), err
	}
	wdCon.kwindowDrawer.kmeshes[mesh.id] = mesh
	return mesh.id, nil
}

type kwindowDrawer3D struct {
	kwindowDrawer
	batch map[kshaderID]map[kmeshID]map[ktextureID][]mgl32.Mat4
}

func newKwindowDrawer3D(config kwindowDrawer3DBuilder) (kwindowDrawer3D, error) {
	if err := gl.Init(); err != nil {
		return kwindowDrawer3D{}, err
	}
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	debug.pf("3D OpenGL initialized: version %s\n", gl.GoStr(gl.GetString(gl.VERSION)))
	kshaders := make(map[kshaderID]*kshader)

	defaultShader := createkshader(
		default3DVertexShader,
		default3DFragmentShader,
		[]string{
			"v", "p",
			"texture0",
		},
	)

	camera := newkdrawerCamera3D(config.windowWidth, config.windowHeight, config.fov)

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
		ktextures:        make(map[ktextureID]ktexture),
	}

	return kwindowDrawer3D{
		kwindowDrawer: kwd,
		batch:         make(map[kshaderID]map[kmeshID]map[ktextureID][]mgl32.Mat4),
	}, nil
}

func (d *kwindowDrawer3D) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (d *kwindowDrawer3D) AppendDrawable(kdrawable Kdrawable, translation mgl32.Mat4) {
	dw, ok := kdrawable.(Kdrawable3d)
	if !ok {
		panic("Drawer3D needs drawable3D")
	}
	if _, shaderIDHasAdded := d.batch[dw.ShaderID]; !shaderIDHasAdded {
		d.batch[dw.ShaderID] = make(map[kmeshID]map[ktextureID][]mgl32.Mat4)
	}
	if _, meshIDHasAdded := d.batch[dw.ShaderID][dw.MeshID]; !meshIDHasAdded {
		d.batch[dw.ShaderID][dw.MeshID] = make(map[ktextureID][]mgl32.Mat4)
	}
	if _, textureIDHasAdded := d.batch[dw.ShaderID][dw.MeshID][dw.TextureID]; !textureIDHasAdded {
		d.batch[dw.ShaderID][dw.MeshID][dw.TextureID] = []mgl32.Mat4{}
	}
	d.batch[dw.ShaderID][dw.MeshID][dw.TextureID] = append(d.batch[dw.ShaderID][dw.MeshID][dw.TextureID], translation)
}

func (d *kwindowDrawer3D) draw() {
	for kshaderID, meshIDtoTextureIDtoTintColorIDtoTrans := range d.batch {
		d.kshaders[kshaderID].use()
		d.kshaders[kshaderID].setUniformMat4F("p", d.kdrawerCamera.projectionMatrix())
		d.kshaders[kshaderID].setUniformMat4F("v", d.kdrawerCamera.viewMatrix())
		for kmeshID, textureIDtoTintColorIDtoTrans := range meshIDtoTextureIDtoTintColorIDtoTrans {
			kmesh := d.kmeshes[kmeshID]
			totalMeshToDraw := int32(0)
			var modelMat4Datas = []float32{}
			for ktextureID, trans := range textureIDtoTintColorIDtoTrans {
				ktexture := d.ktextures[ktextureID]
				ktexture.startDraw()
				for _, t := range trans {
					totalMeshToDraw++
					// d.kshaders[kshaderID].setUniformMat4F("m", t)
					for _, tData := range [16]float32(t) {
						modelMat4Datas = append(modelMat4Datas, tData)
					}
					// gl.DrawElements(gl.TRIANGLES, kmesh.elementSize, gl.UNSIGNED_INT, gl.PtrOffset(0))
				}
				// ktexture.stopDraw()
				delete(d.batch[kshaderID][kmeshID], ktextureID)
			}
			// kmesh.stopDraw()
			kmesh.vao.updateModelMat4VBO(modelMat4Datas)
			kmesh.startDraw()
			gl.DrawElementsInstanced(gl.TRIANGLES, kmesh.elementSize, gl.UNSIGNED_INT, gl.PtrOffset(0), totalMeshToDraw)
			delete(d.batch[kshaderID], kmeshID)
		}
		// d.kshaders[kshaderID].unuse()
		delete(d.batch, kshaderID)
	}
}

func (kwd3d *kwindowDrawer3D) GetCamera() *kdrawerCamera {
	return &kwd3d.kdrawerCamera
}

func (d *kwindowDrawer3D) DefaultShaderID() kshaderID {
	return d.defaultkshaderID
}

func (d *kwindowDrawer3D) onWindowSizeChange(newWidth, newHeight float32) {
	d.kdrawerCamera.onWindowSizeChange(newWidth, newHeight)
}
