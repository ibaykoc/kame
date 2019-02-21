package kame

import "fmt"

type meshDataComponent int

const (
	meshDataComponentPosition meshDataComponent = iota
	meshDataComponentUV
	meshDataComponentNormal
	meshDataComponentElements
)

type kmeshID uint32

type kmesh struct {
	id          kmeshID
	vao         VAO
	elementSize int32
}

func (kmesh *kmesh) startDraw() {
	kmesh.vao.use()
}
func (kmesh *kmesh) stopDraw() {
	kmesh.vao.unuse()
}

func newkmeshPosUV(positions []float32, uvs []float32, elements []uint32) (kmesh, error) {
	// positions data validation
	if foundErr, err := validatemeshDataComponentLength(meshDataComponentPosition, len(positions)); foundErr {
		return kmesh{}, err
	}

	// uvs data validation
	if foundErr, err := validatemeshDataComponentLength(meshDataComponentUV, len(uvs)); foundErr {
		return kmesh{}, err
	}
	if foundErr, err := validateVertexDataMatchToPosition(meshDataComponentUV, uvs, positions); foundErr {
		return kmesh{}, err
	}

	// indices data validation
	if foundErr, err := validatemeshDataComponentLength(meshDataComponentElements, len(elements)); foundErr {
		return kmesh{}, err
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

	posUVvbo := createFloat32VBO(perVertexDataCount,
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
	iModelMatrixVBO := createFloat32InstanceVBO(16,
		[]VBOData{
			//col0
			VBOData{
				count:      4,
				byteOffset: 0,
			},
			//col1
			VBOData{
				count:      4,
				byteOffset: 4 * 4,
			},
			//col2
			VBOData{
				count:      4,
				byteOffset: 8 * 4,
			},
			//col3
			VBOData{
				count:      4,
				byteOffset: 12 * 4,
			},
		},
		make([]float32, 64),
	)
	vao := createVAO()
	vao.storeVBO(posUVvbo)
	vao.storeInstanceVBO(iModelMatrixVBO)
	vao.storeEBO(elements)
	return kmesh{
		id:          kmeshID(vao.id),
		vao:         vao,
		elementSize: int32(len(elements)),
	}, nil
}

func newkmeshPosUVNormals(positions []float32, uvs []float32, normals []float32, elements []uint32) (kmesh, error) {
	// positions data validation
	if foundErr, err := validatemeshDataComponentLength(meshDataComponentPosition, len(positions)); foundErr {
		return kmesh{}, err
	}

	// uvs data validation
	if foundErr, err := validatemeshDataComponentLength(meshDataComponentUV, len(uvs)); foundErr {
		return kmesh{}, err
	}
	if foundErr, err := validateVertexDataMatchToPosition(meshDataComponentUV, uvs, positions); foundErr {
		return kmesh{}, err
	}

	// 	// normals data validation
	if foundErr, err := validatemeshDataComponentLength(meshDataComponentNormal, len(normals)); foundErr {
		return kmesh{}, err
	}
	if foundErr, err := validateVertexDataMatchToPosition(meshDataComponentNormal, normals, positions); foundErr {
		return kmesh{}, err
	}

	// indices data validation
	if foundErr, err := validatemeshDataComponentLength(meshDataComponentElements, len(elements)); foundErr {
		return kmesh{}, err
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

	posUVNormvbo := createFloat32VBO(perVertexDataCount,
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
	iModelMatrixVBO := createFloat32InstanceVBO(16,
		[]VBOData{
			//col0
			VBOData{
				count:      4,
				byteOffset: 0,
			},
			//col1
			VBOData{
				count:      4,
				byteOffset: 4 * 4,
			},
			//col2
			VBOData{
				count:      4,
				byteOffset: 8 * 4,
			},
			//col3
			VBOData{
				count:      4,
				byteOffset: 12 * 4,
			},
		},
		make([]float32, 64),
	)
	vao := createVAO()
	vao.storeVBO(posUVNormvbo)
	vao.storeInstanceVBO(iModelMatrixVBO)
	vao.storeEBO(elements)
	return kmesh{
		id:          kmeshID(vao.id),
		vao:         vao,
		elementSize: int32(len(elements)),
	}, nil
}

func validatemeshDataComponentLength(meshDataComponent meshDataComponent, vertexDataToValidateLength int) (foundError bool, err error) {
	var shouldBeDivisibleBy int
	var dataName string

	switch meshDataComponent {
	case meshDataComponentPosition:
		shouldBeDivisibleBy = 3
		dataName = "POSITION"
	case meshDataComponentUV:
		shouldBeDivisibleBy = 2
		dataName = "UV"
	case meshDataComponentNormal:
		shouldBeDivisibleBy = 3
		dataName = "NORMAL"
	case meshDataComponentElements:
		shouldBeDivisibleBy = 3
		dataName = "INDICES"
	}
	if vertexDataToValidateLength%shouldBeDivisibleBy != 0 {
		return true, fmt.Errorf("\n***\t%s length should be divisible by %d\n***\t%s length found: %d", dataName, shouldBeDivisibleBy, dataName, vertexDataToValidateLength)
	}
	return false, nil
}

func validateVertexDataMatchToPosition(meshDataComponent meshDataComponent, vertexDataToValidate []float32, positionData []float32) (foundError bool, err error) {
	var dataName string
	var toValidateVertexLen int
	positionVertexLen := len(positionData) / 3
	switch meshDataComponent {
	case meshDataComponentUV:
		toValidateVertexLen = len(vertexDataToValidate) / 2
		dataName = "UV"
	case meshDataComponentNormal:
		toValidateVertexLen = len(vertexDataToValidate) / 3
		dataName = "NORMAL"
	}
	if toValidateVertexLen != positionVertexLen {
		return true, fmt.Errorf("\n***\t%s VERTEX length should be match with POSITION VERTEX length\n***\t%s VERTEX length : %d\n***\tPOSITION VERTEX length : %d", dataName, dataName, toValidateVertexLen, positionVertexLen)
	}
	return false, nil
}
