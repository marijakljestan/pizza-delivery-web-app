package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/marijakljestan/golang-web-app/client/dto"
	"net/http"
)

var baseUrl = "http://localhost:8080"

func RegisterUser() {
	var username string
	fmt.Println("Enter username:")
	fmt.Scan(&username)

	var password string
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	url := baseUrl + "/user/register"
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

	url := baseUrl + "/user/login"
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
	switch choice {
	case "1":
		RegisterUser()
	case "2":
		Login()
		handleCustomerChoice()
	case "3":
		Login()
		handleAdminChoice()
	case "0":
		return 0
	}
	return 1
}

func handleCustomerChoice() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("----------- Choose Operation ------------")
		fmt.Println("-----------------------------------------")
		fmt.Println("1) List the menu")
		fmt.Println("2) Create an order")
		fmt.Println("3) Check order status")
		fmt.Println("4) Cancel order")
		fmt.Println("0) BACK")
		var choice string
		fmt.Scan(&choice)
		fmt.Println("-----------------------------------------")
		switch choice {
		case "1":
			ListMenu()
		case "2":
			CreateOrder()
		case "3":
			CheckOrderStatus()
		case "4":
			CancelOrder()
		case "0":
			return
		default:
			fmt.Println("INPUT ERROR !")
		}
	}
}

func handleAdminChoice() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("----------- Choose Operation ------------")
		fmt.Println("-----------------------------------------")
		fmt.Println("1) List the menu")
		fmt.Println("2) Add pizza")
		fmt.Println("3) Delete pizza")
		fmt.Println("4) Cancel order")
		fmt.Println("0) BACK")
		var choice string
		fmt.Scan(&choice)
		fmt.Println("-----------------------------------------")
		switch choice {
		case "1":
			ListMenu()
		case "2":
			AddPizza()
		case "3":
			DeletePizza()
		case "4":
			CancelOrderRegardlessStatus()
		case "0":
			return
		default:
			fmt.Println("INPUT ERROR !")
		}
	}
}
