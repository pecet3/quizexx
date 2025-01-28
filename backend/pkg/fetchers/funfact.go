package fetchers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pecet3/quizex/pkg/logger"
)

type FetcherFunFact struct {
}

func (f FetcherFunFact) Fetch(ctx context.Context, i interface{}) (string, error) {
	s, ok := i.(string)
	if !ok {
		return "", errors.New("wrong prompt interface")

	}
	prompt := `return a fun fact about this topic in maximum four sentences. Topic is: ` + s
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
	return content, nil
}
