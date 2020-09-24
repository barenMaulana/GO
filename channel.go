package main

import (
	"fmt"
	"runtime"
)


func main(){
	runtime.GOMAXPROCS(2)
	var pesanSapa = make( chan string )

	menyapa := func (students []string) {
		for _,each := range students{
			data := fmt.Sprintf("hai %s", each)
			pesanSapa <- data
		}
	}
	
	students1 := []string{"baren","didi","rian"}
	students2 := []string{"ajril","pebi","rangga"}
	students3 := []string{"manto","pikri","kupa"}

	go menyapa(students1)
	go menyapa(students2)
	go menyapa(students3)

	for i := 1; i<10; i++{
		var pesani = <- pesanSapa
		fmt.Println(pesani)	
	}

}

