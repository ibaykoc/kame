package kame

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Rect struct {
	Min mgl32.Vec3
	Max mgl32.Vec3
}

func (r Rect) GetSize() (width, height float32) {
	w := r.Max.X() - r.Min.X()
	h := r.Max.Y() - r.Min.Y()
	return w, h
}
