package quiz

import (
	"context"
	"log"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func fetchQuestionSet(ctx context.Context, category, maxRounds, difficulty, lang string) (string, error) {
	log.Println("> Fetching Question Set")
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	fmt.Println("error env")
	// }

	// apiKey := os.Getenv("GPT_KEY")
	// log.Println(apiKey[:16])

	// options := "Options for this quiz:" +
	// 	" category: " + category +
	// 	", diffuculty:" + difficulty +
	// 	", content language: " + lang

	// prompt := "return json for quiz game with " +
	// 	maxRounds + " questions." +
	// 	options +
	// 	" You have to return correct struct. This is just array of objects. Nothing more, start struct: [{ question, 4x answers, correctAnswer(index)}] "

	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	// reqBody, err := json.Marshal(map[string]interface{}{
	// 	"model":      "gpt-3.5-turbo",
	// 	"messages":   []interface{}{map[string]interface{}{"role": "system", "content": prompt}},
	// 	"max_tokens": 1200,
	// })
	// if err != nil {
	// 	return "", err
	// }

	// req, err := http.NewRequestWithContext(ctx, "POST", apiEndpoint, bytes.NewBuffer(reqBody))
	// if err != nil {
	// 	return "", err
	// }
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer "+apiKey)

	// client := &http.Client{}
	// response, err := client.Do(req)
	// if err != nil {
	// 	return "", err
	// }
	// defer response.Body.Close()

	// var data map[string]interface{}

	// err = json.NewDecoder(response.Body).Decode(&data)
	// if err != nil {
	// 	return "", err
	// }

	// content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	content := `[ { "question": "Która marka samochodu pochodzi z Włoch?", "answers": ["Volkswagen", "Toyota", "Fiat", "Ford"], "correctAnswer": 2 }, { "question": "Jak nazywa się popularny model auta marki Mercedes-Benz?", "answers": ["Astra", "Passat", "Clio", "Klasa E"], "correctAnswer": 3 }, { "question": "Co oznacza skrót 'SUV' w motoryzacji?", "answers": ["Super Ultra Vitesse", "Sport Utility Vehicle", "Special Upgrade Version", "Society of United Vehicles"], "correctAnswer": 1 }, { "question": "Która marka samochodu pochodzi z Japonii?", "answers": ["BMW", "Honda", "Audi", "Chevrolet"], "correctAnswer": 1 } ]`
	log.Println("> Question Set Fetching Success")

	return content, nil
}
