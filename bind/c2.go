package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	network   string = "tcp"
	port      string = "9000"
	host      string = "localhost"
	seperator string = ":"
	hacker    string = "sadsec"
	symbol1          = "@"
	symbol2          = "$ "
)

func main() {
	for {
		buf := make([]byte, 2048)
		conn, err := net.Dial(network, host+seperator+port)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(hacker + symbol1 + conn.RemoteAddr().String() + symbol2)
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}

		nbytes, err := conn.Read(buf[:])
		if err == io.EOF {
			//nop
		} else if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(buf[:nbytes]))

		conn.Close()
	}
}
