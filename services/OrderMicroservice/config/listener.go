package config

import (
	"context"
	"encoding/json"
	"log"
	"orderPaymentMicroservice/models"
	"orderPaymentMicroservice/repository"

	"cloud.google.com/go/pubsub"
)

func ListenForPayments(repo repository.OrderRepository) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "johnny-projectt")
	if err != nil {
		log.Fatalf("PubSub Client error: %v", err)
	}
	sub := client.Subscription("order-status-sub")

	go func() {
		err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
			var evt models.PaymentEvent
			if err := json.Unmarshal(m.Data, &evt); err != nil {
				log.Printf("Invalid message: %v", err)
				m.Nack()
				return
			}

			log.Printf("Received order event: %+v", evt)

			_, err = repo.UpdateOrderStatus(evt.OrderID, string(evt.Status))
			if err != nil {
				log.Printf("Payment status update failed: %v", err)
				m.Nack()
				return
			}

			m.Ack()
		})

		if err != nil {
			log.Fatalf("PubSub subscription error: %v", err)
		}
	}()
}
