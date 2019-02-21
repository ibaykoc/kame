package kame

// import (
// 	"github.com/go-gl/gl/v4.1-core/gl"
// )

type Kdrawable2d struct {
	Shader    kshaderID
	Mesh      kmeshID
	Texture   ktextureID
	TintColor kcolorID
}

// type kdrawable2dBuilder struct {
// 	shader    kshaderID
// 	mesh      kmeshID
// 	texture   ktextureID
// 	tintColor kcolorID
// }

// func Kdrawable2dBuilder(shader kshaderID, mesh kmeshID, texture ktextureID, tintColor kcolorID) *kdrawable2dBuilder {
// 	return &kdrawable2dBuilder{
// 		shader:    shader,
// 		mesh:      mesh,
// 		texture:   texture,
// 		tintColor: tintColor,
// 	}
// }

// func (kd2db *kdrawable2dBuilder) BuildTo(windowDrawer KwindowDrawerID) kdrawable2d {
// 	d := windows[KwindowID(windowDrawer)].kwindowDrawer
// 	d2d := d.(*KwindowDrawer2D)
// 	d2d.
// 	// return kd2db
// }

// func (kdrawable *kdrawable2d) startDraw() {
// 	gl.ActiveTexture(gl.TEXTURE0)
// 	gl.BindTexture(gl.TEXTURE_2D, uint32(kdrawable.ktextureID))
// 	kdrawable.kmesh.startDraw()
// }

// func (kdrawable *kdrawable2d) stopDraw() {
// 	kdrawable.kmesh.stopDraw()
// 	gl.ActiveTexture(gl.TEXTURE0)
// 	gl.BindTexture(gl.TEXTURE_2D, 0)
// }

// // func newkdrawable2D(vbo VBO, ebo []uint32, textureID ktextureID, tintColor Kcolor) kdrawable2d {
// // 	fmt.Printf("Create new drawable model\n")
// // 	vao := createVAO()
// // 	vao.storeVBO(vbo)
// // 	vao.storeEBO(ebo)
// // 	return kdrawable2d{
// // 		vao:         vao,
// // 		ktextureID:  textureID,
// // 		elementSize: int32(len(ebo)),
// // 		tintColor:   tintColor,
// // 	}
// // }
