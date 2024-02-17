package external

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func FetchBodyFromGPT(category string, difficulity string, maxRounds int) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error env")
	}

	apiKey := os.Getenv("GPT_KEY")
	maxQuestions := strconv.Itoa(maxRounds)
	client := resty.New()
	language := "polish"
	options := "category: " + category + ", diffuculty:" + difficulity + ", content language: " + language
	prompt := "return json for quiz game with" + maxQuestions + "questions with fields:{ questions, 4x answers, correct answer(index)} " + options

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
		return "", err
	}

	body := response.Body()
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	log.Println(content)
	return content, nil
}
