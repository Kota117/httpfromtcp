package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}
	defer file.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")

	linesChannel := getLinesChannel(file)

	for line := range linesChannel {
		fmt.Println("read:", line)
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer file.Close()
		defer close(lines)

		// Go automatically sets the capacity to the length if not provided.
		// Since we explicitly only use what was overwritten from the read (buffer[:n]), it is ok to declare the buffer once instead of on each iteration of the loop.
		buffer := make([]byte, 8)

		current_line := ""

		for {
			n, err := file.Read(buffer)
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
