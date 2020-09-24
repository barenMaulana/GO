package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan int, 2)

	go func() {
		for {
			i := <-message
			fmt.Println("terima data", i)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("kirim data", i)
		message <- i
	}

}
