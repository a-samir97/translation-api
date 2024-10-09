package tests

import (
	"GinniBackend/models"
	"GinniBackend/services"
	"context"
	"testing"

	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOpenAI struct {
	mock.Mock
}

func (m *MockOpenAI) CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(openai.ChatCompletionResponse), args.Error(1)
}

func TestIsArabic(t *testing.T) {
	// Test case for Arabic text
	service := services.NewTranslateService(nil)
	arabicText := "مرحبا"
	assert.True(t, service.IsArabic((arabicText)))

	// Test case for English text
	englishText := "Hello"
	assert.False(t, service.IsArabic((englishText)))
}

func TestTranslateToEnglish(t *testing.T) {
	mockOpenAI := new(MockOpenAI)
	tests := []struct {
		input        string
		expected     string
		mockResponse string
	}{
		{
			input:        "ماذا تفعل؟",
			expected:     "What are you doing?",
			mockResponse: "What are you doing?",
		},
		{
			input:        "مرحبا",
			expected:     "Hello",
			mockResponse: "Hello",
		},
	}
	service := services.NewTranslateService(mockOpenAI)
	for _, test := range tests {
		mockOpenAI.On("CreateChatCompletion", mock.Anything, mock.Anything).Return(openai.ChatCompletionResponse{
			Choices: []openai.ChatCompletionChoice{
				{
					Message: openai.ChatCompletionMessage{
						Content: test.mockResponse,
					},
				},
			},
		}, nil).Once()
		result, err := service.TranslateToEnglish(test.input)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestTranslateTranscript(t *testing.T) {
	mockOpenAI := new(MockOpenAI)
	tests := []struct {
		input        []models.Transcription
		expected     []models.Transcription
		mockResponse []models.Transcription
	}{
		{
			input:        []models.Transcription{{Speaker: "A", Time: "00:00:00", Sentence: "ماذا تفعل؟"}},
			expected:     []models.Transcription{{Speaker: "A", Time: "00:00:00", Sentence: "What are you doing?"}},
			mockResponse: []models.Transcription{{Speaker: "A", Time: "00:00:00", Sentence: "What are you doing?"}},
		},
	}
	service := services.NewTranslateService(mockOpenAI)
	for _, test := range tests {
		mockOpenAI.On("CreateChatCompletion", mock.Anything, mock.Anything).Return(openai.ChatCompletionResponse{
			Choices: []openai.ChatCompletionChoice{
				{
					Message: openai.ChatCompletionMessage{
						Content: test.mockResponse[0].Sentence,
					},
				},
			},
		}, nil).Once()
		result, err := service.TranslateTranscript(test.input)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}
