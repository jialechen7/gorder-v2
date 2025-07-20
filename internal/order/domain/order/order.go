package order

import (
	"errors"
	"fmt"

	"github.com/jialechen7/gorder-v2/order/entity"
	"github.com/stripe/stripe-go/v82"
)

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items       []*entity.Item
}

func NewOrder(ID string, customerID string, status string, paymentLink string, items []*entity.Item) (*Order, error) {
	if ID == "" {
		return nil, errors.New("empty ID")
	}
	if customerID == "" {
		return nil, errors.New("empty customerID")
	}
	if status == "" {
		return nil, errors.New("empty status")
	}
	if items == nil {
		return nil, errors.New("empty items")
	}
	return &Order{ID: ID, CustomerID: customerID, Status: status, PaymentLink: paymentLink, Items: items}, nil
}

func (o *Order) IsPaid() error {
	if o.Status == string(stripe.CheckoutSessionPaymentStatusPaid) {
		return nil
	}
	return fmt.Errorf("order status not paid, order id = %s, status = %s", o.ID, o.Status)
}
