package ports

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	"github.com/jialechen7/gorder-v2/common/tracing"
	"github.com/jialechen7/gorder-v2/order/app"
	"github.com/jialechen7/gorder-v2/order/app/command"
	"github.com/jialechen7/gorder-v2/order/app/query"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{app: app}
}

func (H HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {
	ctx, span := tracing.Start(c, "PostCustomerCustomerIDOrders")
	defer span.End()

	var req orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := H.app.Commands.CreateOrder.Handle(ctx, command.CreateOrder{
		CustomerID: req.CustomerID,
		Items:      req.Items,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	traceID := tracing.TraceID(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message":      "OK",
		"trace_id":     traceID,
		"customer_id":  req.CustomerID,
		"order_id":     r.OrderID,
		"redirect_url": fmt.Sprintf("http://localhost:8282/success?customerID=%s&orderID=%s", req.CustomerID, r.OrderID),
	})
}

func (H HTTPServer) GetCustomerCustomerIDOrdersOrderID(c *gin.Context, customerID string, orderID string) {
	ctx, span := tracing.Start(c, "GetCustomerCustomerIDOrdersOrderID")
	defer span.End()

	o, err := H.app.Queries.GetCustomerOrder.Handle(ctx, query.GetCustomerOrder{
		OrderID:    orderID,
		CustomerID: customerID,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	traceID := tracing.TraceID(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"trace_id": traceID,
		"data": gin.H{
			"Order": o,
		},
	})
}
