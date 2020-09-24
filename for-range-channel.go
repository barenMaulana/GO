package main

import (
	"fmt"
	"runtime"
)

func setName(nama []string, ch chan<- string){
	for i:=0;i<len(nama);i++{
		ch <- nama[i]
	}
	close(ch)
}

func getName(ch <-chan string){
	for name := range ch{
		fmt.Println(name)
	}
}

func main(){
	runtime.GOMAXPROCS(2)

	names := []string{"baren","didi","rian"}
	ch := make(chan string)
	
	go setName(names,ch)
	getName(ch)
}
