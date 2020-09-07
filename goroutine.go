package main

import (
	"fmt"
	"runtime"
)

func sayHello(jumlah int, name string){
	for i := 0; i < jumlah;i++{
		fmt.Println("Hello", name, i + 1)
	}
}

func main(){
	runtime.GOMAXPROCS(2)
	
	go sayHello(5,"baren")
	sayHello(100,"didi")

	var input string
	fmt.Scanln(&input)
}