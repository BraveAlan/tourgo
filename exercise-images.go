package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x+y)/2), uint8((x+y)/2), 255,255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}


