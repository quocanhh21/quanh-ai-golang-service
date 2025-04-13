package model

type AIAnalyzeResponse struct {
	Feedback  string   `json:"feedback"`
	CEFRLevel string   `json:"cefr_level"`
	Roadmap   []string `json:"roadmap"`
}
