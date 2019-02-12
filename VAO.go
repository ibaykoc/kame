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

func (vao *VAO) storeVBO(attributeIndex uint32, vbo VBO) {
	vao.bind()
	vbo.bind()
	gl.VertexAttribPointer(attributeIndex, vbo.singleDataSize, vbo.dataType, false, vbo.stride, nil)
	vbo.unbind()
	vao.unbind()
	vao.attributes[attributeIndex] = vbo
}

func (vao *VAO) storeEBO(indices []uint32) {
	vao.bind()
	var eboID uint32
	gl.GenBuffers(1, &eboID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, eboID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	vao.unbind()
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
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
