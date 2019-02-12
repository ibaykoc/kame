package kame

type DrawableModel struct {
	vao        VAO
	vertexSize int32
}

func CreateDrawableModel(window *Window, vertexPositions []float32) DrawableModel {
	window.MakeContextCurrent()
	vao := createVAO()
	vbo := createFloat32VBO(3, vertexPositions)
	vao.storeVBO(0, vbo)
	return DrawableModel{
		vao:        vao,
		vertexSize: int32(len(vertexPositions)) / 3,
	}
}

func (dm *DrawableModel) Dispose() {
	dm.vao.dispose()
}
