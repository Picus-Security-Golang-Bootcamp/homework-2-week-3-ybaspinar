package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"
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

var (
	search    = flag.String("search", "", "Book Title")
	list      = flag.Bool("list", false, "Book List")
	get       = flag.Int("get", -1, "Book ID")
	deleteKey = flag.Int("delete", -1, "Book ID")
	buy       = flag.Int("buy", -1, "Book ID")
	amount    = flag.Int("amount", -1, "Book ID")
)
var usage = `Usage: go run main.go -search god of flies -delete 1 -buy 1 -amount 4 ...
Options:
	-list    List all the books. 
	-search  Search given words in list.
	-get     Search given ıd in list.
	-delete  Delete given id from list.
	-buy     Buy given id from list. needs amount.
	-amount  Used for buy command.
`

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	set := true
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

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

	if isFlagPassed("list") {
		List(Books)
		set = false
	}
	if isFlagPassed("search") {
		Search(Books, *search)
		set = false

	}
	if isFlagPassed("get") {
		Get(Books, *get)
		set = false

	}
	if isFlagPassed("buy") {
		Buy(Books, *buy, *amount)
		set = false

	}
	if isFlagPassed("delete") {
		Delete(Books, *deleteKey)
		set = false

	}
	if set {
		usageAndExit()
	}

}
func usageAndExit() {
	fmt.Fprintf(os.Stderr, "\n\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

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
		} else if Key == -1 {
			return errors.New("Please give available ıd")
		}
	}
	byteArray, err := json.Marshal(Books)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("books.json", byteArray, 0644)
	return nil
}
func Delete(Books []Books, Key int) error {
	Lock := true
	for i := 0; i < len(Books); i++ {
		if Books[i].Bookid == Key {
			Books[i].Available = false
			Lock = false
			println(Books[i].Available)
		} else if i == len(Books)-1 && Lock {
			return errors.New("Book is not on the list")
		} else if Key == -1 {
			return errors.New("Please give available ıd")
		}
	}

	byteArray, err := json.Marshal(Books)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("books.json", byteArray, 0644)
	return nil
}
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
		} else if Key == -1 {
			return errors.New("Please give available ıd")
		}
	}
	return nil
}
