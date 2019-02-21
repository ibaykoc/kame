package kame

// import (
// 	"fmt"
// 	"math"

// 	"github.com/go-gl/mathgl/mgl32"

// 	"github.com/go-gl/gl/v4.1-core/gl"
// )

// type MeshType int

// const (
// 	Triangle MeshType = iota
// 	Quad
// )

// type drawableDataComponent int

// const (
// 	drawableDataComponentPosition drawableDataComponent = iota
// 	drawableDataComponentUV
// 	drawableDataComponentNormal
// 	drawableDataComponentElements
// )

// type drawableModel struct {
// 	vao         VAO
// 	textureID   uint32
// 	tintColor   mgl32.Vec3
// 	elementSize int32
// 	shaderID    ShaderID
// }

// func newDrawableModel(vbo VBO, ebo []uint32, shaderID ShaderID) drawableModel {
// 	fmt.Printf("Create new drawable model\n")
// 	vao := createVAO()
// 	vao.storeVBO(vbo)
// 	vao.storeEBO(ebo)
// 	return drawableModel{
// 		vao:         vao,
// 		textureID:   math.MaxUint32,
// 		elementSize: int32(len(ebo)),
// 		shaderID:    shaderID,
// 		tintColor:   mgl32.Vec3{1, 1, 1},
// 	}
// }

// func (dm *drawableModel) startDraw() {
// 	gl.ActiveTexture(gl.TEXTURE0)
// 	gl.BindTexture(gl.TEXTURE_2D, dm.textureID)
// 	dm.vao.use()
// }

// // LoadTextureFile load texture file into drawable model
// func (dm *drawableModel) loadTextureFile(path string) error {
// 	loadedTexture, textureLoaded := window.kdrawer.loadedTextureFile[path]
// 	if textureLoaded {
// 		dm.textureID = loadedTexture
// 		return nil
// 	}
// 	textureID, err := loadTextureFile(path)
// 	if err != nil {
// 		return err
// 	}
// 	window.kdrawer.loadedTextureFile[path] = textureID
// 	dm.textureID = textureID
// 	return nil
// }

// func (dm *drawableModel) stopDraw() {
// 	dm.vao.unuse()
// 	gl.ActiveTexture(gl.TEXTURE0)
// 	gl.BindTexture(gl.TEXTURE_2D, 0)
// }

// // CreateDrawableModel create drawable model with built in mesh
// func newBuiltInDrawableModel(meshType MeshType) drawableModel {
// 	var positions []float32
// 	var uvs []float32
// 	var elements []uint32
// 	switch meshType {
// 	case Triangle:
// 		positions = []float32{
// 			+0.0, +1.0, +0.0, // Center Top
// 			-1.0, -1.0, +0.0, // Left Bottom
// 			+1.0, -1.0, +0.0, // Right Bottom
// 		}

// 		uvs = []float32{
// 			// UVs
// 			+0.5, +1.0,
// 			+0.0, +0.0,
// 			+1.0, +0.0,
// 		}

// 		elements = []uint32{
// 			0, 1, 2, // First Triangle CCW (front facing you)

// 		}
// 	case Quad:
// 		positions = []float32{
// 			//Positions
// 			-0.5, +0.5, +0.0, // Left Top
// 			+0.5, +0.5, +0.0, // Right Top
// 			-0.5, -0.5, +0.0, // Left Bottom
// 			+0.5, -0.5, +0.0, // Right Bottom
// 		}

// 		uvs = []float32{
// 			// UVs
// 			+0.0, +1.0,
// 			+1.0, +1.0,
// 			+0.0, +0.0,
// 			+1.0, +0.0,
// 		}

// 		elements = []uint32{
// 			0, 2, 3, // first triangle CCW (front facing you)
// 			0, 3, 1,
// 		}
// 	}

// 	model, err := newDrawableModel1(positions, uvs, elements)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return model
// }

// // CreateDrawableModelT create drawable model with built in mesh and defined texture
// func GetBuiltInDrawableModel(meshType MeshType) DrawableModelID {
// 	var modelID DrawableModelID
// 	switch meshType {
// 	case Quad:
// 		modelID = quadDrawableModelID
// 	}
// 	return modelID
// }

// // CreateDrawableModel0 creates drawable models with position data
// func newDrawableModel0(positions []float32, elements []uint32) (drawableModel, error) {
// 	// positions data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentPosition, len(positions)); foundErr {
// 		return drawableModel{}, err
// 	}

// 	// elements data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentElements, len(elements)); foundErr {
// 		return drawableModel{}, err
// 	}

// 	perVertexDataCount := int32(3)
// 	data := positions
// 	vbo := createFloat32VBO(perVertexDataCount,
// 		[]VBOData{
// 			// Pos
// 			VBOData{
// 				count:      3,
// 				byteOffset: 0,
// 			},
// 		},
// 		data)
// 	return newDrawableModel(vbo, elements, window.kdrawer.defaultShaderProgramID), nil

// }

// // CreateDrawableModel1 creates drawable models with position, and UV data
// func newDrawableModel1(positions []float32, uvs []float32, elements []uint32) (drawableModel, error) {
// 	{ //Validation
// 		// positions data validation
// 		if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentPosition, len(positions)); foundErr {
// 			return drawableModel{}, err
// 		}

// 		// uvs data validation
// 		if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentUV, len(uvs)); foundErr {
// 			return drawableModel{}, err
// 		}
// 		if foundErr, err := validateVertexDataMatchToPosition(drawableDataComponentUV, uvs, positions); foundErr {
// 			return drawableModel{}, err
// 		}

// 		// indices data validation
// 		if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentElements, len(elements)); foundErr {
// 			return drawableModel{}, err
// 		}
// 	}

// 	perVertexDataCount := int32(5)
// 	data := make([]float32, len(positions)+len(uvs))
// 	for i := 0; i < len(data)/int(perVertexDataCount); i++ {
// 		dataI := i * int(perVertexDataCount)
// 		// position
// 		positionI := i * 3
// 		data[dataI] = positions[positionI]
// 		data[dataI+1] = positions[positionI+1]
// 		data[dataI+2] = positions[positionI+2]
// 		// uv
// 		uvI := i * 2
// 		data[dataI+3] = uvs[uvI]
// 		data[dataI+4] = uvs[uvI+1]
// 	}
// 	vbo := createFloat32VBO(perVertexDataCount,
// 		[]VBOData{
// 			// Pos
// 			VBOData{
// 				count:      3,
// 				byteOffset: 0,
// 			},
// 			// UV
// 			VBOData{
// 				count:      2,
// 				byteOffset: 3 * 4,
// 			},
// 		},
// 		data)
// 	return newDrawableModel(vbo, elements, window.kdrawer.defaultShaderProgramID), nil
// }

// // CreateDrawableModel2 creates drawable models with position, and normal data
// func newDrawableModel2(positions []float32, normals []float32, elements []uint32) (drawableModel, error) {
// 	// positions data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentPosition, len(positions)); foundErr {
// 		return drawableModel{}, err
// 	}

// 	// normals data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentNormal, len(normals)); foundErr {
// 		return drawableModel{}, err
// 	}
// 	if foundErr, err := validateVertexDataMatchToPosition(drawableDataComponentNormal, normals, positions); foundErr {
// 		return drawableModel{}, err
// 	}

// 	// indices data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentElements, len(elements)); foundErr {
// 		return drawableModel{}, err
// 	}
// 	perVertexDataCount := int32(6)
// 	data := make([]float32, len(positions)+len(normals))
// 	for i := 0; i < len(data)/int(perVertexDataCount); i++ {
// 		dataI := i * int(perVertexDataCount)
// 		// position
// 		positionI := i * 3
// 		data[dataI] = positions[positionI]
// 		data[dataI+1] = positions[positionI+1]
// 		data[dataI+2] = positions[positionI+2]
// 		// normal
// 		normalI := i * 3
// 		data[dataI+3] = normals[normalI]
// 		data[dataI+4] = normals[normalI+1]
// 		data[dataI+5] = normals[normalI+2]
// 	}
// 	vbo := createFloat32VBO(perVertexDataCount,
// 		[]VBOData{
// 			// Pos
// 			VBOData{
// 				count:      3,
// 				byteOffset: 0,
// 			},
// 			// Norms
// 			VBOData{
// 				count:      3,
// 				byteOffset: 5 * 4,
// 			},
// 		},
// 		data)

// 	return newDrawableModel(vbo, elements, window.kdrawer.defaultShaderProgramID), nil
// }

// // newDrawableModel3 creates drawable models with position, UV, and normal data
// func newDrawableModel3(positions []float32, uvs []float32, normals []float32, elements []uint32) (drawableModel, error) {
// 	// positions data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentPosition, len(positions)); foundErr {
// 		return drawableModel{}, err
// 	}

// 	// uvs data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentUV, len(uvs)); foundErr {
// 		return drawableModel{}, err
// 	}
// 	if foundErr, err := validateVertexDataMatchToPosition(drawableDataComponentUV, uvs, positions); foundErr {
// 		return drawableModel{}, err
// 	}

// 	// normals data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentNormal, len(normals)); foundErr {
// 		return drawableModel{}, err
// 	}
// 	if foundErr, err := validateVertexDataMatchToPosition(drawableDataComponentNormal, normals, positions); foundErr {
// 		return drawableModel{}, err
// 	}

// 	// indices data validation
// 	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentElements, len(elements)); foundErr {
// 		return drawableModel{}, err
// 	}

// 	perVertexDataCount := int32(8)
// 	data := make([]float32, len(positions)+len(uvs)+len(normals))
// 	for i := 0; i < len(data)/int(perVertexDataCount); i++ {
// 		dataI := i * int(perVertexDataCount)
// 		// position
// 		positionI := i * 3
// 		data[dataI] = positions[positionI]
// 		data[dataI+1] = positions[positionI+1]
// 		data[dataI+2] = positions[positionI+2]
// 		// uv
// 		uvI := i * 2
// 		data[dataI+3] = uvs[uvI]
// 		data[dataI+4] = uvs[uvI+1]
// 		// normal
// 		normalI := i * 3
// 		data[dataI+5] = normals[normalI]
// 		data[dataI+6] = normals[normalI+1]
// 		data[dataI+7] = normals[normalI+2]
// 	}

// 	vbo := createFloat32VBO(perVertexDataCount,
// 		[]VBOData{
// 			// Pos
// 			VBOData{
// 				count:      3,
// 				byteOffset: 0,
// 			},
// 			// UV
// 			VBOData{
// 				count:      2,
// 				byteOffset: 3 * 4,
// 			},
// 			// Norms
// 			VBOData{
// 				count:      3,
// 				byteOffset: 5 * 4,
// 			},
// 		},
// 		data)

// 	return newDrawableModel(vbo, elements, window.kdrawer.defaultShaderProgramID), nil
// }

// // Dispose dispose drawable model
// func (dm *drawableModel) Dispose() {
// 	dm.vao.dispose()
// }

// func validateDrawableDataComponentLength(drawableDataComponent drawableDataComponent, vertexDataToValidateLength int) (foundError bool, err error) {
// 	var shouldBeDivisibleBy int
// 	var dataName string

// 	switch drawableDataComponent {
// 	case drawableDataComponentPosition:
// 		shouldBeDivisibleBy = 3
// 		dataName = "POSITION"
// 	case drawableDataComponentUV:
// 		shouldBeDivisibleBy = 2
// 		dataName = "UV"
// 	case drawableDataComponentNormal:
// 		shouldBeDivisibleBy = 3
// 		dataName = "NORMAL"
// 	case drawableDataComponentElements:
// 		shouldBeDivisibleBy = 3
// 		dataName = "INDICES"
// 	}
// 	if vertexDataToValidateLength%shouldBeDivisibleBy != 0 {
// 		return true, fmt.Errorf("\n***\t%s length should be divisible by %d\n***\t%s length found: %d", dataName, shouldBeDivisibleBy, dataName, vertexDataToValidateLength)
// 	}
// 	return false, nil
// }

// func validateVertexDataMatchToPosition(drawableDataComponent drawableDataComponent, vertexDataToValidate []float32, positionData []float32) (foundError bool, err error) {
// 	var dataName string
// 	var toValidateVertexLen int
// 	positionVertexLen := len(positionData) / 3
// 	switch drawableDataComponent {
// 	case drawableDataComponentUV:
// 		toValidateVertexLen = len(vertexDataToValidate) / 2
// 		dataName = "UV"
// 	case drawableDataComponentNormal:
// 		toValidateVertexLen = len(vertexDataToValidate) / 3
// 		dataName = "NORMAL"
// 	}
// 	if toValidateVertexLen != positionVertexLen {
// 		return true, fmt.Errorf("\n***\t%s VERTEX length should be match with POSITION VERTEX length\n***\t%s VERTEX length : %d\n***\tPOSITION VERTEX length : %d", dataName, dataName, toValidateVertexLen, positionVertexLen)
// 	}
// 	return false, nil
// }

// func (dmID DrawableModelID) LoadTexture(filePath string) {
// 	window.kdrawer.models[dmID].loadTextureFile(filePath)
// }
