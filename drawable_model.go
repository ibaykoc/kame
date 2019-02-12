package kame

type DrawableModel struct {
	vao        VAO
	vertexSize int32
}

func CreateDrawableModel(window *Window, vertexPositions []float32, indices []uint32) DrawableModel {
	window.MakeContextCurrent()
	vao := createVAO()
	vbo := createFloat32VBO(3, vertexPositions)
	vao.storeVBO(0, vbo)
	vao.storeEBO(indices)
	return DrawableModel{
		vao:        vao,
		vertexSize: int32(len(indices)),
	}
}

func (dm *DrawableModel) Dispose() {
	dm.vao.dispose()
}
