package weaviate

import (
	"HexMaster/utils"
	"context"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"log"
)

func InsertData(texts []string, class, id string) {
	cfg := weaviate.Config{
		Host:   utils.GetEnv("WEAVIATE_HOST", "localhost:8080"),
		Scheme: utils.GetEnv("WEAVIATE_SCHEME", "http"),
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error initializing Weaviate client: %v", err)
	}

	schema := &models.Class{
		Class: class,
		Properties: []*models.Property{
			{
				Name:     "text",
				DataType: []string{"string"},
			},
		},
	}

	_, err = client.Schema().ClassGetter().WithClassName(class).Do(context.Background())
	if err != nil {
		err = client.Schema().ClassCreator().WithClass(schema).Do(context.Background())
		if err != nil {
			log.Fatalf("Error creating schema: %v", err)
		}
	}

	for index, text := range texts {
		_, err := client.Data().Creator().
			WithClassName(class).
			WithID(id + string(index)).
			//default tenant for testing later is access (group) related
			WithTenant("default").
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
