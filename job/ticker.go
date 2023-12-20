package job

import (
	"log"
	"match_statistics_scrapper/processor"
	"time"
)

func StartTicker() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("######################### Started Fetching #########################")
			processor.FetchPlayerStats()
			log.Println("######################### Finished Fetching #########################")
		}
	}
}
