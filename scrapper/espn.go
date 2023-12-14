package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
)

func EspnScrap(url string) {
	// creating a new Colly instance
	c := colly.NewCollector()

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
	c.OnHTML("section.gamelogWidget", func(e *colly.HTMLElement) {
		// Collect all matching tbody elements
		allTr := e.DOM.Find("tbody").Find("tr")

		var allStat []models.EspnStat
		allTr.Each(func(i int, element *goquery.Selection) {
			var trData []string
			allTd := element.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})
			if len(trData) == 14 {
				statEspn := models.EspnStat{
					Date:   trData[0],
					Opp:    trData[1],
					Result: trData[2],
					Min:    trData[3],
					FG:     trData[4],
					FT:     trData[5],
					ThreeP: trData[6],
					REB:    trData[7],
					AST:    trData[8],
					BLK:    trData[9],
					STL:    trData[10],
					PF:     trData[11],
					TO:     trData[12],
					PTS:    trData[13],
				}
				allStat = append(allStat, statEspn)
			}
		})
		fmt.Println("Espn Scrapper Result", allStat)
	})

	// visiting the target page
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error while visiting url : ", err)
		return
	}
	return
}
