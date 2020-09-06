package main

import "fmt"

type sekolah interface{
	getName() 
	getGrade()
}

type siswaStruct struct{
	name string
	grade int
}


func ( s siswaStruct ) getName(){
	fmt.Println("name : " + s.name)
} 

func ( s siswaStruct ) getGrade(){
	fmt.Println("grade : ",s.grade)
}

func data(sk sekolah){
	sk.getName()
	sk.getGrade()
}

func main(){
	var siswa siswaStruct 
	siswa.name = "baren maulana"
	siswa.grade = 12
	data(siswa)
}


