package simple_socket

import (
	"bytes"
	"fmt"
	"io"
	net "net"
	"os"
)

func socketLearn() {

	address := "xueyuanjun.com:80"

	con, err := net.Dial("tcp", address)

	checkErr(err)

	_, err = con.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)

	data, err := readFully(con)
	checkErr(err)
	fmt.Println(string(data))
	//	=== RUN   Test_socket_learn
	//	HTTP/1.1 301 Moved Permanently
	//Server: nginx/1.18.0 (Ubuntu)
	//Date: Tue, 14 Jun 2022 10:03:59 GMT
	//	Content-Type: text/html
	//	Content-Length: 178
	//Connection: close
	//Location: https://laravelacademy.org/
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("err ", err)
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	// 读取所有响应数据后主动关闭连接
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
