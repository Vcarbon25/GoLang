package main

import "fmt"

func Imprime(Info string) string {
	return "Funcao Secundaria Recebeu: " + string(Info)
}
func Chama2(arg string) {
	fmt.Printf("Invocou função do 2 programa com argumento %v", arg)
}
