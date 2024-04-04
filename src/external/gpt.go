package external

import (
	"database/sql"
	"log"
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
	log.Println(lang)
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	fmt.Println("error env")
	// }

	// apiKey := os.Getenv("GPT_KEY")

	// client := resty.New()
	// options := "Options for this quiz: category: " + category + ", diffuculty:" + difficulty + ", content language: " + lang
	// prompt := "return json for quiz game with " + maxRounds + " questions." + options + " You have to return correct struct. This is just array of objects. Nothing more, start struct: [{ question, 4x answers, correctAnswer(index)}] "

	// response, err := client.R().
	// 	SetAuthToken(apiKey).
	// 	SetHeader("Content-Type", "application/json").
	// 	SetBody(map[string]interface{}{
	// 		"model":      "gpt-3.5-turbo",
	// 		"messages":   []interface{}{map[string]interface{}{"role": "system", "content": prompt}},
	// 		"max_tokens": 1200,
	// 	}).
	// 	Post(apiEndpoint)

	// if err != nil {
	// 	fmt.Println("connecting with api error")
	// 	return "", err
	// }

	// body := response.Body()
	// var data map[string]interface{}

	// err = json.Unmarshal(body, &data)
	// if err != nil {
	// 	return "", err
	// }
	// content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	content := `[ { "question": "Która marka samochodu pochodzi z Włoch?", "answers": ["Volkswagen", "Toyota", "Fiat", "Ford"], "correctAnswer": 2 }, { "question": "Jak nazywa się popularny model auta marki Mercedes-Benz?", "answers": ["Astra", "Passat", "Clio", "Klasa E"], "correctAnswer": 3 }, { "question": "Co oznacza skrót 'SUV' w motoryzacji?", "answers": ["Super Ultra Vitesse", "Sport Utility Vehicle", "Special Upgrade Version", "Society of United Vehicles"], "correctAnswer": 1 }, { "question": "Która marka samochodu pochodzi z Japonii?", "answers": ["BMW", "Honda", "Audi", "Chevrolet"], "correctAnswer": 1 } ]`

	return content, nil
}
