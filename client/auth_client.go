package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/marijakljestan/golang-web-app/client/dto"
	"net/http"
)

func RegisterUser() {
	var username string
	fmt.Println("Enter username:")
	fmt.Scan(&username)

	var password string
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	url := "http://localhost:8080/user/register"
	reqBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var user dto.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Body)
	fmt.Println(user)
	fmt.Println("Successfully registered with username:", user)
}

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Println("-------- Library Management CLI ---------")
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("----------- Choose Operation ------------")
		fmt.Println("0) EXIT")
		fmt.Println("1) Register as a customer")
		fmt.Println("2) Login as a customer")
		fmt.Println("3) Login as an admin")
		var flag int
		flag = handlePrompets()
		if flag == 0 {
			break
		}
	}
}

func handlePrompets() int {
	var choice string
	fmt.Scan(&choice)
	//utils.ClearTerminal()
	switch choice {
	case "1":
		RegisterUser()
	/*case "2":
		handleReadersPrompets()
	case "3":
		handleReadersPrompets()*/
	case "0":
		return 0
	}
	return 1
}

func handleUserChoice() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("--------- Choose Operation ----------")
		fmt.Println("-----------------------------------------")
		fmt.Println("1) Get all books")
		fmt.Println("2) Get all books sorted by title")
		fmt.Println("3) Get all books sorted by publication date")
		fmt.Println("4) Search by id")
		fmt.Println("5) Search by title")
		fmt.Println("6) Add book")
		fmt.Println("0) BACK")
		var choice string
		fmt.Scan(&choice)
		//utils.ClearTerminal()
		fmt.Println("-----------------------------------------")
		/*switch choice {
		case "1":
			testPort()
			flag, books := getBooks(HOST + ":" + PORT + "/books")
			if flag != 1 {
				fmt.Println("ERROR !")
			} else {
				books.PrintAll()
			}
		case "2":
			testPort()
			flag, books := getBooks(HOST + ":" + PORT + "/books?sortBy=title")
			if flag != 1 {
				fmt.Println("ERROR !")
			} else {
				books.PrintAll()
			}
		case "3":
			testPort()
			flag, books := getBooks(HOST + ":" + PORT + "/books?sortBy=publication%20date")
			if flag != 1 {
				fmt.Println("ERROR !")
			} else {
				books.PrintAll()
			}
		case "4":
			fmt.Print("Enter book ID: ")
			var id string
			fmt.Scan(&id)
			testPort()
			flag, book := getBook(HOST + ":" + PORT + "/books/search?id=" + id)
			if flag != 1 {
				fmt.Println("There is No book with this ID")
			} else {
				fmt.Println("Book is:\n" + book.ToString() + "\n")
			}
		case "5":
			title := utils.ScanLine("Enter book title: ", "ERROR, Enter Book title: ")
			title = utils.ReplaceURLSpaces(title)
			testPort()
			flag, book := getBook(HOST + ":" + PORT + "/books/search?title=" + title)
			if flag != 1 {
				fmt.Println("There is No book with this title")
			} else {
				fmt.Println("Book is:\n" + book.ToString() + "\n")
			}
		case "6":
			book := scanBook()
			json, _ := json.Marshal(book)
			testPort()
			response, err := http.Post(HOST+":"+PORT+"/books", "application/json", bytes.NewBuffer([]byte(json)))
			if err != nil {
				fmt.Println("ERROR !")
			} else {
				if response.StatusCode != 200 {
					fmt.Println("ERROR !")
				} else {
					fmt.Println("Book added successfully!")
				}
		}
		case "0":
			return
		default:
			fmt.Println("INPUT ERROR !")
		}*/
	}
}
