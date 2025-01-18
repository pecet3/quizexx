package quiz

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pecet3/quizex/pkg/logger"
)

const (
	testContent = `[ { "question": "Która marka samochodu pochodzi z Włoch?", "answers": ["Volkswagen", "Toyota", "Fiat", "Ford"], "correct_answer": 2 }, { "question": "Jak nazywa się popularny model auta marki Mercedes-Benz?", "answers": ["Astra", "Passat", "Clio", "Klasa E"], "correct_answer": 3 }, { "question": "Co oznacza skrót 'SUV' w motoryzacji?", "answers": ["Super Ultra Vitesse", "Sport Utility Vehicle", "Special Upgrade Version", "Society of United Vehicles"], "correct_answer": 1 }, { "question": "Która marka samochodu pochodzi z Japonii?", "answers": ["BMW", "Honda", "Audi", "Chevrolet"], "correct_answer": 1 } ]`

	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func fetchFromGPT(ctx context.Context, prompt string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		logger.Error(err)
		return "", err
	}

	apiKey := os.Getenv("GPT_KEY")

	reqBody, err := json.Marshal(map[string]interface{}{
		"model":       "gpt-4o-mini",
		"messages":    []interface{}{map[string]interface{}{"role": "system", "content": prompt}},
		"max_tokens":  1200,
		"temperature": 0.8,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	content = strings.ReplaceAll(content, "```json", "")
	content = strings.ReplaceAll(content, "```", "")
	return content, nil
}
