package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("this is handler")
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	var m = "hello"
	fmt.Println("hello ", m)
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
