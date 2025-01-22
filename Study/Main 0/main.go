
package main

import (

	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello from VSCode!"))
}

func main() {
	
	fmt.Println("hello world")

	router := httprouter.New()

	router.GET("/", IndexHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
