package chapter8

import (
	"io"
	"log"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

func InitServer() {

	listen, err := net.Listen("tcp", "127.0.0.1:10086")
	if err != nil {
		log.Fatalln(err)
		return
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}
		go handleCon(accept)
	}
}

var x uint64 = 0

func handleCon(conn net.Conn) {
	defer conn.Close()
	addInt32 := atomic.AddUint64(&x, 1)
	formatUint := strconv.FormatUint(addInt32, 10)
	count := 0
	for count < 30 {
		count++
		io.WriteString(conn, formatUint+" "+time.Now().Format("15:04:05\n")) //nolint:errcheck
		time.Sleep(1 * time.Second)
	}

}
