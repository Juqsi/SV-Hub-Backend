package weaviate

import (
	"HexMaster/utils"
	"context"
	"fmt"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"log"
)

func DeleteSchema() {
	cfg := weaviate.Config{
		Host:   utils.GetEnv("WEAVIATE_HOST", "localhost:8080"),
		Scheme: utils.GetEnv("WEAVIATE_SCHEME", "http"),
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error initializing Weaviate client: %v", err)
	}

	err = client.Schema().AllDeleter().Do(context.Background())
	if err != nil {
		log.Fatalf("Error deleting schema: %v", err)
	}

	fmt.Println("All data and schema deleted successfully.")
}
