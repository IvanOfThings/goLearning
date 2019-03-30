package main

import (
	"fmt"
)

func main() {

	ch := make(chan int) // channel que únicamente acepta valores de tipo int (con un tamaño tendríamos un bufferd channel)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	//time.Sleep(1 * time.Second)
}
