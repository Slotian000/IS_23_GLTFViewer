package Utils

import (
	"image"
	"image/draw"
	"math"
	"os"
)

func ReadRGBA(path string) ([]uint8, error) {
	imgFile, err := os.Open("Sources/wall.jpg")
	defer imgFile.Close()
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	return rgba.Pix, nil
}

func Sin(f float32) float32 {
	return float32(math.Sin(float64(f)))
}
func Cos(f float32) float32 {
	return float32(math.Cos(float64(f)))
}
func MinMax(min float32, val float32, max float32) float32 {
	if val > max {
		return max
	}
	if val < min {
		return min
	}
	return val

}
