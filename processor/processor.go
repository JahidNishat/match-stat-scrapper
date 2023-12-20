package processor

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"match_statistics_scrapper/db"
	"match_statistics_scrapper/db/repos"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/notifier"
	"match_statistics_scrapper/scrapper"
	"match_statistics_scrapper/utils"
	"strings"
)

func FetchPlayerStats() {
	ps := repos.PlayersStore{DB: db.ConnectDB()}
	ss := repos.StatsStore{DB: db.ConnectDB()}

	log.Println("Getting player urls")
	playersData, err := ps.GetPlayersData()
	if err != nil {
		log.Println("Failed to get players data")
		return
	}

	for _, player := range playersData {
		log.Println("Scrapping for ", player.PlayerName)
		url := player.Url
		dataScrapped := scrapFromUrl(url)

		for _, data := range dataScrapped {
			data.PlayerName = player.PlayerName
			data.Source = url

			stringified := utils.StringifyStruct(*data)
			hashed := utils.GetHashOfData(stringified)

			statsData, err := ss.FindHash(hashed)
			if errors.Is(err, gorm.ErrRecordNotFound) && statsData == nil {
				err = notifier.PublishToSubscribers(stringified)
				if err != nil {
					fmt.Println("Failed while sending the update")
					continue
				}
				// updating hash data
				err = ss.InsertData(&models.ScrappedData{
					Hash: hashed,
					Data: stringified,
				})
				if err != nil {
					return
				}
				log.Printf("\n\nSuccessfully published stats of\nPlayer Name : %v,\nSource : %v,\nDate: %v\nData : %v\n\n", data.PlayerName, url, data.Date, stringified)
			}
		}
	}
}

func scrapFromUrl(url string) []*models.MatchStatResponse {
	if strings.Contains(url, "basketball.eurobasket") {
		return scrapper.ScrapsBasketBallEuroBasket(url)
	} else if strings.Contains(url, "eurobasket.com") {
		return scrapper.ScrapsEuroBasket(url)
	} else if strings.Contains(url, "espn.co") {
		return scrapper.EspnScrap(url)
	} else if strings.Contains(url, "nbl.com.au") {
		return scrapper.NblScrap(url)
	} else if strings.Contains(url, "bleague.jp") {
		return scrapper.ScrapsBleague(url)
	} else if strings.Contains(url, "bnxtleague.com") {
		return scrapper.ScrapsBnxt(url)
	} else if strings.Contains(url, "britishbasketballleague.com") {
		return scrapper.ScrapsBritishBasketBall(url)
	} else if strings.Contains(url, "210.140.77.209") {
		return scrapper.ScrapsB3league(url)
	}

	fmt.Println("URL not supported : ", url)
	return nil
}
