package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func Delete(Books []Books, Key int) error {
	Lock := true
	for i := 0; i < len(Books); i++ {
		if Books[i].Bookid == Key {
			Books[i].Available = false
			Lock = false
			println(Books[i].Available)
		} else if i == len(Books)-1 && Lock {
			return errors.New("Book is not on the list")
		}
	}

	byteArray, err := json.Marshal(Books)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("books.json", byteArray, 0644)
	return nil
}
