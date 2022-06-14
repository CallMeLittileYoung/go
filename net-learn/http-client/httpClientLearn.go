package http_client

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getLearn() {
	resp, err := http.Get("https://xueyuanjun.com")
	if err != nil {
		fmt.Println("get err", err)
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	fmt.Println(code)
	io.Copy(os.Stdout, resp.Body)
}
