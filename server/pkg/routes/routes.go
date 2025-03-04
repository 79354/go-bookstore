package routes

import (
	"github.com/gorilla/mux"
	"go-bookstore/server/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")               // get all books from /books
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")            // add a book at /books
	router.HandleFunc("/books/{bookId}", controllers.GetBookByID).Methods("GET")   // get one book, /books/{id}
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")    // update a book, /book/{id}
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE") // delete a book, /book/{id}
}