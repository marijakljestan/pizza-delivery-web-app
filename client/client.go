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

	var registeredUser dto.User
	err = json.NewDecoder(resp.Body).Decode(&registeredUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully registered with username:", registeredUser.Username)
}

func Login() {
	var username string
	fmt.Println("Enter username:")
	fmt.Scan(&username)

	var password string
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	url := "http://localhost:8080/user/login"
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

	var respBody dto.LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(respBody.Message)
	fmt.Println("Your authorization token:", respBody.Token)
}

func ListMenu() {
	url := "http://localhost:8080/pizza"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var menu dto.GetMenuResponse
	err = json.NewDecoder(resp.Body).Decode(&menu)
	if err != nil {
		fmt.Println(err)
		return
	}

	formatAndPrintResponse(menu)
}

func CreateOrder() {
	
}

func CheckOrderStatus() {
	var orderId string
	fmt.Println("Enter order id:")
	fmt.Scan(&orderId)

	url := fmt.Sprintf("http://localhost:8080/order/status/%s", orderId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var orderStatusResponse dto.GetOrderStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&orderStatusResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Your order has status:", orderStatusResponse.OrderStatus)
}

func formatAndPrintResponse(response any) {
	jsonData, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
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
	case "2":
		Login()
		handleCustomerChoice()
	case "3":
		Login()
	case "0":
		return 0
	}
	return 1
}

func handleCustomerChoice() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("--------- Choose Operation ----------")
		fmt.Println("-----------------------------------------")
		fmt.Println("1) List the menu")
		fmt.Println("2) Create an order")
		fmt.Println("3) Check order status")
		fmt.Println("4) Cancel order")
		fmt.Println("0) BACK")
		var choice string
		fmt.Scan(&choice)
		//utils.ClearTerminal()
		fmt.Println("-----------------------------------------")
		switch choice {
		case "1":
			ListMenu()
		case "2":
			CreateOrder()
		case "3":
			CheckOrderStatus()
		/*case "4":
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
		*/
		case "0":
			return
		default:
			fmt.Println("INPUT ERROR !")
		}
	}
}
