package chapter1

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchUrl() {

	var s string = "https://www.google.com.hk/?hl=zh-cn"
	get, err := http.Get(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err, err)
		os.Exit(-1)
	}
	code := get.StatusCode
	fmt.Printf("code is %v \n", code)

	stdout := os.Stdout
	//io.Copy(stdout)
	// io.copy 新函数
	_, err = io.Copy(stdout, get.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err, err)
		os.Exit(-1)
	}
	//all, err := ioutil.ReadAll(get.Body)
	//get.Body.Close()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", s, err)
	//	os.Exit(-1)
	//}
	//fmt.Printf("%s", all)
}
