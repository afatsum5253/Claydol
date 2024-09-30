package clients

import (
	"bytes"
	"claydol/domain"
	"encoding/json"
	"io"
	"log"
	"net/http"
	_ "os"
)

func CallOpenAI(prompt, apiKey string) (string, error) {

	const openaiAPIUrl = "https://api.openai.com/v1/chat/completions"

	reqData := domain.OpenAIRequest{
		Model: "gpt-4o-mini",
		ResponseFormat: domain.OpenAIResponseFormat{
			Type: "json_object",
		},
		Messages: []domain.OpenAIRequestMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		panic(err)
	}

	// Make the request to OpenAI
	req, err := http.NewRequest("POST", openaiAPIUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var openAIResponse domain.OpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		log.Fatalln(err)
	}

	return openAIResponse.Choices[0].Message.Content, nil
}
