package ports

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jialechen7/gorder-v2/common"
	client "github.com/jialechen7/gorder-v2/common/client/order"
	"github.com/jialechen7/gorder-v2/order/app"
	"github.com/jialechen7/gorder-v2/order/app/command"
	"github.com/jialechen7/gorder-v2/order/app/dto"
	"github.com/jialechen7/gorder-v2/order/app/query"
	"github.com/jialechen7/gorder-v2/order/convertor"
)

type HTTPServer struct {
	common.BaseResponse
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{app: app}
}

func (H HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerID string) {
	var (
		req  client.CreateOrderRequest
		resp dto.CreateOrderResponse
		err  error
	)
	defer func() {
		H.Response(c, err, resp)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}
	r, err := H.app.Commands.CreateOrder.Handle(c.Request.Context(), command.CreateOrder{
		CustomerID: req.CustomerId,
		Items:      convertor.NewItemWithQuantityConvertor().ClientsToEntities(req.Items),
	})
	if err != nil {
		return
	}

	resp = dto.CreateOrderResponse{
		CustomerID:  req.CustomerId,
		OrderID:     r.OrderID,
		RedirectUrl: fmt.Sprintf("http://localhost:8282/success?customerID=%s&orderID=%s", req.CustomerId, r.OrderID),
	}
}

func (H HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerID string, orderID string) {
	var (
		err  error
		resp dto.GetOrderResponse
	)
	defer func() {
		H.Response(c, err, resp)
	}()

	o, err := H.app.Queries.GetCustomerOrder.Handle(c.Request.Context(), query.GetCustomerOrder{
		OrderID:    orderID,
		CustomerID: customerID,
	})
	if err != nil {
		return
	}
	resp.Order = convertor.NewOrderConvertor().EntityToClient(o)
}
