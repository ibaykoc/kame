package kame

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Krect struct {
	TopLeft, TopRight, BottomLeft, BottomRight, Center mgl32.Vec3
}

func (r Krect) GetSize() (width, height float32) {
	w := r.TopLeft.X() - r.TopRight.X()
	h := r.TopLeft.Y() - r.BottomLeft.Y()
	return w, h
}
