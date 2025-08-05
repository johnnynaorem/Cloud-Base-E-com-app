package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"orderPaymentMicroservice/models"

	"cloud.google.com/go/pubsub"
)

const OrderTopic = "order-events"

func PublishOrderCreated(ctx context.Context, client *pubsub.Client, order models.OrderEvent) {
	data, _ := json.Marshal(order)
	topic := client.Topic(OrderTopic)
	fmt.Println(order)
	result := topic.Publish(ctx, &pubsub.Message{Data: data})
	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("Publish error: %v", err)
		return
	}
	fmt.Printf("✅ Order created and published! Message ID: %s", id)
}
