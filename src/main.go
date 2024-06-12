package main

import (
	"fmt"

	"net/http"
	"strconv"

	router "init/project/src/infrastructure/routes"

	"github.com/labstack/echo/v4"
)

type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var books = []Book{
	{ID: 1, Isbn: "12345", Title: "Book 1", Author: &Author{FirstName: "John", LastName: "Doe"}},
	{ID: 2, Isbn: "23456", Title: "Book 2", Author: &Author{FirstName: "Jane", LastName: "Doe"}},
}

func getBooks(c echo.Context) error {
	// เปลี่ยนจากการใช้ http.ResponseWriter และ *http.Request เป็น echo.Context

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, books)
}

// ใช้ echo.Context แทน http.ResponseWriter และ *http.Request
func getBook(c echo.Context) error {
	idParam := c.Param("ID")
	ID, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	for _, book := range books {
		if book.ID == ID {
			return c.JSON(http.StatusOK, idParam)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
}

// ใช้ echo.Context แทน http.ResponseWriter
func addBook(c echo.Context) error {
	FirstName := c.FormValue("FirstName")
	first_name := c.FormValue("first_name")

	book := new(Book)

	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	book.ID = len(books) + 1
	books = append(books, *book)

	fmt.Println("request to add book:", c.Echo())
	fmt.Println("request to add FirstName:", book.Author.FirstName)
	fmt.Println("request to add first_name:", first_name)

	// ตรวจสอบว่า first_name เป็น Alice หรือไม่
	if first_name == "Alice" || FirstName == "Alice" {
		return c.JSON(http.StatusCreated, map[string]bool{"is_alice": true})
	}

	return c.JSON(http.StatusCreated, map[string]string{"FirstName": book.Author.FirstName})
}

type Shape interface {
	Area() float64
}

func main() {
	e := echo.New()
	router.UserRoute(e)
	router.ProductRoutes(e)

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	e.GET("/api/books", getBooks /* ใช้ echo.Context แทน http.ResponseWriter และ *http.Request */)
	e.GET("/api/book/:ID", getBook /* ใช้ echo.Context แทน http.ResponseWriter และ *http.Request */)
	e.POST("/api/books", addBook)

	e.Logger.Fatal(e.Start(":1323"))
}
