package main

type EmailRequest struct {
	Sender    string `json:"sender" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Body      string `json:"body" binding:"required"`
}

type EmailStats struct {
	TotalSent     int               `json:"total_sent"`
	EmailsPerUser map[string]int    `json:"emails_per_user"`
	DailyLimit    int               `json:"daily_limit"`
	LastSentTimes map[string]string `json:"last_sent_times"`
}
