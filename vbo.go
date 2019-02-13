package kame

import "github.com/go-gl/gl/v4.1-core/gl"

type VBO struct {
	id                 uint32
	data               []VBOData
	perVertexDataCount int32
	stride             int32
	dataType           uint32
}

func createFloat32VBO(perVertexDataCount int32, data []VBOData, buffer []float32) VBO {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(buffer)*4, gl.Ptr(buffer), gl.STATIC_DRAW)
	return VBO{
		id:                 vboID,
		data:               data,
		perVertexDataCount: perVertexDataCount,
		stride:             perVertexDataCount * 4,
		dataType:           gl.FLOAT,
	}
}

func (vbo *VBO) bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.id)
}
func (vbo *VBO) unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// func (vbo *VBO) dispose() {
// 	gl.DeleteBuffers(1, &vbo.id)
// }
