package main

import "fmt"

func change(original *int, value int){
	 *original = value
}

func main()  {
	a := 4

	change(&a, 10)
	fmt.Println(a)
}