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
	player := models.PlayersData{
		PlayerName: "Max Mackinnon",
		Source:     "ESPN",
		Url:        "https://www.espn.co.uk/mens-college-basketball/player/_/id/5105785/max-mackinnon",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Jordan Derkack",
		Source:     "ESPN",
		Url:        "https://www.espn.com/mens-college-basketball/player/_/id/5106600/jordan-derkack",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Jacob Lampkin",
		Source:     "eurobasket",
		Url:        "https://basketball.eurobasket.com/player/Jacob-Lampkin/335999",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Jesse Ghee",
		Source:     "eurobasket",
		Url:        "https://www.eurobasket.com/player/Jesse-Ghee/365098",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Sean Bairstow",
		Source:     "ESPN",
		Url:        "https://www.espn.co.uk/mens-college-basketball/player/_/id/4593102/sean-bairstow",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Travis Roberts",
		Source:     "ESPN",
		Url:        "https://www.espn.co.uk/mens-college-basketball/player/_/id/4684734/travis-roberts",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Lachlan Anderson",
		Source:     "nbl",
		Url:        "https://nbl.com.au/player/3713/853140/lachlan-anderson",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Brock Motum",
		Source:     "bleague",
		Url:        "https://www.bleague.jp/roster_detail/?PlayerID=51000158",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Jamell Anderson",
		Source:     "eurobasket",
		Url:        "https://basketball.eurobasket.com/player/Jamell-Anderson/179475",
	}
	tm.InsertData(&player)

	player = models.PlayersData{
		PlayerName: "Jamell Anderson",
		Source:     "britishbasketballleague",
		Url:        "https://www.britishbasketballleague.com/competitions/?WHurl=%2Fperson%2F6661%2Fgamelog%3F",
	}
	tm.InsertData(&player)
}
