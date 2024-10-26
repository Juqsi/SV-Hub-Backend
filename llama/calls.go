package llama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struktur für die Antwort von der LLaMA-API
type ResponsePayload struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int64  `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int64  `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
}

// doRequest sendet eine Anfrage an die LLaMA-API und gibt die Antwort zurück
func DoRequest(prompt string) (string, error) {
	// Die URL zur LLaMA-API
	url := "http://localhost:11434/api/generate"

	// Anfrage-Payload erstellen
	payload := RequestPayload{
		Model:  "llama3",
		Prompt: prompt,
		Stream: false,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Erstellen der Anfrage: %v", err)
	}

	// HTTP-Request erstellen
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("Fehler beim Erstellen der HTTP-Anfrage: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Anfrage senden
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Senden der Anfrage: %v", err)
	}
	defer resp.Body.Close()

	// Antwort einlesen
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Lesen der Antwort: %v", err)
	}

	// Antwort dekodieren
	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Dekodieren der Antwort: %v", err)
	}
	fmt.Println(responsePayload)
	return responsePayload.Response, nil
}
