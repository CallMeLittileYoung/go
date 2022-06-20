package http_server

import (
	"fmt"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		fmt.Println(params)
		fmt.Fprintf(writer, "你好, %s", params.Get("name"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动 HTTP 服务器失败: %v", err)
	}
}
