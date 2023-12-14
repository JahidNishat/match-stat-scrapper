package main

import (
	"github.com/spf13/viper"
	"log"
	"match_statistics_scrapper/job"
)

func main() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	//
	job.StartTicker()
	//notifier.TgBot()
	//scrapper.ScrapsBnxt(`https://bnxtleague.com/en/player-statistics/?player_id=2882&amp;team\_id=162`)
	//scrapper.EspnScrap("https://www.espn.co.uk/mens-college-basketball/player/_/id/5105785/max-mackinnon")
	//scrapper.ScrapsEuroBasket("https://basketball.eurobasket.com/player/Jacob-Lampkin/335999")
	//scrapper.NblScrap("https://nbl.com.au/player/3713/853140/lachlan-anderson")
	//scrapper.ScrapsBleague("https://www.bleague.jp/roster_detail/?PlayerID=51000158")
	//scrapper.ScrapsBritishBasketBall("https://www.britishbasketballleague.com/competitions/?WHurl=%2Fperson%2F6661%2Fgamelog%3F")
	//scrapper.ScrapsB3league("http://210.140.77.209/player/?key=93&amp;team=715&amp;player=43239") // Working one : https://www.b3league.jp/player/?key=93&team=2725&player=9208
}
