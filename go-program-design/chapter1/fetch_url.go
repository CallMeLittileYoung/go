package chapter1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func FetchUrl(s string) {

	//var s string = "https://www.google.com.hk/?hl=zh-cn"
	get, err := http.Get(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(-1)
	}
	code := get.StatusCode
	fmt.Printf("code is %v \n", code)

	stdout := os.Stdout
	//io.Copy(stdout)
	// io.copy 新函数
	_, err = io.Copy(stdout, get.Body)
	defer get.Body.Close() //关闭资源
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
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

func FetchUrlWithChan(s string, ch chan<- string) {

	now := time.Now()
	//var s string = "https://www.google.com.hk/?hl=zh-cn"
	get, err := http.Get(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		ch <- fmt.Sprint(err)
		return
	}
	code := get.StatusCode
	fmt.Printf("code is %v \n", code)

	//io.Copy(stdout)
	// io.copy 新函数  ioutil.Discard 只返回成功  不做任何处理
	nbytes, err := io.Copy(ioutil.Discard, get.Body)
	defer get.Body.Close() //关闭资源
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v ", s, err)
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	cost := time.Since(now).Seconds()
	ch <- fmt.Sprintf("cost: %.2fs get data :%7d  url:%s", cost, nbytes, s)
	//all, err := ioutil.ReadAll(get.Body)
	//get.Body.Close()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", s, err)
	//	os.Exit(-1)
	//}
	//fmt.Printf("%s", all)
}

func FetchMulti(s []string) {
	strings := make(chan string)
	for _, s2 := range s {
		go FetchUrlWithChan(s2, strings)
		//开启协程
	}
	for range s {
		fmt.Println(<-strings)
	}
	fmt.Println("aaaa")
}
