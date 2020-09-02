package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Ready to GO!</h1>")
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000",nil)
}