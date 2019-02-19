package kame

type ComponentsCreator interface {
	CreateComponents()
}

type ComponentsOwner interface {
	GetComponentPointers() []*Component
}

type IDReceiver interface {
	ReceiveID(id int)
}

type IDOwner interface {
	GetID() int
}
type Entity interface {
	IDOwner
	IDReceiver
	ComponentsCreator
	ComponentsOwner
}

// type Entity struct {
// 	modelMatrix   mgl.Mat4
// 	DrawableModel *DrawableModel
// }
//
// func CreateEntity(dm *DrawableModel) Entity {
// 	return Entity{
// 		modelMatrix:   mgl.Translate3D(0, 0, 0),
// 		DrawableModel: dm,
// 	}
// }

// func (e *Entity) GetPosition() (x, y, z float32) {
// 	px, py, pz, _ := e.modelMatrix.Col(3).Elem()
// 	return px, py, pz
// }

// func (e *Entity) GetRotation() (x, y, z float32) {
// 	rMat := e.GetRotationMat()
// 	m00 := float64(rMat.At(0, 0))
// 	m10 := float64(rMat.At(1, 0))
// 	m11 := float64(rMat.At(1, 1))
// 	m12 := float64(rMat.At(1, 2))
// 	m20 := float64(rMat.At(2, 0))
// 	m21 := float64(rMat.At(2, 1))
// 	m22 := float64(rMat.At(2, 2))
// 	sy := math.Sqrt(m00*m00 + m10*m10)

// 	singular := sy < 1e-6 // If

// 	var rx, ry, rz float64
// 	if !singular {
// 		rx = math.Atan2(m21, m22)
// 		ry = math.Atan2(m20, sy)
// 		rz = math.Atan2(m10, m00)
// 	} else {
// 		rx = math.Atan2(-m12, m11)
// 		ry = math.Atan2(-m20, sy)
// 		rz = 0
// 	}
// 	return mgl.RadToDeg(float32(rx)), mgl.RadToDeg(float32(ry)), mgl.RadToDeg(float32(rz))
// }

// func (e *Entity) GetScale() (x, y, z float32) {
// 	return e.modelMatrix.Col(0).Vec3().Len(), e.modelMatrix.Col(1).Vec3().Len(), e.modelMatrix.Col(2).Vec3().Len()
// }

// func (e *Entity) GetPositionMat() mgl.Mat4 {
// 	px, py, pz := e.GetPosition()
// 	return mgl.Translate3D(px, py, pz)
// }

// func (e *Entity) GetRotationMat() mgl.Mat4 {
// 	sx, sy, sz := e.GetScale()
// 	col0 := e.modelMatrix.Col(0).Vec3().Mul(1 / sx)
// 	col1 := e.modelMatrix.Col(1).Vec3().Mul(1 / sy)
// 	col2 := e.modelMatrix.Col(2).Vec3().Mul(1 / sz)
// 	rMat := mgl.Ident4()
// 	rMat.SetCol(0, col0.Vec4(0))
// 	rMat.SetCol(1, col1.Vec4(0))
// 	rMat.SetCol(2, col2.Vec4(0))
// 	return rMat
// }

// func (e *Entity) GetScaleMat() mgl.Mat4 {
// 	sx, sy, sz := e.GetScale()
// 	return mgl.Scale3D(sx, sy, sz)
// }

// func (e *Entity) SetPosition(x, y, z float32) {
// 	e.modelMatrix.SetCol(3,
// 		mgl.Vec4{x, y, z, 1})
// }

// func (e *Entity) Move(x, y, z float32) {
// 	px, py, pz := e.GetPosition()
// 	e.SetPosition(px+x, py+y, z+pz)
// }

// func (e *Entity) MoveRel(x, y, z float32) {
// 	e.modelMatrix = e.modelMatrix.Mul4(mgl.Translate3D(x, y, z))
// }

// func (e *Entity) SetRotation(x, y, z float32) {

// 	pMat := e.GetPositionMat()
// 	sMat := e.GetScaleMat()

// 	if x+y+z == 0 {
// 		e.modelMatrix = mgl.Ident4().Mul4(pMat).Mul4(sMat)
// 		return
// 	}

// 	rV := mgl.Vec3{mgl.DegToRad(x), mgl.DegToRad(y), mgl.DegToRad(z)}
// 	rVal := rV.Len()
// 	rAxis := rV.Normalize()

// 	e.modelMatrix = pMat.Mul4(mgl.HomogRotate3D(rVal, rAxis)).Mul4(sMat)
// }

// func (e *Entity) Rotate(x, y, z float32) {
// 	rV := mgl.Vec3{mgl.DegToRad(x), mgl.DegToRad(y), mgl.DegToRad(z)}
// 	rVal := rV.Len()
// 	rAxis := rV.Normalize()
// 	e.modelMatrix = e.modelMatrix.Mul4(mgl.HomogRotate3D(rVal, rAxis))
// }

// func (e *Entity) Scale(x, y, z float32) {

// 	pMat := e.GetPositionMat()
// 	rMat := e.GetRotationMat()

// 	e.modelMatrix = pMat.Mul4(rMat).Mul4(mgl.Scale3D(x, y, z))
// }
