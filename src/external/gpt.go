package external

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func FetchBodyFromGPT(category string, difficulty string, maxRounds string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error env")
	}

	apiKey := os.Getenv("GPT_KEY")

	client := resty.New()
	language := "polish"
	options := "This is Options for this quiz: category: " + category + ", diffuculty:" + difficulty + ", content language: " + language
	prompt := "return json for quiz game with " + maxRounds + " questions." + options + " You have to return correct struct. This is just array of objects. Nothing more, start struct: [{ question, 4x answers, correctAnswer(index)}] "

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": prompt}},
			"max_tokens": 1200,
		}).
		Post(apiEndpoint)

	if err != nil {
		fmt.Println("connecting with api error")
		return "", err
	}

	body := response.Body()
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	log.Println(category, difficulty, maxRounds)
	return content, nil
}
