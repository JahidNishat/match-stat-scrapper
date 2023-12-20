package repos

import (
	"github.com/spf13/viper"
	"log"
	"match_statistics_scrapper/db"
	"match_statistics_scrapper/models"
	"testing"
)

func TestPlayersStore_InsertData(t *testing.T) {
	viper.SetConfigFile("../../config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	tm := &PlayersStore{DB: db.ConnectDB()}
	players := []models.PlayersData{
		{
			PlayerName: "Max Mackinnon",
			Source:     "ESPN",
			Url:        "https://www.espn.co.uk/mens-college-basketball/player/_/id/5105785/max-mackinnon",
		},
		{
			PlayerName: "Jordan Derkack",
			Source:     "ESPN",
			Url:        "https://www.espn.com/mens-college-basketball/player/_/id/5106600/jordan-derkack",
		},
		{
			PlayerName: "Sean Bairstow",
			Source:     "ESPN",
			Url:        "https://www.espn.co.uk/mens-college-basketball/player/_/id/4593102/sean-bairstow",
		},
		{
			PlayerName: "Travis Roberts",
			Source:     "ESPN",
			Url:        "https://www.espn.co.uk/mens-college-basketball/player/_/id/4684734/travis-roberts",
		},
		{
			PlayerName: "Alex Schumacher",
			Source:     "ESPN",
			Url:        "https://www.espn.co.uk/mens-college-basketball/player/_/id/4712280/alex-schumacher",
		},
		{
			PlayerName: "Jacob Lampkin",
			Source:     "eurobasket",
			Url:        "https://basketball.eurobasket.com/player/Jacob-Lampkin/335999",
		},
		{
			PlayerName: "Jesse Ghee",
			Source:     "eurobasket",
			Url:        "https://www.eurobasket.com/player/Jesse-Ghee/365098",
		},
		{
			PlayerName: "Koen Sapwell",
			Source:     "eurobasket",
			Url:        "https://basketball.eurobasket.com/player/Koen-Sapwell/393132",
		},
		{
			PlayerName: "Jamell Anderson",
			Source:     "eurobasket",
			Url:        "https://basketball.eurobasket.com/player/Jamell-Anderson/179475",
		},
		{
			PlayerName: "Lachlan Anderson",
			Source:     "nbl",
			Url:        "https://nbl.com.au/player/3713/853140/lachlan-anderson",
		},
		{
			PlayerName: "Brock Motum",
			Source:     "bleague",
			Url:        "https://www.bleague.jp/roster_detail/?PlayerID=51000158",
		},
		{
			PlayerName: "AJ Edu",
			Source:     "bleague",
			Url:        "https://www.bleague.jp/roster_detail/?PlayerID=51000301",
		},
		{
			PlayerName: "Jamell Anderson",
			Source:     "britishbasketballleague",
			Url:        "https://www.britishbasketballleague.com/competitions/?WHurl=%2Fperson%2F6661%2Fgamelog%3F",
		},
		{
			PlayerName: "James Moors",
			Source:     "bnxt",
			Url:        "https://bnxtleague.com/en/player-statistics/?player_id=2882&amp;team\\_id=162",
		},
		//Invalid Url
		//{
		//	PlayerName: "***",
		//	Source:     "asiabasket",
		//	Url:        "https://basketball.asia-basket.com/player/AJ-Edu/",
		//},
		//{
		//	PlayerName: "***",
		//	Source:     "***",
		//	Url:        "http://210.140.77.209/player/?key=93&amp;team=715&amp;player=43239",
		//},
	}

	for _, player := range players {
		err := tm.InsertData(&player)
		if err != nil {
			continue
		}
	}
}
