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

var urlScrapFuncs = map[string]func(string) []*models.MatchStatResponse{
	//"basketball.eurobasket":       scrapper.ScrapsBasketBallEuroBasket,
	"eurobasket.com":              scrapper.ScrapsEuroBasket,
	"espn.co":                     scrapper.EspnScrap,
	"nbl.com.au":                  scrapper.NblScrap,
	"bleague.jp":                  scrapper.ScrapsBleague,
	"bnxtleague.com":              scrapper.ScrapsBnxt,
	"britishbasketballleague.com": scrapper.ScrapsBritishBasketBall,
	"210.140.77.209":              scrapper.ScrapsB3league,
}

func scrapFromUrl(url string) []*models.MatchStatResponse {
	for key, scraperFunc := range urlScrapFuncs {
		if strings.Contains(url, key) {
			return scraperFunc(url)
		}
	}

	fmt.Println("URL not supported:", url)
	return nil
}
