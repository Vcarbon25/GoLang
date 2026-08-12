package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf(("Iniciou\n"))
	time.Sleep(3 * time.Second)
	fmt.Printf("Slept for 3 seconds\n")
	// Create a timer for 3 seconds
	fmt.Printf("Iniciou non blocking Timer\n")
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop() // Always clean up to prevent memory leaks

	fmt.Println("Waiting for timer to fire...\n")
	fmt.Println("10")
	fmt.Println("9")

	// Block until the timer's channel receives a value
	<-timer.C
	fmt.Println("Timer fired!")
}
