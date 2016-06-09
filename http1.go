package main

import (
	"io"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	fmt.Println("Hi Avraham1")
	http.HandleFunc("/", hello)
	fmt.Println("Hi Avraham2")
		
	err:=http.ListenAndServe("", nil)
	fmt.Println("Hi Avraham3")
	fmt.Println(err)
	fmt.Println("Hi Avraham4")
}
