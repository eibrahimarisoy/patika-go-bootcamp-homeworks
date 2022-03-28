# Book Store

This repository contains Book Store application written by Go.
Book store has a lot of book. 
You can "list" all books or "search" for a book(case insensitive) using the following commands.

## Clone the project
```
$ git clone https://github.com/Picus-Security-Golang-Bootcamp/homework-1-week-2-eibrahimarisoy.git
$ cd homework-1-week-2-eibrahimarisoy
```

## Commands
### - list command
```
go run main.go list
```
This command displays all books in the list.

###  - search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Ring: The Return of 
```
This command searches for a book in the list. If the book is found, it displays the book. If the book is not found, it displays the message "We don't have that book".