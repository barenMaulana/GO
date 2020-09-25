package main

import "fmt"

type studentStruct struct{
	name string
	grade int
}

func ( students studentStruct ) getBio() (string,int) {
	return students.name, students.grade
}

func main(){
	var student studentStruct
	student.name = "baren maulana"
	student.grade = 12

	getName,getGrade := student.getBio()

	fmt.Printf("nama siswa : %v\n", getName)
	fmt.Printf("kelas siswa : %v\n", getGrade)
}