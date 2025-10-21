package main

import (
	"fmt"

	"github.com/mutsaevz/go-methods-creative/methods"
)

func main() {
	l := methods.Library{Name: "INTOCODE", BooksCount: 15, FoundedYear: 1999, Librarian: "Kuduzow Akhmad", ReadersCount: 22, Sections: []string{"Go", "JS"}, IsOpen: true}

	fmt.Println(l)

	l.AddBook(3)
	l.DeleteBook(1)

	l.RegisterReader()
	l.DeleteReader()

	l.AddSection("Python")

	fmt.Println(l.IsSectionExists("Go"))

	l.Close()
	l.Open()

	l.Info()

	l.ChangeLibrarian("intocode")

	fmt.Println(l)
}
