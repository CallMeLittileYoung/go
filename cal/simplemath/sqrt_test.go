package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	v := Sqrt(9)
	if v != 3 {
		t.Errorf("error result error msg: %d", v)
	}
}
