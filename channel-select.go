package main

import (
	"fmt"
	"runtime"
)

func FirstName(FirstName string, ch chan string){
	ch <- FirstName
}

func LastName(LastName string,ch chan string){
	ch <- LastName
}


func main(){
	runtime.GOMAXPROCS(2)

	channel1 := make(chan string)	
	channel2 := make(chan string)

	go FirstName("baren",channel1)
	go FirstName("maulana",channel2)

	for i := 0; i < 2; i++{
		select{
		case firstName := <- channel1 :
			fmt.Println("First name :",firstName)
		case lastName := <- channel2 :
			fmt.Println("Last name :",lastName)
		}
	}

}