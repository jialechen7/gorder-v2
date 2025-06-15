package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/jialechen7/gorder-v2/order/app"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{app: app}
}

func (H HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) GetCustomerCustomerIDOrdersOrderID(c *gin.Context, customerID string, orderID string) {
	//TODO implement me
	panic("implement me")
}
