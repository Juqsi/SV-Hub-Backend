package weaviate

import (
	"context"
	"fmt"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"log"
)

func DeleteSchema() {
	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error initializing Weaviate client: %v", err)
	}

	// Delete the entire schema
	err = client.Schema().AllDeleter().Do(context.Background())
	if err != nil {
		log.Fatalf("Error deleting schema: %v", err)
	}

	fmt.Println("All data and schema deleted successfully.")
}
