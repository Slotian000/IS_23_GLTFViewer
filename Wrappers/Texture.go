package Wrappers

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	"image/draw"
	"os"
)

type Texture struct {
	ID uint32
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.ID)
}
func (t *Texture) UnBind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func NewTexture(path string) (Texture, error) {
	texture := Texture{}

	imgFile, err := os.Open("Sources/wall.jpg")
	defer imgFile.Close()
	if err != nil {
		return Texture{}, err
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return Texture{}, err
	}

	rect := img.Bounds()
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	gl.GenTextures(1, &texture.ID)
	texture.Bind()

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(width), int32(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	texture.UnBind()

	return texture, nil
}
