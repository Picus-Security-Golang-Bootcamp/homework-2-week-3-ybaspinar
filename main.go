package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Books struct {
	Bookid      int     `json:"bookid"`
	Booktitle   string  `json:"booktitle"`
	Pages       int     `json:"pages"`
	Stockamount int     `json:"stockamount"`
	Price       float64 `json:"price"`
	Stockid     int     `json:"stockid"`
	Sbn         int     `json:"ısbn"`
	Author      struct {
		Authorid   int    `json:"authorid"`
		Authorname string `json:"authorname"`
	} `json:"author"`
	Available bool `json:"available"`
}

// List Prints all the books from the

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
