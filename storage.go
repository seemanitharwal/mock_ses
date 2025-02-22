package main

import (
	"sync"
	"time"
)

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
