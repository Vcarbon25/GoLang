//whem calling the code in cmd with go run Name, the user can put more arguments after 
//the name and the code will use those
//spaces ' ' separeted values
package main

import (
	"fmt"
	"os"
)

func main() {
	ExecName := os.Args[0]
	// argument 0 always will be the executable path
	fmt.Println(ExecName)
	Args := os.Args[1:]
	fmt.Println(Args)
	//the args are all parameters typed after the program name, separeted by spaces ' '
}
