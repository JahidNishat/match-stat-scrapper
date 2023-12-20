package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/utils"
)

func ScrapsBasketBallEuroBasket(url string) []*models.MatchStatResponse {
	// creating a new Colly instance
	c := colly.NewCollector()

	var rows [][]string

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	// scraping logic section.gamelogWidget
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("#23", func(_ int, row *colly.HTMLElement) {
			var cells []string

			row.ForEach("td", func(_ int, cell *colly.HTMLElement) {
				cells = append(cells, cell.Text)
			})

			rows = append(rows, cells)
		})
	})

	// visiting the target page
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Couldn't Visit")
	}

	stats := []*models.MatchStatResponse{}
	for _, row := range rows {
		statResp := &models.MatchStatResponse{
			Date:   utils.EuroBasketDate(row[0]),
			Opp:    row[2],
			Result: row[3],
			Min:    row[4],
			FGP:    row[6],
			FTP:    row[8],
			ThreeP: row[7],
			REB:    row[11],
			AST:    row[12],
			BLK:    row[14],
			STL:    row[15],
			PF:     row[13],
			TO:     row[16],
			PTS:    row[5],
		}
		stats = append(stats, statResp)
	}
	fmt.Println("EuroBasketStat Scrapper result : ", stats)
	return stats
}
