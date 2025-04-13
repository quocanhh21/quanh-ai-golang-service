package model

type AnalyzeRequest struct {
    Text string `json:"text"`
}

type AnalyzeResponse struct {
    Feedback string `json:"feedback"`
}