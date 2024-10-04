package main

import (
	"claydol/clients"
	"claydol/domain"
	"claydol/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if os.Getenv("CLAYDOL_ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	apiKey := os.Getenv("CLAYDOL_OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	r.POST("/generate", func(c *gin.Context) {

		var cardRequest domain.CardGenerateRequest
		if err := c.ShouldBindJSON(&cardRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		prompt := utils.ConstructPrompt(cardRequest.Prompt)

		response, err := clients.CallOpenAI(prompt, apiKey)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to communicate with OpenAI"})
			return
		}

		var jsonMap map[string]interface{}
		err = json.Unmarshal([]byte(response), &jsonMap)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse OpenAI response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": jsonMap})
	})

	r.Run()
}
