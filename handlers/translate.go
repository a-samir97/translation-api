package handlers

import (
	"GinniBackend/models"
	"GinniBackend/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

// TranslateHandler translates the given text to English
func TranslateHandler(c *gin.Context) {

	var transcriptions []models.Transcription

	if err := c.ShouldBindJSON(&transcriptions); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	service := services.NewTranslateService(client)
	response, err := service.TranslateTranscript(transcriptions)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": response,
	})
}
