package main

import (
	"fmt"
	"log"
	"net/http"
)

func fromHandler(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		fmt.Fprint(writer, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(writer, "Post successful")
	name := request.FormValue("name")
	addy := request.FormValue("addy")

	fmt.Fprintf(writer, "name = %s \n", name)
	fmt.Fprintf(writer, "address = %s\n", addy)

}

func helloHandler(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "not suported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "hello")

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", fromHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("serving....")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
