package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Writing data on file"))
	// size, err := f.WriteString("Hello world!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created successfully! Size: %d bytes\n", size)

	f.Close()

	/// Reading
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// Reading2
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Deleting file
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
