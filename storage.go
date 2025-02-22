package main

var emailStats = EmailStats{TotalSent: 0}

func incrementEmailCount() {
	emailStats.TotalSent++
}

func getEmailStats() EmailStats {
	return emailStats
}
