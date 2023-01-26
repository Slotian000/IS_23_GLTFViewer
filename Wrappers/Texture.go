package Wrappers

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	"os"
)

type Texture struct {
	ID uint32
}

func NewTextureFromFile(path string) (Texture, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return Texture{}, err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return Texture{}, err
	}

	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	return NewTexture(image.NewRGBA(img.Bounds()).Pix, width, height)
}

func NewTexture(data []uint8, width int, height int) (Texture, error) {
	texture := Texture{}

	gl.GenTextures(1, &texture.ID)
	gl.BindTexture(gl.TEXTURE_2D, texture.ID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(width), int32(height), 0, gl.RGB, gl.UNSIGNED_BYTE, gl.Ptr(&data))
	return texture, nil
}
