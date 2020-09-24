package main

import (
	"fmt"
	"strings"
	"strconv"
)

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

func gmailGenerate(name string,callback func(string)bool) string {
	var result string
	if mutated := callback(name); mutated{
		result = name + strings.Repeat(strconv.Itoa(int(name[0])), 1) + "@gmail.com"
	}
	return result
}

func main(){
	students := []string{"baren","didi","mbek"}
	search := "baren"
	fullname := getFullName(search,students,filterStr)
	fmt.Println(fullname)

	realName := "rian"
	repeatName := gmailGenerate(realName, func(name string) bool {
		return name == realName
	})
	

	fmt.Println(repeatName)
}