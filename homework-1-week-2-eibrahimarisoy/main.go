package main

import (
	"fmt"
	"os"
	"strings"
)

type person struct {
	firstName string
	lastName  string
	birth     int
}

func (p *person) personInfo() string {
	return strings.Join([]string{p.firstName, p.lastName, fmt.Sprint(p.birth)}, " ")
}

type book struct {
	name   string
	author person
	genres []string
	year   int
}

func (b *book) bookInfo() {
	fmt.Printf("Name : %s \n", b.name)
	fmt.Printf("Author : %v \n", b.author.personInfo())
	fmt.Printf("Genres : %s \n", b.genres)
	fmt.Printf("Year : %d \n", b.year)
}

var books = []book{
	{
		name:   "Hunger Games",
		author: person{firstName: "Suzanne", lastName: "Collins", birth: 1962},
		genres: []string{"Dystopian", "Science fiction", "Drama", "Action"},
		year:   2008,
	},
	{
		name:   "Animal Farm",
		author: person{firstName: "George", lastName: "Orwell", birth: 1903},
		genres: []string{"Political satire"},
		year:   1945,
	},
	{
		name:   "Alchemist",
		author: person{firstName: "Paulo", lastName: "Coelho", birth: 1947},
		genres: []string{"Novel"},
		year:   1988,
	},
	{
		name:   "Harry Potter and the Order of the Phoenix",
		author: person{firstName: "Joanne Kathleen", lastName: "Rowling", birth: 1965},
		genres: []string{"fantasy"},
		year:   2003,
	},
}

// define console information
const info = `
> list : lists all the books
> search <bookName> : queries a book by name
`

func main() {
	args := os.Args[1:]

	// if the user has not provided any arguments
	if len(args) < 1 {
		return
	}

	switch args[0] {

	case "list":
		// list all the books
		list(books)

	case "search":
		// if the user has not provided <bookName>
		if len(args) < 2 {
			return
		}

		// checks bookName is in the books slice and get index
		book := strings.Join(args[1:], " ")
		if b, e := contains(book, &books); e {
			b.bookInfo()
			return
		}
		fmt.Println("We don't have that book")
		fmt.Printf(info)
	}
	return
}

// list prints all the books
func list(books []book) {
	for _, v := range books {
		v.bookInfo()
		fmt.Println("-", strings.Repeat("-", 50))
	}
}

// contains checks if a book is in the books slice if exists return its index
func contains(book string, books *[]book) (*book, bool) {
	for _, v := range *books {
		if strings.ToLower(v.name) == strings.ToLower(book) {
			return &v, true
		}
	}
	return nil, false
}
