package main

import (
	"fmt"
	"reflect"
)

func checkTipeData(data interface{}) {
	check := reflect.TypeOf(data)
	fmt.Println(check)
}

func main() {
	nama := "baren maulana"
	umur := 17
	reflectValue := reflect.ValueOf(umur)

	fmt.Println("tipe data pada variable umur :", reflectValue.Type())
	fmt.Println("value data pada variable umur :", reflectValue.Interface())

	checkTipeData(umur)
	checkTipeData(nama)
}
