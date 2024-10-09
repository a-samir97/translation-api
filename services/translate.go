package services

import (
	"GinniBackend/models"
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClient interface {
	CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error)
}

type TranslateService struct {
	Client OpenAIClient
}

func NewTranslateService(client OpenAIClient) *TranslateService {
	return &TranslateService{
		Client: client,
	}
}

func (t *TranslateService) TranslateTranscript(transcriptions []models.Transcription) ([]models.Transcription, error) {
	var translatedTranscriptions []models.Transcription

	for _, transcription := range transcriptions {
		translatedText, err := t.TranslateToEnglish(transcription.Sentence)
		if err != nil {
			return nil, err
		}
		transcription.Sentence = translatedText
		translatedTranscriptions = append(translatedTranscriptions, transcription)
	}
	return translatedTranscriptions, nil
}

func (t *TranslateService) TranslateToEnglish(text string) (string, error) {
	resp, err := t.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Translate the following text to English",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		})

	if err != nil {
		return "", err
	}
	fmt.Println(resp.Choices[0].Message.Content, text)
	return resp.Choices[0].Message.Content, nil
}

func (t *TranslateService) IsArabic(text string) bool {
	for _, r := range text {
		if r >= '\u0600' && r <= '\u06FF' {
			return true
		}
	}
	return false
}
