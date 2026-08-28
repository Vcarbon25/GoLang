package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("Hello_World.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	d := 5
	fmt.Fprintf(file, "Hello World!!!!  %d\n", d)

}
