// Session 1 pre-exercise.

package pics

import (
	"image"
)

type Image struct {
	x, y int
	a [][]uint8
}

func (i Image) At(x, y int) image.Color {
	p := i.a[x][y]
	return image.RGBAColor{p, p, p, 255}
}

func (i Image) ColorModel() image.ColorModel {
	return image.RGBAColorModel
}

func (i Image) Height() int {
	return i.y
}

func (i Image) Width()  int {
	return i.x
}

func Pic(dx, dy int) image.Image {
	pic := make([][]uint8, dy)
	for y := range pic {
		sx := make([] uint8, dx)
		pic[y] = sx
		for x := range(sx) {
			sx[x] = uint8(x*y)
		}
	}
	return Image{x:dx, y:dy, a:pic}
}

func (i Image) ServeHTTP (c *http.Conn, r *http.Request) {

}


