package chatgptAPI

import (
	b64 "encoding/base64"
	"errors"
	"github.com/go-resty/resty/v2"
	"os"
)

func Init() (*OpenAIClient, error) {
	var openAIClient OpenAIClient
	openAIClient.Client = resty.New()

	apiKey := os.Getenv("CHATGPT_API_KEY")
	if apiKey == "" {
		return nil, errors.New("API key not found")
	}

	b64encodedAPIKey := b64.StdEncoding.EncodeToString([]byte(apiKey))

	openAIClient.Client.
		SetBaseURL("http://localhost:8080").
		SetHeaders(map[string]string{
			"Content-Type": "multipart/form-data",
			"User-Agent":   "chatgpt-client-v0.0.1",
			"APIKey":       b64encodedAPIKey,
		})

	return &openAIClient, nil
}

func (openAIClient *OpenAIClient) AskGPT(question string) (*string, error) {
	response, err := openAIClient.Client.R().
		SetMultipartFormData(map[string]string{"question": question}).
		Post(API_ENDPOINT)
	if err != nil {
		return nil, err
	}

	answer := string(response.Body())

	return &answer, nil
}
