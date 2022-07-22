package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var clockWall = []*string{
	flag.String("NewYork", "localhost:8000", "Connection for US/Eastern clock"),
	flag.String("London", "localhost:8000", "Connection for Europe/London clock"),
	flag.String("Tokyo", "localhost:8000", "Connection for Asia/Tokyo clock"),
}

func main() {
	flag.Parse()
	for i := 0; i < 3; i++ {
		conn, err := net.Dial("tcp", *clockWall[i])
		if err != nil {
			log.Fatal(err)
		}
		go shouldCopy(os.Stdout, conn)
	}
	select {}
}

func shouldCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
