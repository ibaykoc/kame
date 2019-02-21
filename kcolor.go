package kame

type kcolorID int

type Kcolor struct {
	R float32
	G float32
	B float32
	A float32
}

func (kcolor *Kcolor) RGBA() (r, g, b, a float32) {
	return kcolor.R, kcolor.G, kcolor.B, kcolor.A
}
