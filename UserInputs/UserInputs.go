package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Type your Input:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The input was: %s\n", input)
	fmt.Println("Insira string sem espaço ")
	var dado string
	//in fmt scan each word separated by space " " will be put in a diferrent variable, bellow only accepts 1
	fmt.Scanln(&dado)
	fmt.Printf("Fmt scan got: %s", dado)

}
