package kame

import (
	"fmt"
	"image/png"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type KtextureID uint32

type ktexture struct {
	id           KtextureID
	unitLocation uint32
}

func (texture *ktexture) startDraw() {
	gl.ActiveTexture(gl.TEXTURE0 + texture.unitLocation)
	gl.BindTexture(gl.TEXTURE_2D, uint32(texture.id))
}

func (texture *ktexture) stopDraw() {
	gl.ActiveTexture(gl.TEXTURE0 + texture.unitLocation)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func newktextureFromPNG(filePath string) (ktexture, error) {
	fmt.Printf("Load new texture file: %s\n", filePath)
	textureFile, err := os.Open(filePath)
	if err != nil {
		return ktexture{}, err
	}
	defer textureFile.Close()
	image, err := png.Decode(textureFile)
	if err != nil {
		panic(err)
	}

	imgWidth := int32(image.Bounds().Max.X)
	imgHeight := int32(image.Bounds().Max.Y)
	pixels := make([]byte, imgWidth*imgHeight*4)
	// Flip texture (0th index start from bottom left)
	pixelIndex := len(pixels) - 1
	for y := 0; y < int(imgHeight); y++ {
		for x := int(imgWidth) - 1; x >= 0; x-- {
			r, g, b, a := image.At(x, y).RGBA()
			pixels[pixelIndex] = byte(a / 256)
			pixelIndex--
			pixels[pixelIndex] = byte(b / 256)
			pixelIndex--
			pixels[pixelIndex] = byte(g / 256)
			pixelIndex--
			pixels[pixelIndex] = byte(r / 256)
			pixelIndex--
		}
	}
	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, imgWidth, imgHeight, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	ktex := ktexture{
		id:           KtextureID(textureID),
		unitLocation: 0,
	}

	return ktex, nil
}

func loadTexturePixel(pixelsData []byte, width int32, height int32) uint32 {
	fmt.Printf("Load new texture pixel: size(%d, %d)\n", width, height)
	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixelsData))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return textureID
}
