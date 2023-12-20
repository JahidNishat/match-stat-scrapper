package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/utils"
)

func EspnScrap(url string) []*models.MatchStatResponse {
	// creating a new Colly instance
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	var allStatResp []*models.MatchStatResponse
	// scraping logic section.gamelogWidget
	c.OnHTML("section.gamelogWidget", func(e *colly.HTMLElement) {
		// Collect all matching tbody elements
		allTr := e.DOM.Find("tbody").Find("tr")
		allTr.Each(func(i int, element *goquery.Selection) {
			var trData []string
			allTd := element.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})
			if len(trData) == 14 {
				statResp := &models.MatchStatResponse{
					Date:   utils.EspnDate(trData[0]),
					Opp:    trData[1],
					Result: trData[2],
					Min:    trData[3],
					FGP:    trData[4],
					FTP:    trData[5],
					ThreeP: trData[6],
					REB:    trData[7],
					AST:    trData[8],
					BLK:    trData[9],
					STL:    trData[10],
					PF:     trData[11],
					TO:     trData[12],
					PTS:    trData[13],
				}
				allStatResp = append(allStatResp, statResp)
			}
		})
		fmt.Println("Espn Scrapper Result", allStatResp)
	})

	// visiting the target page
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error while visiting url : ", err)
		return nil
	}
	return allStatResp
}
