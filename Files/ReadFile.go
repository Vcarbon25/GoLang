package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("Hello_World.txt")
	if err != nil {
		log.Fatal(err)
	}
	//defer file.close()
	fmt.Println(string(file))

}
