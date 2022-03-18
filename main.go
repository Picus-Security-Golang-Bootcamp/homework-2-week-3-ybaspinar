package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Books struct {
	Bookid      int    `json:"bookid"`
	Booktitle   string `json:"booktitle"`
	Pages       int    `json:"pages"`
	Stockamount int    `json:"stockamount"`
	Price       int    `json:"price"`
	Stockid     int    `json:"stockid"`
	Sbn         int    `json:"ısbn"`
	Author      struct {
		Authorid   int    `json:"authorid"`
		Authorname string `json:"authorname"`
	} `json:"author"`
}

// List Prints all the books from the
func List(Books []Books) {
	for i := 0; i < len(Books); i++ {
		fmt.Printf("Book %d : %s \n", i+1, Books[i].Booktitle)
	}
}

func main() {
	var Books []Books
	//Reads JSON file
	file, err := ioutil.ReadFile("books.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Decode JSON file to array
	err = json.Unmarshal(file, &Books)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	List(Books)
}
