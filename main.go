package main

import (
	"fmt"
	"github.com/serranoarevalo/mydict"
)

func main() {
	//dictionary := mydict.Dictionary{"first": "First words"}
	//fmt.Println(dictionary["first"])
	//
	//result, err := dictionary.Search("second")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(result)
	//}

	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
