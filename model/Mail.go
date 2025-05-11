package model

type Mail struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	From      string `json:"from"`
	To        string `json:"to"`
	CreatedAt string `json:"created_at"`
}
