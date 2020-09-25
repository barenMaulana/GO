package model

var students []string

func init() {
	students = []string{"baren", "didi", "rangga"}
}

func GetStudents() interface{} {
	return students
}
