package dto

type Pizza struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type Order struct {
	CustomerUsername string      `json:"customer_username"`
	Price            float32     `json:"price"`
	Items            []OrderItem `json:"items"`
}

type OrderItem struct {
	PizzaName string `json:"pizza_name"`
	Quantity  int    `json:"quantity"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type GetMenuResponse struct {
	Menu []Pizza `json:"menu"`
}

type GetOrderStatusResponse struct {
	OrderStatus string `json:"status"`
}
