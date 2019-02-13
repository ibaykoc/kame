package kame

type DrawableModel struct {
	vao        VAO
	textureID  uint32
	vertexSize int32
}

func CreateDrawableModel(window *Window, vertexPositionUVs []float32, indices []uint32, texturePath string) DrawableModel {
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
		textureID:  tID,
		vertexSize: int32(len(indices)),
	}
}

func (dm *DrawableModel) Dispose() {
	dm.vao.dispose()
}
