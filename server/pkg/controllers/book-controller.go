package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/79354/go-bookstore/pkg/models"
	"github.com/79354/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

/*

 */

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	b:= newBook.AddBook()
	resp, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}


func GetBookByID(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	ID, err:= strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while Parsing")
		return
	}
	bookDetails, _ := models.BookByID(ID)
	resp, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var newBook models.Book
	utils.ParseBody(r, newBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while Parsing")
		return
	}
	bookDetails, db := models.BookByID(ID)
	if newBook.Name != ""{
		bookDetails.Name = newBook.Name 
	}
	if newBook.Author != ""{
		bookDetails.Author = newBook.Author 
	}
	if newBook.Publication != ""{
		bookDetails.Publication = newBook.Publication
	}
	db.Save(&bookDetails)
	resp, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("unable to parse")
		return
	}
	book := models.RemoveBook(ID)
	resp, _:= json.Marshal(book)
	w.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
