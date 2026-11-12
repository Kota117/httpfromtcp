package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		linesChannel := getLinesChannel(conn)

		for line := range linesChannel {
			fmt.Println(line)
		}

		fmt.Println("Connection to", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(conn io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer conn.Close()
		defer close(lines)

		// Go automatically sets the capacity to the length if not provided.
		// Since we explicitly only use what was overwritten from the read (buffer[:n]), it is ok to declare the buffer once instead of on each iteration of the loop.
		buffer := make([]byte, 8)

		current_line := ""

		for {
			n, err := conn.Read(buffer)
			if err != nil {
				if current_line != "" {
					lines <- current_line
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}

			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", current_line, parts[i])
				current_line = ""
			}
			current_line += parts[len(parts)-1]
		}
	}()
	return lines
}
