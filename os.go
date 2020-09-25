package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err.Error())
	}

	arguments := os.Args
	fmt.Println(arguments)
	fmt.Println(arguments[0])
	fmt.Println(arguments[1])
	fmt.Println(arguments[2])

	errChdir := os.Chdir("baren")

	if errChdir != nil {
		fmt.Println("error:", errChdir)
	}

	errChmod := os.Chmod("test-chmod.txt", 555)
	if errChmod != nil {
		fmt.Println("error chmod:", errChmod)
	} else {
		fmt.Println("permission untuk file test-chmod.txt berubah")
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)

	os.Clearenv()

	if pesan, err := os.Executable(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pesan)
	}
}
