package services

import (
	"GinniBackend/models"
	"context"

	"github.com/sashabaranov/go-openai"
)

const (
	// number of tokens
	MAX_TOKEN = 5000
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

// NOTE: Implement HandleLongTranscription
func (t *TranslateService) HandleLongTranscription(transcripts []models.Transcription) ([]models.Transcription, error) {
	var result []models.Transcription

	for _, batch := range transcripts {
		batchChars := len(batch.Sentence)
		// check of the batch size is greater than the max token
		if batchChars <= MAX_TOKEN {
			translatedBatch, err := t.TranslateToEnglish(batch.Sentence)
			if err != nil {
				return nil, err
			}
			batch.Sentence = translatedBatch
			result = append(result, batch)
			continue
		} else {
			chunks := batchChars / MAX_TOKEN
			for i := 0; i < chunks; i++ {
				translatedBatch, err := t.TranslateToEnglish(batch.Sentence[i*MAX_TOKEN : (i+1)*MAX_TOKEN])
				if err != nil {
					return nil, err
				}
				batch.Sentence = translatedBatch
				result = append(result, batch)
			}
		}
	}
	return result, nil
}
