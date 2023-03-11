package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mongo-crud/handler"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/api/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handler.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handler.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handler.DeleteBook).Methods("DELETE")

	// set our port address
	fmt.Println("server started at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
