package main

import "fmt"

type FilterType func(string, []string)string

func filterName(name string, kataKasar []string) string {
	var newName string
	for i := 0;i<len(kataKasar);i++{
		if name == kataKasar[i]{
			newName = "***"
			break
		}else if i == len(kataKasar) - 1{
			newName = name
			break
		}
	}
	return newName
}

func main() {

	getFullName := func (kataKasar []string, name string, filter FilterType) string{
		filteredName := filter(name, kataKasar)
		return filteredName
	}

	kata_kasar := []string{"bodoh","anjay","kera","sial"}
	hasil_filter := getFullName(kata_kasar,"baren", filterName)
	fmt.Printf("nama : %v\n", hasil_filter)
}