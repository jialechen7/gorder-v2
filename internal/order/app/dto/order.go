package dto

import client "github.com/jialechen7/gorder-v2/common/client/order"

type CreateOrderResponse struct {
	CustomerID  string `json:"customer_id"`
	OrderID     string `json:"order_id"`
	RedirectUrl string `json:"redirect_url"`
}

type GetOrderResponse struct {
	Order *client.Order `json:"order"`
}
