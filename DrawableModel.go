package kame

type DrawableModel struct {
	vao        VAO
	vertexSize int32
}

func CreateDrawableModel(window *Window, vertexPositions []float32) DrawableModel {
	window.MakeContextCurrent()
	vao := createVAO()
	vao.storeFloat32Buffer(0, 3, vertexPositions)
	return DrawableModel{
		vao:        vao,
		vertexSize: int32(len(vertexPositions)) / 3,
	}
}

func (dm *DrawableModel) Dispose() {
	dm.vao.dispose()
}
