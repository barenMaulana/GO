package main

import (
	"fmt"
)

type Student struct{
	Name,Major string
	Grade int
}

func ChangeName1(student Student){
	student.Name = "Gatot kaca"
}

func ChangeName2(student *Student){
	student.Name = "Gatot kaca"
}

func main(){
	// struct value
	student := Student{"baren maulana","rekayasa perangkat lunak",12}
	fmt.Println(student)
	fmt.Println("")
	
	// value by pass
	ChangeName1(student)
	fmt.Println(student)

	// pointer by pass
	ChangeName2(&student)
	fmt.Println(student)
}