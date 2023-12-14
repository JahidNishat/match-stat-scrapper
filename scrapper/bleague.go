package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
	"strings"
)

func ScrapsBleague(url string) {
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

	// scraping logic section.gamelogWidget
	c.OnHTML("#scores", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
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

	stats := []models.BleagueStat{}
	for _, row := range rows {
		data := models.BleagueStat{
			Day:       strings.TrimSpace(row[0]),
			VS:        strings.TrimSpace(row[1]),
			HA:        strings.TrimSpace(row[2]),
			WL:        strings.TrimSpace(row[3]),
			Min:       strings.TrimSpace(row[5]),
			Pts:       strings.TrimSpace(row[6]),
			FgP:       strings.TrimSpace(row[9]),
			TwoFgP:    strings.TrimSpace(row[12]),
			ThreeFgP:  strings.TrimSpace(row[15]),
			FtP:       strings.TrimSpace(row[18]),
			EfgP:      strings.TrimSpace(row[19]),
			TsP:       strings.TrimSpace(row[20]),
			Or:        strings.TrimSpace(row[21]),
			Dr:        strings.TrimSpace(row[22]),
			Tr:        strings.TrimSpace(row[23]),
			As:        strings.TrimSpace(row[24]),
			Ast:       strings.TrimSpace(row[25]),
			To:        strings.TrimSpace(row[26]),
			St:        strings.TrimSpace(row[27]),
			Bs:        strings.TrimSpace(row[28]),
			Bsr:       strings.TrimSpace(row[29]),
			F:         strings.TrimSpace(row[30]),
			Fd:        strings.TrimSpace(row[31]),
			Eff:       strings.TrimSpace(row[32]),
			PlusMinus: strings.TrimSpace(row[33]),
		}
		stats = append(stats, data)
	}
	fmt.Println("Bleague Scrapper result : ", stats)
	return
}
