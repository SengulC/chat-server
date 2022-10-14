package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func simpReader(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n') // read the passed string
	fmt.Println(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	for {
		fmt.Printf("Enter text ->")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		simpReader(conn)
	}
}

// run server 1st
