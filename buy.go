package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
)

func Buy(Books []Books, Key int, Amount int) error {
	Lock := true
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer w.Flush()
	fmt.Fprint(w, "Bookid\tBooktitle\tPrice\tPages\tStockamount\tStockid\tIsbn\tAuthorname\tAuthorid\t\n")
	for i := 0; i < len(Books); i++ {
		if Books[i].Bookid == Key && Books[i].Available {
			if Books[i].Stockamount < Amount {
				return errors.New("Not enough books available")
			} else {
				Books[i].Stockamount = Books[i].Stockamount - Amount
			}
			fmt.Fprintf(w, "%d\t%q\t%f\t%d\t%d\t%d\t%d\t%q\t%d\n", Books[i].Bookid, Books[i].Booktitle, Books[i].Price, Books[i].Pages, Books[i].Stockamount, Books[i].Stockid, Books[i].Sbn, Books[i].Author.Authorname, Books[i].Author.Authorid)
			Lock = false
		} else if Books[i].Bookid == Key {
			return errors.New("Book is deleted")
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
