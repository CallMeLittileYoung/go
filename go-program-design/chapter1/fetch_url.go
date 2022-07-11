package chapter1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func FetchUrl() {

	var s string = "https://www.google.com.hk/?hl=zh-cn"
	get, err := http.Get(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(-1)
	}
	all, err := ioutil.ReadAll(get.Body)
	get.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", s, err)
		os.Exit(-1)
	}
	fmt.Printf("%s", all)
}
