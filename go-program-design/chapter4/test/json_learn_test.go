package test

import (
	"encoding/json"
	"fmt"
	"go-program-design/chapter4"
	"testing"
)

func Test_Json(t *testing.T) {

	movie := chapter4.Movie{
		Title:    "aa",
		Released: "bb",
		Actors:   []string{"aa", "b"},
	}

	indent, _ := json.MarshalIndent(movie, "", " ")
	fmt.Printf("%s\n", indent)
}
