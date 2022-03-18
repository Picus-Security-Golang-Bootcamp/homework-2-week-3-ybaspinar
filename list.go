package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func List(Books []Books) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer w.Flush()
	fmt.Fprint(w, "Bookid\tBooktitle\tPrice\tPages\tStockamount\tStockid\tIsbn\tAuthorname\tAuthorid\t\n")
	for i := 0; i < len(Books); i++ {
		if Books[i].Available {
			fmt.Fprintf(w, "%d\t%q\t%f\t%d\t%d\t%d\t%d\t%q\t%d\n", Books[i].Bookid, Books[i].Booktitle, Books[i].Price, Books[i].Pages, Books[i].Stockamount, Books[i].Stockid, Books[i].Sbn, Books[i].Author.Authorname, Books[i].Author.Authorid)
		} else {
			fmt.Fprintf(w, "%d\t%q\t%f\t%d\t%d\t%d\t%d\t%q\t%d\t%q\n", Books[i].Bookid, Books[i].Booktitle, Books[i].Price, Books[i].Pages, Books[i].Stockamount, Books[i].Stockid, Books[i].Sbn, Books[i].Author.Authorname, Books[i].Author.Authorid, "Book is deleted")
		}
	}
}
