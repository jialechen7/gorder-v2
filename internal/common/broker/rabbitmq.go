package broker

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

func Connect(user, password, host, port string) (*amqp.Channel, func() error) {
	address := fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port)
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, func() error { return nil }
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, func() error { return nil }
	}
	err = ch.ExchangeDeclare(EventOrderCreated, "direct", true, false, false, false, nil)
	if err != nil {
		return nil, func() error { return nil }
	}

	err = ch.ExchangeDeclare(EventOrderPaid, "fanout", true, false, false, false, nil)
	if err != nil {
		return nil, func() error { return nil }
	}
	return ch, ch.Close
}

type RabbitMQHeaderCarrier map[string]interface{}

func (r RabbitMQHeaderCarrier) Get(key string) string {
	value, ok := r[key]
	if !ok {
		return ""
	}
	return value.(string)
}

func (r RabbitMQHeaderCarrier) Set(key string, value string) {
	r[key] = value
}

func (r RabbitMQHeaderCarrier) Keys() []string {
	keys := make([]string, 0)
	for key := range r {
		keys = append(keys, key)
	}
	return keys
}

func InjectRabbitMQHeaders(ctx context.Context) map[string]interface{} {
	carrier := make(RabbitMQHeaderCarrier)
	otel.GetTextMapPropagator().Inject(ctx, &carrier)
	return carrier
}

func ExtractRabbitMQHeaders(ctx context.Context, headers map[string]interface{}) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, RabbitMQHeaderCarrier(headers))
}
