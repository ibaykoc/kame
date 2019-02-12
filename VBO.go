package kame

import "github.com/go-gl/gl/v4.1-core/gl"

type VBO struct {
	id uint32
}

func createVBO() VBO {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	return VBO{
		id: vboID,
	}
}

func (vbo *VBO) bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.id)
}
func (vbo *VBO) unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
func (vbo *VBO) dispose() {
	gl.DeleteBuffers(1, &vbo.id)
}
