package app

import (
	"github.com/jialechen7/gorder-v2/payment/app/command"
)

type Application struct {
	Commands Commands
}

type Commands struct {
	CreatePayment command.CreatePaymentHandler
}
