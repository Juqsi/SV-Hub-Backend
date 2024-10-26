package weaviate

import (
	"HexMaster/utils"
	"context"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"log"
)

func InsertData(texts []string) {
	cfg := weaviate.Config{
		Host:   utils.GetEnv("WEAVIATE_HOST", "localhost:8080"),
		Scheme: utils.GetEnv("WEAVIATE_SCHEME", "http"),
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error initializing Weaviate client: %v", err)
	}

	schema := &models.Class{
		Class: "Document",
		Properties: []*models.Property{
			{
				Name:     "text",
				DataType: []string{"string"},
			},
		},
	}

	_, err = client.Schema().ClassGetter().WithClassName("Document").Do(context.Background())
	if err != nil {
		err = client.Schema().ClassCreator().WithClass(schema).Do(context.Background())
		if err != nil {
			log.Fatalf("Error creating schema: %v", err)
		}
	}

	for _, text := range texts {
		_, err := client.Data().Creator().
			WithClassName("Document").
			WithID("uuid").
			WithProperties(map[string]interface{}{
				"text": text,
			}).
			Do(context.Background())
		if err != nil {
			log.Printf("Error inserting into Weaviate: %v", err)
			continue
		}
	}

	log.Println("Texts successfully inserted and vectorized.")
}
