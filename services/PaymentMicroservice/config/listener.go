package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"paymentMicroservice/models"
	"paymentMicroservice/repository"

	"cloud.google.com/go/pubsub"
)

func SubscribeToOrderEvents(repo *repository.PaymentRepoImpl) {
	fmt.Println("📡 Subscribing to 'payment-sub'...")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "johnny-projectt")
	if err != nil {
		log.Fatalf("PubSub Client error: %v", err)
	}

	sub := client.Subscription("payment-sub")

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("✅ Payment Service received: %s\n", string(msg.Data))

		var event models.OrderCreatedEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("❌ Failed to parse message: %v", err)
			msg.Ack()
			return
		}

		newPayment := &models.PaymentCreate{
			OrderID: event.OrderID,
			Amount:  event.TotalAmount,
			Method:  models.DebitCard,
		}

		repo.MakePayment(client, ctx, newPayment)
		msg.Ack()
	})

	if err != nil {
		log.Fatalf("❌ Failed to receive messages: %v", err)
	}
}
