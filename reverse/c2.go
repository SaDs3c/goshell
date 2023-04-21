package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	network   = "tcp"
	host      = "localhost"
	port      = "9000"
	seperator = ":"
	hacker    = "sadsec"
	symbol1   = "@"
	symbol2   = " $ "
)

func main() {
	fmt.Println("Listening on port", port)
	listener, err := net.Listen(network, host+seperator+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 2048)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(hacker + symbol1 + conn.RemoteAddr().String() + symbol2)
		reader, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err == io.EOF {
			//NOP
		} else if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write(reader)
		if err != nil {
			fmt.Println(err)
			return
		}

		nbytes, err := conn.Read(buf)
		if err == io.EOF {
			//NOP
		} else if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(buf[:nbytes]))
	}

}
