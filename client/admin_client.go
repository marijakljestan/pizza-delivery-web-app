package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/marijakljestan/golang-web-app/client/dto"
	"net/http"
)

func AddPizza() {
	var name string
	fmt.Println("Enter name:")
	fmt.Scan(&name)

	var description string
	fmt.Println("Enter description:")
	fmt.Scan(&description)

	var price float32
	fmt.Println("Enter price:")
	fmt.Scan(&price)

	bearerToken := getAuthorizationToken()
	url := baseUrl + "/pizza"

	reqBody, err := json.Marshal(map[string]any{
		"name":        name,
		"description": description,
		"price":       price,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var respBody dto.MenuResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	formatAndPrintResponse(respBody)
}

func DeletePizza() {
	var pizzaName string
	fmt.Println("Enter name of pizza you want to delete:")
	fmt.Scan(&pizzaName)

	bearerToken := getAuthorizationToken()
	url := fmt.Sprintf(baseUrl+"/pizza/%s", pizzaName)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var menu dto.MenuResponse
	err = json.NewDecoder(resp.Body).Decode(&menu)
	if err != nil {
		fmt.Println(err)
		return
	}

	formatAndPrintResponse(menu)
}

func CancelOrderRegardlessStatus() {
	var orderId string
	fmt.Println("Enter order id:")
	fmt.Scan(&orderId)

	bearerToken := getAuthorizationToken()
	url := fmt.Sprintf(baseUrl+"/order/%s", orderId)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var cancelOrderResponse dto.CancelOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&cancelOrderResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully cancelled order")
	formatAndPrintResponse(cancelOrderResponse)
}

func getAuthorizationToken() string {
	var token string
	fmt.Println("Your authorization token:")
	fmt.Scan(&token)
	bearerToken := "Bearer " + token
	return bearerToken
}