package chapter5

import (
	"testing"
	"time"
)

func Test_Rocket(t *testing.T) {
	r := &Rocket{}
	time.AfterFunc(10*time.Second, r.Launch)

	time.Sleep(20 * time.Second)
}
