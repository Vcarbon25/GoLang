package main

/*
don't need to import the second file, is already acessible when running module as whole
to run these packages at the same time use in the terminal for the folder
		go run .
	instead of go run Principal.go
*/

import (
	"fmt"
	S ".second"
)

func main() {
	fmt.Println("Main code\n")
	a := S.Imprime("Enviado do programa principal")
	fmt.Println(a)
	S.Chama2("it worked")
}
