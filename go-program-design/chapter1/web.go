package chapter1

import (
	"fmt"
	"log"
	"net/http"
)

func InitWeb() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gif", gif2)
	log.Fatal(http.ListenAndServe("127.0.0.1:10086", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path is %q \n", r.URL.Path)
}

func gif2(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}
