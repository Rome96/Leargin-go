package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\n Author: %s\n Pages: %d\n", b.Title, b.Author, b.Pages)
}

// Constructor of Book - devuelve un puntero de tipo Book
func newBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func StructAndMethod() {

	//var myBook = Book{
	//	Title:  "Pepito",
	//	Author: "Pepito",
	//	Pages:  300,
	//}

	myBook := newBook("Mody Dick", "Hernan", 300)

	myBook.PrintInfo()
}
