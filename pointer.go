package main

import "fmt"

type Student struct{
	Name,Major string
	Grade uint8
}

func main()  {
	var siswa1 = Student{"baren maulana","rekayasa perangkat lunak",12}
	var siswa2 *Student = &siswa1
	var siswa3 *Student = siswa2

	fmt.Println("siswa 1",siswa1)
	fmt.Println("siswa 2",siswa2)
	fmt.Println("siswa 3",siswa3)

	fmt.Println("")
	siswa3.Name = "gatot kaca"

	fmt.Println("siswa 1",siswa1)
	fmt.Println("siswa 2",siswa2)
	fmt.Println("siswa 3",siswa3)
}