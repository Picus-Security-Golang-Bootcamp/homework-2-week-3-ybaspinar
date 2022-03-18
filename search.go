package main

import (
	"errors"
	fmt "fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func Search(Books []Books, Key string) error {
	Lock := true
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer w.Flush()
	fmt.Fprint(w, "Bookid\tBooktitle\tPrice\tPages\tStockamount\tStockid\tIsbn\tAuthorname\tAuthorid\t\n")
	for i := 0; i < len(Books); i++ {
		if strings.Contains(strings.ToLower(Books[i].Booktitle), strings.ToLower(Key)) && Books[i].Available {
			fmt.Fprintf(w, "%d\t%q\t%f\t%d\t%d\t%d\t%d\t%q\t%d\n", Books[i].Bookid, Books[i].Booktitle, Books[i].Price, Books[i].Pages, Books[i].Stockamount, Books[i].Stockid, Books[i].Sbn, Books[i].Author.Authorname, Books[i].Author.Authorid)
			Lock = false
		} else if strings.Contains(strings.ToLower(Books[i].Booktitle), strings.ToLower(Key)) {
			fmt.Fprintf(w, "%d\t%q\t%f\t%d\t%d\t%d\t%d\t%q\t%d\t%q\n", Books[i].Bookid, Books[i].Booktitle, Books[i].Price, Books[i].Pages, Books[i].Stockamount, Books[i].Stockid, Books[i].Sbn, Books[i].Author.Authorname, Books[i].Author.Authorid, "Book is deleted")
		} else if i == len(Books)-1 && Lock {
			return errors.New("Book is not on the list")
		}
	}
	return nil
}
func Get(Books []Books, Key int) error {
	Lock := true
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer w.Flush()
	fmt.Fprint(w, "Bookid\tBooktitle\tPrice\tPages\tStockamount\tStockid\tIsbn\tAuthorname\tAuthorid\t\n")
	for i := 0; i < len(Books); i++ {
		if Books[i].Bookid == Key && Books[i].Available {
			fmt.Fprintf(w, "%d\t%q\t%f\t%d\t%d\t%d\t%d\t%q\t%d\n", Books[i].Bookid, Books[i].Booktitle, Books[i].Price, Books[i].Pages, Books[i].Stockamount, Books[i].Stockid, Books[i].Sbn, Books[i].Author.Authorname, Books[i].Author.Authorid)
			Lock = false
		} else if Books[i].Bookid == Key {
			fmt.Fprintf(w, "%d\t%q\t%f\t%d\t%d\t%d\t%d\t%q\t%d\t%q\n", Books[i].Bookid, Books[i].Booktitle, Books[i].Price, Books[i].Pages, Books[i].Stockamount, Books[i].Stockid, Books[i].Sbn, Books[i].Author.Authorname, Books[i].Author.Authorid, "Book is deleted")
		} else if i == len(Books)-1 && Lock {
			return errors.New("Book is not on the list")
		}
	}
	return nil
}
