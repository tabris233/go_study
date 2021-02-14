// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	go getTime("NewYork", "8001")
	go getTime("Tokyo", "8002")
	go getTime("London", "8003")

	/// TODO  怎么同时打印在一行呢？   这得阻塞了啊。。。
	/// 网友写的也好搓。。。
	time.Sleep(100 * time.Second)
}

func getTime(TZ string, port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
