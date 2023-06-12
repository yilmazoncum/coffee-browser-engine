package rasterizer

import (
	"image"
	"image/color"
	"image/png"
	"main/engine"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func CreateCanvas() *image.RGBA {
	width := 2000
	height := 1000

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}
	return img
}

func Paint(img *image.RGBA, node *engine.LayoutNode) *image.RGBA {
	var queue Queue

	for _, child := range node.Children {
		queue.Enqueue(child)
	}

	for {
		node := queue.Top()
		//1. Enqueue the Children
		for _, child := range node.Children {
			queue.Enqueue(child)
		}

		//2. Paint Current Node
		color := ParseColor(node.RenderNode.Style["background-color"])

		for x := node.Box.X; x < (node.Box.Width + node.Box.X); x++ {
			for y := node.Box.Y; y < (node.Box.Height + node.Box.Y); y++ {
				img.Set(x, y, color)
			}
		}

		if node.RenderNode.Text != "" {
			drawText(img, node)
		}

		//3.Pop
		queue.Dequeue()
		if queue.IsEmpty() {
			break
		}
	}

	return img
}

func CreateImage(img *image.RGBA) {

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func ParseColor(str string) color.RGBA {
	//rgb(0, 126, 164) :(
	if str == "" {
		return color.RGBA{0, 0, 0, 0}
	}
	tmp := strings.SplitN(str[4:len(str)-1], ",", 3)

	var color color.RGBA

	r, _ := strconv.ParseUint(strings.TrimSpace(tmp[0]), 10, 8)
	color.R = uint8(r)

	g, _ := strconv.ParseUint(strings.TrimSpace(tmp[1]), 10, 8)
	color.G = uint8(g)

	b, _ := strconv.ParseUint(strings.TrimSpace(tmp[2]), 10, 8)
	color.B = uint8(b)

	color.A = 255
	return color

}

func drawText(img *image.RGBA, node *engine.LayoutNode) {

	f := basicfont.Face7x13
	text := node.RenderNode.Text

	// Draw the text on the image
	point := fixed.Point26_6{
		X: fixed.I(node.Box.X) + ((fixed.I(node.Box.Width) / 2) - (font.MeasureString(f, text) / 2)),
		Y: fixed.I(node.Box.Y + (node.Box.Height / 2)),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: f,
		Dot:  point,
	}
	d.DrawString(text)
}
