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

	// content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	content := `[ { "question": "Jakie są najpopularniejsze dystrybucje Linuxa?", "answers": ["Ubuntu", "Fedora", "Debian", "Mint"], "correctAnswer": 0 }, { "question": "Która dystrybucja Linuxa jest sponsorowana głównie przez firmę Red Hat?", "answers": ["Ubuntu", "Fedora", "Debian", "Mint"], "correctAnswer": 1 }, { "question": "Która dystrybucja Linuxa jest znana z dbałości o stabilność i system pakietów .deb?", "answers": ["Ubuntu", "Fedora", "Debian", "Mint"], "correctAnswer": 2 }, { "question": "Która dystrybucja Linuxa często jest polecana dla początkujących użytkowników?", "answers": ["Ubuntu", "Fedora", "Debian", "Mint"], "correctAnswer": 0 }, { "question": "Która dystrybucja Linuxa znana jest z lekkości i wyglądu przypominającego Windowsa?", "answers": ["Ubuntu", "Fedora", "Debian", "Mint"], "correctAnswer": 3 } ]`
	log.Println(content)
	return content, nil
}
