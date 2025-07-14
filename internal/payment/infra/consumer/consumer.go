package consumer

import (
	"context"
	"encoding/json"

	"github.com/jialechen7/gorder-v2/common/broker"
	"github.com/jialechen7/gorder-v2/common/genproto/orderpb"
	"github.com/jialechen7/gorder-v2/payment/app"
	"github.com/jialechen7/gorder-v2/payment/app/command"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	app app.Application
}

func NewConsumer(app app.Application) *Consumer {
	return &Consumer{
		app: app,
	}
}

func (c *Consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(broker.EventOrderCreated, true, false, false, false, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		logrus.Warnf("fail to consume: queue=%s, err=%v", q.Name, err)
	}

	forever := make(chan struct{})
	go func() {
		for msg := range msgs {
			c.handleMessage(msg)
		}
	}()

	<-forever
}

func (c *Consumer) handleMessage(msg amqp.Delivery) {
	o := &orderpb.Order{}
	if err := json.Unmarshal(msg.Body, o); err != nil {
		logrus.Infof("fail to unmarshal order: %s", err)
		_ = msg.Nack(false, false)
		return
	}
	_, err := c.app.Commands.CreatePayment.Handle(context.Background(), command.CreatePayment{
		Order: o,
	})
	if err != nil {
		// TODO: retry
		logrus.Infof("fail to create payment: %s", err)
		_ = msg.Nack(false, false)
		return
	}
	_ = msg.Ack(false)
}
