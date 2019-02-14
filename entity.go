package kame

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Entity struct {
	Position      mgl.Vec3
	Rotation      mgl.Vec3
	Scale         mgl.Vec3
	DrawableModel *DrawableModel
}

func CreateEntity(position, rotation, scale mgl.Vec3, drawableModel *DrawableModel) Entity {
	return Entity{
		Position:      position,
		Rotation:      rotation,
		Scale:         scale,
		DrawableModel: drawableModel,
	}
}

func (e *Entity) modelMatrix() mgl.Mat4 {
	rAxis := e.Rotation.Normalize()
	rValue := e.Rotation.Len()
	mMat := mgl.Ident4()
	mMat = mMat.Mul4(mgl.Translate3D(e.Position.X(), e.Position.Y(), -e.Position.Z()))
	mMat = mMat.Mul4(mgl.Scale3D(e.Scale.X(), e.Scale.Y(), e.Scale.Z()))
	if rValue != 0 {
		mMat = mMat.Mul4(mgl.HomogRotate3D(mgl.DegToRad(rValue), rAxis))
	}
	return mMat
}
