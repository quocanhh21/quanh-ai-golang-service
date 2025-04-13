package service

import (
    "context"
    "os"

    "github.com/sashabaranov/go-openai"
    "github.com/yourname/quanh-ai-golang-service/internal/model"
)

func AnalyzeWithOpenAI(req model.AnalyzeRequest) (*model.AnalyzeResponse, error) {
    client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

    prompt := "Evaluate the following essay and give feedback:\n\n" + req.Text

    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT4,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: prompt,
                },
            },
        },
    )

    if err != nil {
        return nil, err
    }

    return &model.AnalyzeResponse{
        Feedback: resp.Choices[0].Message.Content,
    }, nil
}