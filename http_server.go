package main

import(
	"fmt"
	"log"
	"net/http"
)

func hello_world(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func main(){
	http.HandleFunc("/", hello_world)
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
