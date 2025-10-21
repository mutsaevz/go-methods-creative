package main

import (
	"/github.com/mutsaevz/go-methods-creative/methods"
	"fmt"
)

func main() {
	library := methods.Library{}
	fmt.Println(library.Name([]string{"INTOCODE", ""}))
}
