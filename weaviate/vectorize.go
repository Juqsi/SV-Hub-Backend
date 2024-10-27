package weaviate

import (
	"HexMaster/utils"
	"context"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	models "github.com/weaviate/weaviate/entities/models"
	"log"
)

func VectorizeAndStoreData(texts []string) {
	cfg := weaviate.Config{
		Host:   utils.GetEnv("WEAVIATE_HOST", "localhost:8080"),
		Scheme: utils.GetEnv("WEAVIATE_SCHEME", "http"),
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatalf("Fehler beim Initialisieren des Weaviate-Clients: %v", err)
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

	err = client.Schema().ClassCreator().WithClass(schema).Do(context.Background())
	if err != nil {
		log.Fatalf("Fehler beim Erstellen des Schemas: %v", err)
	}

	for _, text := range texts {
		_, err := client.Data().Creator().
			WithClassName("Document").
			WithProperties(map[string]interface{}{
				"text": text,
			}).
			Do(context.Background())
		if err != nil {
			log.Printf("Fehler beim Einfügen in Weaviate: %v", err)
			continue
		}
	}

	log.Println("Texte erfolgreich eingefügt und vektorisiert.")
}
