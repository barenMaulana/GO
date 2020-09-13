package main

import "fmt"

type Students struct{
	Name string
	Expertise string	
	Grade int
}

func (student Students) absen(name string)(absen bool){
	if name == student.Name{
		absen = true
	}else{
		absen = false
	}
	return
}

func main(){

	var student Students
	student.Name = "Baren"
	student.Expertise = "Programming"
	student.Grade = 12
	fmt.Println(student)

	fmt.Println(student.absen("Baren"))

}