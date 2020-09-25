package main

import (
	"belajar-golang/auth"
	"fmt"
)

func main() {
	students := []string{"baren", "didi", "kudil"}
	fmt.Println(auth.Login("rian", students))
}
