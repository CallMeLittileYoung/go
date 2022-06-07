package details

import (
	"fmt"
	"strconv"
)

func turn() {
	v1 := uint8(16)
	fmt.Println(v1)

	v2 := uint16(255)
	fmt.Println(v2)

	v3 := 99.99
	v4 := int(v3)
	fmt.Println(v4)

	v5 := "100"
	v6, err := strconv.Atoi(v5)
	if err == nil {
		fmt.Println(v6)
	}
}
