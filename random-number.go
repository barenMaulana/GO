package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumber(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(50-10)+10)
}

func main(){
	randNumber()
}