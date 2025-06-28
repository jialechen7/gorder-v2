package app

import "github.com/jialechen7/gorder-v2/order/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct{}

type Queries struct {
	GetCustomerOrderHandler query.GetCustomerOrderHandler
}
