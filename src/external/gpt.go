package external

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

type Response struct {
	Category   string     `json:"category"`
	Difficulty string     `json:"difficulty"`
	Language   string     `json:"language"`
	Questions  []Question `json:"questions"`
}
type Question struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

func FetchBodyFromGPT(category string, difficulity string) Response {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error env")
	}

	apiKey := os.Getenv("GPT_KEY")

	client := resty.New()
	language := "polish"
	options := "category: " + category + ", diffuculty:" + difficulity + ", content language: " + language
	prompt := "return json for quiz game with 5 questions with fields:{ questions, 4x answers, correct answer(index)} " + options

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": prompt}},
			"max_tokens": 400,
		}).
		Post(apiEndpoint)

	if err != nil {
		fmt.Println("connecting with api error")
	}

	body := response.Body()
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		return Response{}
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	var parsedContent Response
	err = json.Unmarshal([]byte(content), &parsedContent)
	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		return Response{}
	}
	return parsedContent
}
