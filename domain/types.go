package domain

type CardGenerateRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type OpenAIRequest struct {
	Model          string                 `json:"model"`
	ResponseFormat OpenAIResponseFormat   `json:"response_format"`
	Messages       []OpenAIRequestMessage `json:"messages"`
}

type OpenAIRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponseFormat struct {
	Type string `json:"type"`
}

type OpenAIResponse struct {
	ID      string                 `json:"id"`
	Choices []OpenAIResponseChoice `json:"choices"`
}

type OpenAIResponseChoice struct {
	Message OpenAIResponseChoiceMessage `json:"message"`
}

type OpenAIResponseChoiceMessage struct {
	Content string `json:"content"`
}
