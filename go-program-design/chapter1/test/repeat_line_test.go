package test

import (
	"go-program-design/chapter1"
	"testing"
)

func Test_FindRepeatLine(t *testing.T) {
	chapter1.FindRepeatLine2()
}

func Test_Gif(t *testing.T) {
	chapter1.Gif()
}

func Test_FetchUrl(t *testing.T) {
	s := "https://draveness.me/golang/"
	chapter1.FetchUrl(s)
}

func Test_Mu(t *testing.T) {

	ss := []string{"https://draveness.me/golang/", "https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-for-range/",
		"https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-panic-recover/"}
	chapter1.FetchMulti(ss)
}

func Test_initWeb(t *testing.T) {
	chapter1.InitWeb()
}
