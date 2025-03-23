package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"net/http"
)

type Book struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	SubTitle    string `json:"subTitle"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	Publisher   string `json:"publisher"`
	Pages       int    `json:"pages"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

var Books = []Book{
	{
		ISBN:        "9781449331818",
		Title:       "Book Title",
		SubTitle:    "Sub Title",
		Author:      "JK",
		PublishDate: "2020-06-04T09:11:40.000Z",
		Publisher:   "Media",
		Pages:       200,
		Description: "With Learning JavaScript Design Patterns",
		Website:     "http://www.addyosmani.com.br",
	},
}

func main() {
	startProvider()
}

func startProvider() {
	router := gin.Default()
	router.GET("/BookStore/v1/Book/ISBN/:isbn", getBookByISBN)

	err := router.Run("localhost:8081")
	if err != nil {
		log.Infof("Failed to start you service")
	}
}

func getBookByISBN(b *gin.Context) {
	for _, book := range Books {
		if book.ISBN == b.Param("isbn") {
			b.Header("Content-Type", "application/json")
			b.JSON(http.StatusOK, book)
			return
		}
	}

	// Return 404 Status Code and error message if no book was found.
	b.Header("Content-Type", "application/json")
	b.JSON(http.StatusNotFound, gin.H{"message": "Requested book is not found"})
}
