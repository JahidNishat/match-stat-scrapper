package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
)

func ScrapsB3league(url string) []*models.MatchStatResponse {
	// creating a new Colly instance
	c := colly.NewCollector()

	var rows [][]string

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	// Counter to keep track of tbody elements
	tbodyCount := 0
	// scraping logic section.gamelogWidget
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		// Increment the counter
		tbodyCount++

		// Check if it's the third tbody
		if tbodyCount == 3 {
			e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
				var cells []string

				row.ForEach("th", func(_ int, cell *colly.HTMLElement) {
					cells = append(cells, cell.Text)
				})

				rows = append(rows, cells)
			})
		}
	})

	// visiting the target page
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Couldn't Visit")
	}

	stats := []*models.MatchStatResponse{}
	for _, row := range rows {
		statResp := &models.MatchStatResponse{
			Date:   row[2],
			Opp:    "",
			Result: "",
			Min:    row[22],
			FGP:    "",
			FTP:    row[12],
			ThreeP: row[6],
			REB:    row[16],
			AST:    row[18],
			BLK:    row[20],
			STL:    row[19],
			PF:     row[13],
			TO:     row[17],
			PTS:    row[3],
		}
		stats = append(stats, statResp)
	}
	fmt.Println("B3league Scrapper result : ", stats)
	return stats
}
