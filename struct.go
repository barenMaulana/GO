package main

import "fmt"

type students struct{
	name string
	grade int8
}

type person struct {
    name string
    age  int
}

func main(){

	var siswa students

	siswa.name = "Baren maulana"
	siswa.grade = 3

	fmt.Printf("struct : %v\n", siswa)
	fmt.Println(siswa.name)
	fmt.Println(siswa.grade)

	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	fmt.Println(allStudents)
	
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}


}