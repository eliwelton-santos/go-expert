package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Empty

	// Thread 2
	go func() {
		canal <- "Hello World!" // Full
	}()

	// Thread 1
	msg := <-canal // Empty
	fmt.Println(msg)
}
