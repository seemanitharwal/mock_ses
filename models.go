package main

type EmailRequest struct {
	To      string `json:"to" binding:"required"`
	From    string `json:"from" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

type EmailStats struct {
	TotalSent int `json:"total_sent"`
}
