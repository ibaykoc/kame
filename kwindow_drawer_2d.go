package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type KwindowDrawer2DID KwindowDrawerID

func (kwd2d KwindowDrawer2DID) StoreTexturePNG(path string) (ktextureID, error) {
	ktex, err := newktextureFromPNG(path)
	if err != nil {
		return ktextureID(0), err
	}
	(windows[KwindowID(kwd2d)].kwindowDrawer.(*KwindowDrawer2D)).ktextures[ktex.id] = ktex
	return ktex.id, nil
}

func (kwd2d KwindowDrawer2DID) StoreMesh(positions []float32, uvs []float32, elements []uint32) (kmeshID, error) {
	mesh, err := newkmesh(positions, uvs, elements)
	if err != nil {
		return kmeshID(0), err
	}
	(windows[KwindowID(kwd2d)].kwindowDrawer.(*KwindowDrawer2D)).kmeshes[mesh.id] = mesh
	return mesh.id, nil
}

func (kwd2d KwindowDrawer2DID) StoreTintColor(color Kcolor) kcolorID {
	colID := kcolorID(len((windows[KwindowID(kwd2d)].kwindowDrawer.(*KwindowDrawer2D)).tintColors))
	(windows[KwindowID(kwd2d)].kwindowDrawer.(*KwindowDrawer2D)).tintColors[colID] = color
	return colID
}

func (kwd2d KwindowDrawer2DID) DefaultShaderID() kshaderID {
	return windows[KwindowID(kwd2d)].kwindowDrawer.DefaultShaderID()
}

func (kwd2d *KwindowDrawer2D) GetCamera() *kdrawerCamera {
	return &kwd2d.kdrawerCamera
}

type KwindowDrawer2D struct {
	backgroundColor  mgl32.Vec4
	kdrawerCamera    kdrawerCamera
	defaultkshaderID kshaderID
	kshaders         map[kshaderID]*kshader
	kmeshes          map[kmeshID]kmesh
	ktextures        map[ktextureID]ktexture
	tintColors       map[kcolorID]Kcolor
	batch            map[kshaderID]map[kmeshID]map[ktextureID]map[kcolorID][]mgl32.Mat4
}

func newKwindowDrawer2D(config kwindowDrawer2DBuilder) (KwindowDrawer2D, error) {
	if err := gl.Init(); err != nil {
		return KwindowDrawer2D{}, err
	}
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	// version := gl.GoStr(gl.GetString(gl.VERSION))
	// fmt.Println("OpenGL initialized: version", version)
	kshaders := make(map[kshaderID]*kshader)

	defaultShader := createkshader(
		defaultVertexShader,
		defaultSpriteFragmentShader,
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
	gl.ClearColor(config.backgroundColor.Elem())

	return KwindowDrawer2D{
		backgroundColor:  config.backgroundColor,
		defaultkshaderID: kshaderID(defaultShader.id),
		kdrawerCamera:    &camera,
		kshaders:         kshaders,
		kmeshes:          make(map[kmeshID]kmesh),
		ktextures:        make(map[ktextureID]ktexture),
		tintColors:       make(map[kcolorID]Kcolor),
		batch:            make(map[kshaderID]map[kmeshID]map[ktextureID]map[kcolorID][]mgl32.Mat4),
	}, nil
}

func (d *KwindowDrawer2D) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (d *KwindowDrawer2D) AppendDrawable(kdrawable Kdrawable, translation mgl32.Mat4) {
	dw := kdrawable.(Kdrawable2d)
	if _, shaderIDHasAdded := d.batch[dw.Shader]; !shaderIDHasAdded {
		d.batch[dw.Shader] = make(map[kmeshID]map[ktextureID]map[kcolorID][]mgl32.Mat4)
	}
	if _, meshIDHasAdded := d.batch[dw.Shader][dw.Mesh]; !meshIDHasAdded {
		d.batch[dw.Shader][dw.Mesh] = make(map[ktextureID]map[kcolorID][]mgl32.Mat4)
	}
	if _, textureIDHasAdded := d.batch[dw.Shader][dw.Mesh][dw.Texture]; !textureIDHasAdded {
		d.batch[dw.Shader][dw.Mesh][dw.Texture] = make(map[kcolorID][]mgl32.Mat4)
	}
	if _, tintColorIDHasAdded := d.batch[dw.Shader][dw.Mesh][dw.Texture][dw.TintColor]; !tintColorIDHasAdded {
		d.batch[dw.Shader][dw.Mesh][dw.Texture][dw.TintColor] = []mgl32.Mat4{}
	}
	d.batch[dw.Shader][dw.Mesh][dw.Texture][dw.TintColor] = append(d.batch[dw.Shader][dw.Mesh][dw.Texture][dw.TintColor], translation)
}

func (d *KwindowDrawer2D) draw() {
	for kshaderID, meshIDtoTextureIDtoTintColorIDtoTrans := range d.batch {
		d.kshaders[kshaderID].use()
		d.kshaders[kshaderID].setUniformMat4F("p", d.kdrawerCamera.(*kdrawerCamera2D).projectionMatrix())
		d.kshaders[kshaderID].setUniformMat4F("v", d.kdrawerCamera.viewMatrix())
		for kmeshID, textureIDtoTintColorIDtoTrans := range meshIDtoTextureIDtoTintColorIDtoTrans {
			kmesh := d.kmeshes[kmeshID]
			totalMeshToDraw := int32(0)
			var modelMat4Datas = []float32{}
			for ktextureID, colorIDtoTrans := range textureIDtoTintColorIDtoTrans {
				ktexture := d.ktextures[ktextureID]
				ktexture.startDraw()
				for tintColorID, trans := range colorIDtoTrans {
					tintColor := d.tintColors[tintColorID]
					r, g, b, _ := tintColor.RGBA()
					d.kshaders[kshaderID].setUniform3F("tintColor", r, g, b)
					for _, t := range trans {
						totalMeshToDraw++
						// d.kshaders[kshaderID].setUniformMat4F("m", t)
						for _, tData := range [16]float32(t) {
							modelMat4Datas = append(modelMat4Datas, tData)
						}
						// gl.DrawElements(gl.TRIANGLES, kmesh.elementSize, gl.UNSIGNED_INT, gl.PtrOffset(0))
					}
					delete(d.batch[kshaderID][kmeshID][ktextureID], tintColorID)
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

func (d *KwindowDrawer2D) DefaultShaderID() kshaderID {
	return d.defaultkshaderID
}

// func (d *KDrawer) addDrawBacth(shaderID ShaderID, drawableModelID DrawableModelID, translation mgl32.Mat4) {
// 	if _, shaderIDHasAdded := d.batch[shaderID]; !shaderIDHasAdded {
// 		d.batch[shaderID] = make(map[DrawableModelID][]mgl32.Mat4)
// 	}
// 	if _, drawableIDHasAdded := d.batch[shaderID][drawableModelID]; !drawableIDHasAdded {
// 		d.batch[shaderID][drawableModelID] = make([]mgl32.Mat4, 1)
// 	}
// 	d.batch[shaderID][drawableModelID] = append(d.batch[shaderID][drawableModelID], translation)
// }
