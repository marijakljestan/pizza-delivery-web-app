package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/marijakljestan/golang-web-app/client/dto"
	"net/http"
)

func ListMenu() {
	url := baseUrl + "/pizza"
	resp, err := http.Get(url)
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

func CreateOrder() {
	var items []dto.OrderItem
	var pizzaName string
	var quantity int
	var orderItem dto.OrderItem
	fmt.Println("-------- Add items to your order ---------")
Loop:
	for {
		fmt.Println("--- Press any key to continue --")
		fmt.Println("--If you want to finish ordering press 0--")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "0":
			fmt.Println("Finishing order...")
			break Loop
		default:
			fmt.Println("Pizza name: ")
			fmt.Scan(&pizzaName)

			fmt.Println("Quantity: ")
			fmt.Scan(&quantity)

			orderItem = dto.OrderItem{
				PizzaName: pizzaName,
				Quantity:  quantity,
			}
			items = append(items, orderItem)
		}
	}

	reqBody, err := json.Marshal(map[string]any{
		"items": items,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	
	url := baseUrl + "/order"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var createOrderResponse dto.CreateOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&createOrderResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully created order!")
	formatAndPrintResponse(createOrderResponse)
}

func CheckOrderStatus() {
	var orderId string
	fmt.Println("Enter order id:")
	fmt.Scan(&orderId)

	url := fmt.Sprintf(baseUrl+"/order/status/%s", orderId)
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

func CancelOrder() {
	var orderId string
	fmt.Println("Enter order id:")
	fmt.Scan(&orderId)

	url := fmt.Sprintf(baseUrl+"/order/cancel/%s", orderId)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
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
