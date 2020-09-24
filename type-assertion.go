package main

import (
	"fmt"
)

func randomType()(name interface{}){
	 name = "baren"
	 return name
}

func main(){
	var result interface{} = randomType()

	switch result.(type){
	case string:
		fmt.Println("value",result,"type string")
	case int:
		fmt.Println("value",result,"type int")
	case bool:
		fmt.Println("value",result,"type bool")
	default:
		fmt.Println("value",result,"type unknow")
	} 

}	