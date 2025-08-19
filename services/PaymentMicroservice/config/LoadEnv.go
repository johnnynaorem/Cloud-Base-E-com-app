package config

import (
	"context"
	"log"
	"os"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func LoadSecretsFromGCP() {
	ctx := context.Background()

	secretName := os.Getenv("SECRET_CREDENTIALS")
	if secretName == "" {
		log.Fatal("❌ SECRET_NAME environment variable not set")
	}

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{Name: secretName}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("❌ Failed to access secret version: %v", err)
	}

	secretData := string(result.Payload.Data)

	lines := strings.Split(secretData, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}
	log.Println("✅ Secrets loaded into environment variables")
}
