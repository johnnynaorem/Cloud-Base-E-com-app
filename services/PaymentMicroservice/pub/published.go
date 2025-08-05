package pub

import (
	"context"
	"encoding/json"
	"log"
	"paymentMicroservice/models"

	"cloud.google.com/go/pubsub"
)

const OrderTopic = "payment-events"

func PublishOrderCreated(ctx context.Context, client *pubsub.Client, order models.PaymentEvent) {
	data, _ := json.Marshal(order)
	topic := client.Topic(OrderTopic)
	result := topic.Publish(ctx, &pubsub.Message{Data: data})
	_, err := result.Get(ctx)
	if err != nil {
		log.Printf("Error publishing order: %v", err)
	}
}
