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
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	apiKey := os.Getenv("CLAYDOL_OPENAI_API_KEY")

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
