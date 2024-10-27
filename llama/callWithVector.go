package llama

import (
	"HexMaster/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"io/ioutil"
	"net/http"
)

type RequestPayload struct {
	Model   string        `json:"model"`
	Prompt  string        `json:"prompt"`
	Stream  bool          `json:"stream"`
	Vectors []interface{} `json:"vectors"`
}

// DoRequestWithVectors sends a request to the LLaMA API with vectorized data
func DoRequestWithVectors(relatedTags []string, prompt, class string) (string, error) {
	cfg := weaviate.Config{
		Host:   utils.GetEnv("WEAVIATE_HOST", "localhost:8080"),
		Scheme: utils.GetEnv("WEAVIATE_SCHEME", "http"),
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Initialisieren des Weaviate-Clients: %v", err)
	}

	var allVectors []interface{}
	for _, tag := range relatedTags {
		// Retrieve vectorized data for each tag
		result, err := client.GraphQL().Get().
			WithClassName(class).
			WithFields(graphql.Field{Name: "text"}).WithLimit(5).
			//default tenant for testing later is access (group) related
			WithTenant("default").
			Do(context.Background())
		if err != nil {
			return "", fmt.Errorf("Fehler beim Abrufen der vektorisierten Daten für Tag %s: %v", tag, err)
		}

		documents := result.Data["Get"].(map[string]interface{})[class].([]interface{})
		for _, doc := range documents {
			allVectors = append(allVectors, doc.(map[string]interface{}))
		}
	}

	prompt = fmt.Sprintf("Context: %s \n Gibt mir für die Frage ausschließlich passende Zitate aus dem Context mit den uuids. Frage: %v", prompt, allVectors)

	payload := RequestPayload{
		Model:  "llama3",
		Prompt: prompt,
		Stream: false,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Erstellen der Anfrage: %v", err)
	}
	fmt.Println("Vektoren:", payload)

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("Fehler beim Erstellen der HTTP-Anfrage: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Senden der Anfrage: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Lesen der Antwort: %v", err)
	}

	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Dekodieren der Antwort: %v", err)
	}
	return responsePayload.Response, nil
}
