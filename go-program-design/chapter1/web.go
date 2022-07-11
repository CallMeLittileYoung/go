package chapter1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//生命两个常量
var mu sync.Mutex
var count int

func InitWeb() {
	http.HandleFunc("/", handler) //所有请求都会走入该方法
	http.HandleFunc("/gif", gif2)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("127.0.0.1:10086", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Println("Again")
	mu.Unlock()
	fmt.Fprintf(w, "URL.path is %q \n", r.URL.Path)
}

func gif2(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "counter is %d ", count)
	mu.Unlock()

}
