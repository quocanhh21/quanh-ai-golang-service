package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/quocanhh21/quanh-ai-golang-service/internal/model"
	"github.com/sashabaranov/go-openai"
)

func AnalyzeWithOpenAI(text string) (*model.AIAnalyzeResponse, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	prompt := fmt.Sprintf(`
You are an English tutor. A student submitted this text:

"%s"

1. Give grammar and clarity feedback.
2. Estimate CEFR level (A1 to C2).
3. Create a 7-day learning roadmap (title each day based on skills they need most).

Respond ONLY in this JSON format:
{
  "feedback": "...",
  "cefr_level": "...",
  "roadmap": ["Day 1: ...", "Day 2: ..."]
}
`, text)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleUser, Content: prompt},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	var result model.AIAnalyzeResponse
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}

	return &result, nil
}
