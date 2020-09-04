package main

import "fmt"

func recursive(students []string, siswa string, len int) string {
	var absensi string
	if  len == -1{
		absensi = "tidak ada siswa yang namanya : " + siswa
	}else if students[len] == siswa{
		absensi = "hadir"
	}else {
		return recursive(students,siswa, len - 1)
	}
	return absensi
}

func main(){
	students := []string{"baren","didi"}
	result := recursive(students, "baren", len(students) - 1)

	fmt.Println(result)
}