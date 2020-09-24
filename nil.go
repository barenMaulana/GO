package main

import "fmt"

func StudentCheck(student map[string]string)map[string]string{
	if(student["Name"] == ""){
		student = nil
	}else{
		fmt.Println("ada")
	}
	return student
}

func main(){
	Students := map[string]string{}

	fmt.Println(Students)

	if(StudentCheck(Students) == nil){
		fmt.Println("Student not found!")
	}else{
		fmt.Printf("hai %v\n",Students["Name"])
	}

}