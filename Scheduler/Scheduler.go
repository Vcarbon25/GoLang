//this code runs and executes a certain function in defined intervals
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	fmt.Println("began program")
  //define the defined period in the ticker argument
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
L:
	for {

		select {
		case t := <-ticker.C:
      //here will insert the code to execute
			fmt.Printf("Executed function %v \n", t)
			i = i + 1

		}
    //run for how many cycles? 
		if i > 5 {
			break L
		}
	}
}
