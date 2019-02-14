package kame

type DrawableModel struct {
	vao        VAO
	hasTexture bool
	textureID  uint32
	vertexSize int32
}

func CreateDrawableModelPositions(window *Window, vertexPositions []float32, indices []uint32) DrawableModel {
	window.MakeContextCurrent()
	vao := createVAO()
	vbo := createFloat32VBO(3,
		[]VBOData{
			// Pos
			VBOData{
				count:      3,
				byteOffset: 0,
			},
		},
		vertexPositions)
	vao.storeVBO(vbo)
	vao.storeEBO(indices)
	return DrawableModel{
		vao:        vao,
		hasTexture: false,
		vertexSize: int32(len(indices)),
	}
}

func CreateDrawableModelPositionUVs(window *Window, vertexPositionUVs []float32, indices []uint32, texturePath string) DrawableModel {
	window.MakeContextCurrent()
	vao := createVAO()
	vbo := createFloat32VBO(5,
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
		vertexPositionUVs)
	vao.storeVBO(vbo)
	vao.storeEBO(indices)
	tID := LoadTexture(texturePath)
	return DrawableModel{
		vao:        vao,
		hasTexture: true,
		textureID:  tID,
		vertexSize: int32(len(indices)),
	}
}

func CreateDrawableModelPositionUVNormals(window *Window, vertexPositionUVNormals []float32, indices []uint32, texturePath string) DrawableModel {
	window.MakeContextCurrent()
	vao := createVAO()
	vbo := createFloat32VBO(8,
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
		vertexPositionUVNormals)
	vao.storeVBO(vbo)
	vao.storeEBO(indices)
	tID := LoadTexture(texturePath)
	return DrawableModel{
		vao:        vao,
		hasTexture: true,
		textureID:  tID,
		vertexSize: int32(len(indices)),
	}
}

func (dm *DrawableModel) Dispose() {
	dm.vao.dispose()
}
