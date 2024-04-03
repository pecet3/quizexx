package external

import (
	"database/sql"
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

type ExternalService struct {
}
type IExternal interface {
	FetchQuestionSet(category, maxRounds, difficulty, lang string) (string, error)
	NewExternalService() *ExternalService
	SaveQuestionSetToDB(db *sql.DB)
}

func (e *ExternalService) NewExternalService() *ExternalService {
	return &ExternalService{}

}

func (e *ExternalService) SaveQuestionSetToDB(db *sql.DB) {

}

func (e *ExternalService) FetchQuestionSet(category, maxRounds, difficulty, lang string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error env")
	}

	apiKey := os.Getenv("GPT_KEY")

	client := resty.New()
	options := "Options for this quiz: category: " + category + ", diffuculty:" + difficulty + ", content language: " + lang
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
		log.Println("error kurwa")
	}
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	log.Println(data)
	// var questions []RoundQuestion

	// err = json.Unmarshal([]byte(content), &questions)
	// if err != nil {
	// 	log.Println("error with unmarshal data")
	// 	log.Println(err)
	// }

	// maxRoundsInt, err := strconv.Atoi(maxRounds)
	// if err != nil {
	// 	return nil, err
	// }
	// if len(questions) != maxRoundsInt {
	// 	log.Println("ChatGPT returned insufficient content. Trying to process again...")
	// 	return e.FetchQuestionSet(category, maxRounds, difficulty, lang)
	// }
	// log.Println(questions)
	return content, nil
}
