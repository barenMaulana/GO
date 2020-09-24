package main

import "fmt"

type StudentInterface interface{
	GetName() string
}

func sayHello(student StudentInterface){
	fmt.Println("hai", student.GetName())
}

type Students struct{
	Name string
}

func (student Students) GetName() string {
	return student.Name
}

func main(){
	var student Students
	student.Name = "baren"

	sayHello(student)
}