package main

import (
	"sync"
	"time"
)

type EmailStats struct {
	TotalSent     int               `json:"total_sent"`
	EmailsPerUser map[string]int    `json:"emails_per_user"`
	DailyLimit    int               `json:"daily_limit"`
	LastSentTimes map[string]string `json:"last_sent_times"`
}

var (
	emailStats = EmailStats{
		TotalSent:     0,
		EmailsPerUser: make(map[string]int),
		DailyLimit:    10,
		LastSentTimes: make(map[string]string),
	}
	mutex sync.Mutex
)

func incrementEmailCount(sender string) {
	mutex.Lock()
	defer mutex.Unlock()

	emailStats.TotalSent++
	emailStats.EmailsPerUser[sender]++
	emailStats.LastSentTimes[sender] = time.Now().Format(time.RFC3339)
}

func getEmailStats() EmailStats {
	mutex.Lock()
	defer mutex.Unlock()
	return emailStats
}
