package kame

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VAO struct {
	id            uint32
	attributeSize uint32
	vboIDs        []uint32
}

func createVAO() VAO {
	var vaoID uint32
	gl.GenVertexArrays(1, &vaoID)
	return VAO{
		id: vaoID,
	}
}

func (vao *VAO) storeVBO(vbo VBO) {
	vao.bind()
	vbo.bind()
	for _, vboData := range vbo.data {
		gl.VertexAttribPointer(vao.attributeSize, vboData.count, vbo.dataType, false, vbo.stride, gl.PtrOffset(vboData.byteOffset))
		vao.attributeSize++
	}
	vao.unbind()
	vbo.unbind()
	vao.vboIDs = append(vao.vboIDs, vbo.id)
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

func (vao *VAO) use() {
	vao.bind()
	for attrID := uint32(0); attrID < vao.attributeSize; attrID++ {
		gl.EnableVertexAttribArray(attrID)
	}
}

func (vao *VAO) unuse() {
	for attrID := uint32(0); attrID < vao.attributeSize; attrID++ {
		gl.DisableVertexAttribArray(attrID)
	}
	vao.unbind()
}

func (vao *VAO) bind() {
	gl.BindVertexArray(vao.id)
}

func (vao *VAO) unbind() {
	gl.BindVertexArray(0)
}

func (vao *VAO) dispose() {
	for _, vboID := range vao.vboIDs {
		gl.DeleteBuffers(1, &vboID)
	}
	gl.DeleteVertexArrays(1, &vao.id)
}
