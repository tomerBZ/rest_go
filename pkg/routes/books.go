package routes

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomerBZ/web/pkg/logs"
	"github.com/tomerBZ/web/pkg/utils"
	"net/http"
	"strconv"
)

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

func booksRouter() http.Handler {
	books = append(books, Book{ID: "1", Isbn: "123123", Title: "First book", Author: &Author{FirstName: "Tomer David", LastName: "Ben Zohar"}})
	books = append(books, Book{ID: "2", Isbn: "112312312323123", Title: "Second book", Author: &Author{FirstName: "Tomer David", LastName: "Ben Zohar"}})
	router := chi.NewRouter()
	router.Get("/", booksIndex)
	router.Get("/{bookID}", getBook)
	router.Delete("/{bookID}", deleteBook)
	router.Post("/", createBook)

	return router
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	bookID := chi.URLParam(request, "bookID")
	for index, book := range books {
		if book.ID == bookID {
			books = append(books[:index], books[index+1:]...)
		}
	}

	utils.ToJson(bookID, writer)
}

func booksIndex(writer http.ResponseWriter, request *http.Request) {
	logs.Error.Println("this is an info", books)
	utils.ToJson(books, writer)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	bookID := chi.URLParam(request, "bookID")
	for _, book := range books {
		if book.ID == bookID {
			utils.ToJson(book, writer)
			return
		}
	}
	utils.ToJson(&Book{}, writer)
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)

	lastID := books[len(books)-1].ID
	i, _ := strconv.Atoi(lastID)
	lastID = strconv.Itoa(i + 1)
	book.ID = lastID
	books = append(books, book)

	utils.ToJson(book, writer)
}
