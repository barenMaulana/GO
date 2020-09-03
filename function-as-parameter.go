package main

import "fmt"

type TypeFunc func(string,[]string)string

func getFullName(search string, name []string, filter TypeFunc)string{
	filteredName := filter(search,name)
	return filteredName
}

func filterStr(name string, students []string) string {
	var result string
	for i:=0;i<len(students);i++{
		 if students[i] == name{
			 result = "ada"
			 break
		 }else if i == len(students) - 1{
			 result = "tidak ada"
			 break
		 }
	}
	return result
}

func main(){
	students := []string{"baren","didi","mbek"}
	search := "baren"
	fullname := getFullName(search,students,filterStr)
	fmt.Println(fullname)
}