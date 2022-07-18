package test

import (
	"crypto/sha256"
	"fmt"
	"go-program-design/chapter4"
	"testing"
)

func Test_arr(t *testing.T) {
	chapter4.Arr()

	sum224 := sha256.Sum224([]byte("X"))
	sum225 := sha256.Sum224([]byte("x"))
	fmt.Println(sum224)
	fmt.Println(sum225)
}

func Test_(t *testing.T) {
	chapter4.Sl()
}

func Test_Reverse(t *testing.T) {
	var s [6]int = [6]int{0, 1, 2, 3, 4, 5}
	chapter4.Reverse(s[:])
	fmt.Println(s)
}

func Test_left_move(t *testing.T) {
	//左移两位
	var s [6]int = [6]int{0, 1, 2, 3, 4, 5}
	chapter4.Reverse(s[:2])
	chapter4.Reverse(s[2:])
	chapter4.Reverse(s[:])
	fmt.Println(s)
}

func Test_right_move(t *testing.T) {
	//左移两位
	var s [6]int = [6]int{0, 1, 2, 3, 4, 5}
	chapter4.Reverse(s[:])
	chapter4.Reverse(s[:2])
	chapter4.Reverse(s[2:])
	fmt.Println(s)
}
