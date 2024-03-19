package main

import (
	"fmt"
	"time"
)

// main.go is the entry point for the application.
// It prints "Hello, World!" to the console and then sleeps for 10 seconds.

func main() {
	fmt.Print("Hello, World!")
	time.Sleep(10 * time.Second)

	// wait for clean ctrl + c clean exit then type exiting
	<-make(chan bool)

	fmt.Println("Exiting")
}
