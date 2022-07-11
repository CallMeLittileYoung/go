package chapter1

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//声命一个切片
var palette = []color.Color{color.Black, color.RGBA{G: 255, B: 68, A: 255}}

//生命常量
const (
	whiteIndex = 0
	blackIndex = 1
)

func Gif() {
	create, _ := os.Create("C:\\Users\\Administrator\\Desktop\\a.gif")
	Lissajous(create)
}
func Lissajous(out io.Writer) {
	const (
		cycles  = 5     //完整的x振荡器变化的个数
		res     = 0.001 //角度的分辨率
		size    = 100
		nframes = 64 //动画中的帧数
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}
