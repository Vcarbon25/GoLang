package main

//concurrency in go is made with those native function
import (
	"fmt"
	"time"
)

func task(id int, tmp float32) {
	fmt.Printf("task %d started\n ", id)
	NotConcurrent(id)
	time.Sleep(time.Duration(tmp) * time.Second)
	fmt.Printf("Task %d finished in %v seconds\n", id, tmp)
}

func NotConcurrent(id int) {
	fmt.Printf("functio %v Called non concurrent funcion\n", id)
}

func main() {
	go task(1, 4) //this will create a separate processing thread
	go task(2, 2)
	time.Sleep(7 * time.Second)
	fmt.Printf("Main task finished ")
}
