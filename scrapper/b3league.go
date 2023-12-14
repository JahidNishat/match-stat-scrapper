package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
)

func ScrapsB3league(url string) []models.B3leagueStat {
	// creating a new Colly instance
	c := colly.NewCollector()

	var rows [][]string

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
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

	stats := []models.B3leagueStat{}
	for _, row := range rows {
		data := models.B3leagueStat{
			Date:   row[2],
			Pts:    row[3],
			ThreeP: row[6],
			TwoP:   row[9],
			FtP:    row[12],
			Pf:     row[13],
			Oreb:   row[14],
			Dreb:   row[15],
			Reb:    row[16],
			Tov:    row[17],
			Ast:    row[18],
			Stl:    row[19],
			Blk:    row[20],
			Eff:    row[21],
			Min:    row[22],
		}
		stats = append(stats, data)
	}
	fmt.Println("B3league Scrapper result : ", stats)
	return stats
}
