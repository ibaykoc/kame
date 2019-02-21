package kame

// import (
// 	"github.com/go-gl/gl/v4.1-core/gl"

// 	"github.com/go-gl/mathgl/mgl32"
// )

// type DrawableModelID int
// type ShaderID uint32

// // KDrawer to draw something onto the screen
// type KDrawer struct {
// 	BackgroundColor              Color
// 	camera                       Camera
// 	defaultShaderProgramID       ShaderID
// 	defaultSpriteShaderProgramID ShaderID
// 	loadedTextureFile            map[string]uint32
// 	models                       map[DrawableModelID]*drawableModel
// 	shaders                      map[ShaderID]*ShaderProgram
// 	batch                        map[ShaderID]map[DrawableModelID][]mgl32.Mat4 //Shader DrawableModelID to Translate
// }

// var quadDrawableModelID DrawableModelID

// func newDrawer2D(backgroundColor Color) (*KDrawer, error) {
// 	return newDrawer(Orthographic, backgroundColor)
// }

// func newDrawer3D(backgroundColor Color) (*KDrawer, error) {
// 	return newDrawer(Perspective, backgroundColor)
// }

// func newDrawer(cameraType ProjectionType, backgroundColor Color) (*KDrawer, error) {
// 	bgColor := backgroundColor
// 	if err := gl.Init(); err != nil {
// 		return nil, err
// 	}
// 	// Enable alpha blending
// 	gl.Enable(gl.BLEND)
// 	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
// 	gl.Enable(gl.DEPTH_TEST)
// 	gl.Enable(gl.CULL_FACE)
// 	// version := gl.GoStr(gl.GetString(gl.VERSION))
// 	// fmt.Println("OpenGL initialized: version", version)

// 	shaders := make(map[ShaderID]*ShaderProgram)

// 	defaultShaderProgram := createShaderProgram(
// 		defaultVertexShader,
// 		defaultFragmentShader,
// 		[]string{
// 			"m",
// 			"v",
// 			"p",
// 			"texture0",
// 			"tintColor",
// 		},
// 	)

// 	shaders[ShaderID(defaultShaderProgram.id)] = &defaultShaderProgram

// 	defaultSpriteShaderProgram := createShaderProgram(
// 		defaultVertexShader,
// 		defaultSpriteFragmentShader,
// 		[]string{
// 			"m",
// 			"v",
// 			"p",
// 			"texture0",
// 			"tintColor",
// 		},
// 	)

// 	shaders[ShaderID(defaultSpriteShaderProgram.id)] = &defaultSpriteShaderProgram

// 	var camera Camera
// 	if cameraType == Orthographic {
// 		camera = createCamera2D(50)
// 	} else {
// 		camera = createCamera3D(mgl32.DegToRad(90))
// 	}

// 	defaultShaderProgram.start()
// 	defaultShaderProgram.setUniformMat4F("v", camera.viewMatrix())
// 	defaultShaderProgram.setUniformMat4F("p", camera.projectionMatrix())
// 	defaultShaderProgram.setUniform1i("texture0", 0)
// 	defaultShaderProgram.stop()

// 	defaultSpriteShaderProgram.start()
// 	defaultSpriteShaderProgram.setUniformMat4F("v", camera.viewMatrix())
// 	defaultSpriteShaderProgram.setUniformMat4F("p", camera.projectionMatrix())
// 	defaultSpriteShaderProgram.setUniform1i("texture0", 0)
// 	defaultSpriteShaderProgram.stop()

// 	gl.ClearColor(
// 		bgColor.R,
// 		bgColor.G,
// 		bgColor.B,
// 		bgColor.A)

// 	d := &KDrawer{
// 		BackgroundColor:              bgColor,
// 		defaultShaderProgramID:       ShaderID(defaultShaderProgram.id),
// 		defaultSpriteShaderProgramID: ShaderID(defaultSpriteShaderProgram.id),
// 		camera:            camera,
// 		shaders:           shaders,
// 		models:            make(map[DrawableModelID]*drawableModel),
// 		loadedTextureFile: make(map[string]uint32),
// 		batch:             make(map[ShaderID]map[DrawableModelID][]mgl32.Mat4),
// 	}
// 	return d, nil
// }

// func (d *KDrawer) start() {
// 	quadModel := newBuiltInDrawableModel(Quad)
// 	quadDrawableModelID = d.storeModel(quadModel)
// }

// func (d *KDrawer) MoveCameraRelative(x, y, z float32) {
// 	d.camera.Move(x, y, z)
// 	for _, shader := range d.shaders {
// 		shader.start()
// 		shader.setUniformMat4F("v", d.camera.viewMatrix())
// 		shader.stop()
// 	}
// }

// func (d *KDrawer) clear() {
// 	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
// }

// // Draw draw model at default position
// func (d *KDrawer) Draw(id DrawableModelID) {
// 	dm := id.getDrawableModel()
// 	d.addDrawBacth(dm.shaderID, id, mgl32.Translate3D(0, 0, 0))
// }

// // Draw0 draw model at specified position
// func (d *KDrawer) DrawAtPosition(id DrawableModelID, position mgl32.Vec3) {
// 	dm := id.getDrawableModel()
// 	d.addDrawBacth(dm.shaderID, id, mgl32.Translate3D(position.Elem()))
// }

// func (d *KDrawer) DrawAtRotation(id DrawableModelID, rotation mgl32.Vec3) {
// 	rValue := rotation.Len()
// 	rAxis := rotation.Normalize()
// 	dm := id.getDrawableModel()
// 	d.addDrawBacth(dm.shaderID, id, mgl32.Translate3D(0, 0, 0).Mul4(mgl32.HomogRotate3D(rValue, rAxis)))
// }
// func (d *KDrawer) DrawAt(id DrawableModelID, translation mgl32.Mat4) {
// 	dm := id.getDrawableModel()
// 	d.addDrawBacth(dm.shaderID, id, translation)
// }

// func (d *KDrawer) addDrawBacth(shaderID ShaderID, drawableModelID DrawableModelID, translation mgl32.Mat4) {
// 	if _, shaderIDHasAdded := d.batch[shaderID]; !shaderIDHasAdded {
// 		d.batch[shaderID] = make(map[DrawableModelID][]mgl32.Mat4)
// 	}
// 	if _, drawableIDHasAdded := d.batch[shaderID][drawableModelID]; !drawableIDHasAdded {
// 		d.batch[shaderID][drawableModelID] = make([]mgl32.Mat4, 1)
// 	}
// 	d.batch[shaderID][drawableModelID] = append(d.batch[shaderID][drawableModelID], translation)
// }

// func (d *KDrawer) drawBatch() {
// 	for shaderID, dmIDToTrans := range d.batch {
// 		d.shaders[shaderID].start()
// 		d.shaders[shaderID].setUniformMat4F("v", d.camera.viewMatrix())

// 		for dmID, trans := range dmIDToTrans {
// 			dm := dmID.getDrawableModel()
// 			dm.startDraw()
// 			for _, t := range trans {
// 				d.shaders[shaderID].setUniformMat4F("m", t)
// 				r, g, b := dm.tintColor.Elem()
// 				d.shaders[shaderID].setUniform3F("tintColor", r, g, b)
// 				gl.DrawElements(gl.TRIANGLES, dm.elementSize, gl.UNSIGNED_INT, gl.PtrOffset(0))
// 			}
// 			dm.stopDraw()
// 			delete(d.batch[shaderID], dmID)
// 		}
// 		d.shaders[shaderID].stop()
// 		delete(d.batch, shaderID)
// 	}
// }

// func (d *KDrawer) changeSize(width int32, height int32) {
// 	for _, shader := range d.shaders {
// 		shader.start()
// 		shader.setUniformMat4F("p", d.camera.projectionMatrix())
// 		shader.stop()
// 	}
// 	gl.Viewport(0, 0, width, height)
// }

// func (d *KDrawer) dispose() {
// 	for _, shader := range d.shaders {
// 		shader.dispose()
// 	}
// 	for _, textureID := range d.loadedTextureFile {
// 		gl.DeleteTextures(1, &textureID)
// 	}
// }

// func (d *KDrawer) storeModel(dm drawableModel) DrawableModelID {
// 	dmID := DrawableModelID(len(d.models))
// 	d.models[dmID] = &dm
// 	return dmID
// }

// func (dm DrawableModelID) getDrawableModel() *drawableModel {
// 	return window.kdrawer.models[dm]
// }

// func (dm *DrawableModelID) SetTintColor(color mgl32.Vec3) {
// 	dm.getDrawableModel().tintColor = color
// }
