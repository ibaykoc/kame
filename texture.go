package kame

import (
	"image/png"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type texture struct {
	id uint32
}

func LoadTexture(texturePath string) uint32 {
	textureFile, err := Resource.Open(texturePath)
	if err != nil {
		panic(err)
	}
	defer textureFile.Close()
	image, err := png.Decode(textureFile)
	if err != nil {
		panic(err)
	}

	imgWidth := int32(image.Bounds().Max.X)
	imgHeight := int32(image.Bounds().Max.Y)
	pixels := make([]byte, imgWidth*imgHeight*4)
	pixelIndex := 0
	for y := 0; y < int(imgHeight); y++ {
		for x := 0; x < int(imgWidth); x++ {
			r, g, b, a := image.At(x, y).RGBA()
			pixels[pixelIndex] = byte(r / 256)
			pixelIndex++
			pixels[pixelIndex] = byte(g / 256)
			pixelIndex++
			pixels[pixelIndex] = byte(b / 256)
			pixelIndex++
			pixels[pixelIndex] = byte(a / 256)
			pixelIndex++
		}
	}
	var tID uint32
	gl.GenTextures(1, &tID)
	gl.BindTexture(gl.TEXTURE_2D, tID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, imgWidth, imgHeight, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	return tID
}
