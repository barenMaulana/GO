package main

import (
	"fmt"
	"errors"
)

 func auth(name, token string)(string, error){
	if funcToken := "abcd";funcToken != token{
		respond := "autentikasi gagal"
		return respond,errors.New("token salah")
	}else{
		respond := "selamat datang " + name
		return respond, nil
	}
 }

func main(){
	result,err := auth("baren","abcd")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}