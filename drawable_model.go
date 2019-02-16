package kame

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type MeshType int

const (
	Triangle MeshType = iota
	Quad
)

type drawableDataComponent int

const (
	drawableDataComponentPosition drawableDataComponent = iota
	drawableDataComponentUV
	drawableDataComponentNormal
	drawableDataComponentElements
)

type DrawableModel struct {
	vao         VAO
	textureIDs  []uint32
	elementSize int32
}

func (dm *DrawableModel) startDraw() {
	if len(dm.textureIDs) <= 0 {
		window.drawer.defaultShaderProgram.SetUniform1F("hasTexture", 0.0)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, window.drawer.defaultShaderProgram.defaultTextureID)
	} else {
		window.drawer.defaultShaderProgram.SetUniform1F("hasTexture", 1.0)
		for i := uint32(0); i < uint32(len(dm.textureIDs)); i++ {
			gl.ActiveTexture(gl.TEXTURE1 + i)
			gl.BindTexture(gl.TEXTURE_2D, dm.textureIDs[i])
		}
	}
	dm.vao.use()
}

// LoadTextureFile load texture file into drawable model
func (dm *DrawableModel) loadTextureFile(path string) error {
	loadedTexture, textureLoaded := window.drawer.loadedTextureFile[path]
	if textureLoaded {
		dm.textureIDs = append(dm.textureIDs, loadedTexture)
		return nil
	}
	textureID, err := loadTextureFile(path)
	if err != nil {
		return err
	}
	window.drawer.loadedTextureFile[path] = textureID
	dm.textureIDs = append(dm.textureIDs, textureID)
	return nil
}

func (dm *DrawableModel) stopDraw() {
	dm.vao.unuse()
	if len(dm.textureIDs) > 0 {
		window.drawer.defaultShaderProgram.SetUniform1F("hasTexture", 0.0)
		for i := uint32(1); i < uint32(len(dm.textureIDs)); i++ {
			gl.ActiveTexture(gl.TEXTURE1 + i)
			gl.BindTexture(gl.TEXTURE_2D, 0)
		}
	}
}

func CreateDrawableModel(meshType MeshType) DrawableModel {
	var positions []float32
	var uvs []float32
	var elements []uint32
	switch meshType {
	case Triangle:
		positions = []float32{
			+0.0, +1.0, +0.0, // Center Top
			-1.0, -1.0, +0.0, // Left Bottom
			+1.0, -1.0, +0.0, // Right Bottom
		}

		uvs = []float32{
			// UVs
			+0.5, +1.0,
			+0.0, +0.0,
			+1.0, +0.0,
		}

		elements = []uint32{
			0, 1, 2, // First Triangle CCW (front facing you)

		}
	case Quad:
		positions = []float32{
			//Positions
			-1.0, +1.0, +0.0, // Left Top
			+1.0, +1.0, +0.0, // Right Top
			-1.0, -1.0, +0.0, // Left Bottom
			+1.0, -1.0, +0.0, // Right Bottom
		}

		uvs = []float32{
			// UVs
			+0.0, +1.0,
			+1.0, +1.0,
			+0.0, +0.0,
			+1.0, +0.0,
		}

		elements = []uint32{
			0, 2, 3, // first triangle CCW (front facing you)
			0, 3, 1,
		}
	}
	model, err := CreateDrawableModel1(positions, uvs, elements)
	if err != nil {
		panic(err)
	}
	return model
}
func CreateDrawableModelT(meshType MeshType, texturePath string) (DrawableModel, error) {
	dm := CreateDrawableModel(meshType)
	if err := dm.loadTextureFile(texturePath); err != nil {
		return DrawableModel{}, err
	}
	return dm, nil
}

// CreateDrawableModel0 creates drawable models with position data
func CreateDrawableModel0(positions []float32, elements []uint32) (DrawableModel, error) {
	// positions data validation
	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentPosition, len(positions)); foundErr {
		return DrawableModel{}, err
	}

	// elements data validation
	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentElements, len(elements)); foundErr {
		return DrawableModel{}, err
	}

	perVertexDataCount := int32(3)
	data := positions
	vao := createVAO()
	vbo := createFloat32VBO(perVertexDataCount,
		[]VBOData{
			// Pos
			VBOData{
				count:      3,
				byteOffset: 0,
			},
		},
		data)
	vao.storeVBO(vbo)
	vao.storeEBO(elements)
	return DrawableModel{
		vao:         vao,
		textureIDs:  []uint32{},
		elementSize: int32(len(elements)),
	}, nil
}

// CreateDrawableModel1 creates drawable models with position, UV, and texture data
func CreateDrawableModel1(positions []float32, uvs []float32, elements []uint32) (DrawableModel, error) {
	{ //Validation
		// positions data validation
		if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentPosition, len(positions)); foundErr {
			return DrawableModel{}, err
		}

		// uvs data validation
		if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentUV, len(uvs)); foundErr {
			return DrawableModel{}, err
		}
		if foundErr, err := validateVertexDataMatchToPosition(drawableDataComponentUV, uvs, positions); foundErr {
			return DrawableModel{}, err
		}

		// indices data validation
		if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentElements, len(elements)); foundErr {
			return DrawableModel{}, err
		}
		//#endregion
	}

	perVertexDataCount := int32(5)
	data := make([]float32, len(positions)+len(uvs))
	for i := 0; i < len(data)/int(perVertexDataCount); i++ {
		dataI := i * int(perVertexDataCount)
		// position
		positionI := i * 3
		data[dataI] = positions[positionI]
		data[dataI+1] = positions[positionI+1]
		data[dataI+2] = positions[positionI+2]
		// uv
		uvI := i * 2
		data[dataI+3] = uvs[uvI]
		data[dataI+4] = uvs[uvI+1]
	}
	vao := createVAO()
	vbo := createFloat32VBO(perVertexDataCount,
		[]VBOData{
			// Pos
			VBOData{
				count:      3,
				byteOffset: 0,
			},
			// UV
			VBOData{
				count:      2,
				byteOffset: 3 * 4,
			},
		},
		data)
	vao.storeVBO(vbo)
	vao.storeEBO(elements)
	return DrawableModel{
		vao:         vao,
		textureIDs:  []uint32{},
		elementSize: int32(len(elements)),
	}, nil
}

func CreateDrawableModel1T(positions []float32, uvs []float32, elements []uint32, texturePath string) (DrawableModel, error) {
	dm, err := CreateDrawableModel1(positions, uvs, elements)
	if err != nil {
		return DrawableModel{}, nil
	}
	textureID, err := loadTextureFile(texturePath)
	if err != nil {
		return DrawableModel{}, nil
	}
	dm.textureIDs = []uint32{textureID}
	return dm, nil
}

// CreateDrawableModel2 creates drawable models with position, UV, normal, and texture data
func CreateDrawableModel2(positions []float32, uvs []float32, normals []float32, elements []uint32, texturePath string) (DrawableModel, error) {
	// positions data validation
	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentPosition, len(positions)); foundErr {
		return DrawableModel{}, err
	}

	// uvs data validation
	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentUV, len(uvs)); foundErr {
		return DrawableModel{}, err
	}
	if foundErr, err := validateVertexDataMatchToPosition(drawableDataComponentUV, uvs, positions); foundErr {
		return DrawableModel{}, err
	}

	// normals data validation
	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentNormal, len(normals)); foundErr {
		return DrawableModel{}, err
	}
	if foundErr, err := validateVertexDataMatchToPosition(drawableDataComponentNormal, normals, positions); foundErr {
		return DrawableModel{}, err
	}

	// indices data validation
	if foundErr, err := validateDrawableDataComponentLength(drawableDataComponentElements, len(elements)); foundErr {
		return DrawableModel{}, err
	}
	perVertexDataCount := int32(8)
	data := make([]float32, len(positions)+len(uvs)+len(normals))
	for i := 0; i < len(data)/int(perVertexDataCount); i++ {
		dataI := i * int(perVertexDataCount)
		// position
		positionI := i * 3
		data[dataI] = positions[positionI]
		data[dataI+1] = positions[positionI+1]
		data[dataI+2] = positions[positionI+2]
		// uv
		uvI := i * 2
		data[dataI+3] = uvs[uvI]
		data[dataI+4] = uvs[uvI+1]
		// normal
		normalI := i * 3
		data[dataI+5] = normals[normalI]
		data[dataI+6] = normals[normalI+1]
		data[dataI+7] = normals[normalI+2]
	}

	vao := createVAO()

	vbo := createFloat32VBO(perVertexDataCount,
		[]VBOData{
			// Pos
			VBOData{
				count:      3,
				byteOffset: 0,
			},
			// UV
			VBOData{
				count:      2,
				byteOffset: 3 * 4,
			},
			// Norms
			VBOData{
				count:      3,
				byteOffset: 5 * 4,
			},
		},
		data)
	vao.storeVBO(vbo)
	vao.storeEBO(elements)
	if loadedTexture, foundLoadedTexture := window.drawer.loadedTextureFile[texturePath]; foundLoadedTexture {
		return DrawableModel{
			vao:         vao,
			textureIDs:  []uint32{loadedTexture},
			elementSize: int32(len(elements)),
		}, nil
	}

	textureID, err := loadTextureFile(texturePath)
	if err != nil {
		return DrawableModel{}, nil
	}
	window.drawer.loadedTextureFile[texturePath] = textureID
	return DrawableModel{
		vao:         vao,
		textureIDs:  []uint32{textureID},
		elementSize: int32(len(elements)),
	}, nil
}

func (dm *DrawableModel) Dispose() {
	dm.vao.dispose()
}

func validateDrawableDataComponentLength(drawableDataComponent drawableDataComponent, vertexDataToValidateLength int) (foundError bool, err error) {
	var shouldBeDivisibleBy int
	var dataName string

	switch drawableDataComponent {
	case drawableDataComponentPosition:
		shouldBeDivisibleBy = 3
		dataName = "POSITION"
	case drawableDataComponentUV:
		shouldBeDivisibleBy = 2
		dataName = "UV"
	case drawableDataComponentNormal:
		shouldBeDivisibleBy = 3
		dataName = "NORMAL"
	case drawableDataComponentElements:
		shouldBeDivisibleBy = 3
		dataName = "INDICES"
	}
	if vertexDataToValidateLength%shouldBeDivisibleBy != 0 {
		return true, fmt.Errorf("\n***\t%s length should be divisible by %d\n***\t%s length found: %d", dataName, shouldBeDivisibleBy, dataName, vertexDataToValidateLength)
	}
	return false, nil
}

func validateVertexDataMatchToPosition(drawableDataComponent drawableDataComponent, vertexDataToValidate []float32, positionData []float32) (foundError bool, err error) {
	var dataName string
	var toValidateVertexLen int
	positionVertexLen := len(positionData) / 3
	switch drawableDataComponent {
	case drawableDataComponentUV:
		toValidateVertexLen = len(vertexDataToValidate) / 2
		dataName = "UV"
	case drawableDataComponentNormal:
		toValidateVertexLen = len(vertexDataToValidate) / 3
		dataName = "NORMAL"
	}
	if toValidateVertexLen != positionVertexLen {
		return true, fmt.Errorf("\n***\t%s VERTEX length should be match with POSITION VERTEX length\n***\t%s VERTEX length : %d\n***\tPOSITION VERTEX length : %d", dataName, dataName, toValidateVertexLen, positionVertexLen)
	}
	return false, nil
}
