package model

import "fmt"

var students []string

func init() {
	fmt.Println("fungsi inisialisasi dipanggil")
	students = []string{"baren", "didi", "rangga"}
}

func GetStudents() interface{} {
	return students
}
