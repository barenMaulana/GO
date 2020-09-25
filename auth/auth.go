package auth

func Login(student string, students []string) interface{} {
	var returnValue interface{}

	for i := 0; i < len(students); i++ {
		if students[i] == student {
			returnValue = students[i] + " hadir"
			break
		} else if i == len(students)-1 {
			returnValue = student + " tidak hadir"
			break
		}
	}

	return returnValue
}
