package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VAO struct {
	id         uint32
	attributes map[uint32]VBO
}

func createVAO() VAO {
	var vaoID uint32
	gl.GenVertexArrays(1, &vaoID)
	return VAO{
		id:         vaoID,
		attributes: make(map[uint32]VBO),
	}
}

func (vao *VAO) storeFloat32Buffer(attributeIndex uint32, dataSizePerVertex int32, buffer []float32) {
	vbo := createVBO()
	vao.bind()
	vbo.bind()                                                                    // bind vbo that we are about to store buffer into
	gl.BufferData(gl.ARRAY_BUFFER, len(buffer)*4, gl.Ptr(buffer), gl.STATIC_DRAW) // store data
	gl.VertexAttribPointer(attributeIndex, dataSizePerVertex, gl.FLOAT, false, 0, nil)
	vbo.unbind()
	vao.unbind()
	vao.attributes[attributeIndex] = vbo
}

func (vao *VAO) bind() {
	gl.BindVertexArray(vao.id)
}

func (vao *VAO) unbind() {
	gl.BindVertexArray(0)
}

func (vao *VAO) dispose() {
	for _, vbo := range vao.attributes {
		vbo.dispose()
	}
	gl.DeleteVertexArrays(1, &vao.id)
}
