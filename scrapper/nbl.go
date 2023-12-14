package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"log"
	"match_statistics_scrapper/models"
	"strings"
)

func NblScrap(url string) []models.NblStat {

	page := rod.New().MustConnect().MustPage(url).MustWaitLoad()

	// Get the HTML content after JavaScript execution
	pageStr := page.MustHTML()
	page.MustClose()
	if pageStr == "" {
		fmt.Println("Page not found in BNXT scrapping, please retry")
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	if err != nil {
		log.Fatal(err)
	}
	// Find elements by tag name or class
	var allStat []models.NblStat
	doc.Find(".player-game-logs").Each(func(i int, s *goquery.Selection) {
		allTr := s.Find("tbody").Find("tr")
		allTr.Each(func(i int, s2 *goquery.Selection) {
			var trData []string
			allTd := s2.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})
			if len(trData) == 12 {
				statNbl := models.NblStat{
					Date: strings.TrimSpace(trData[0]),
					//Opp:  strings.TrimSpace(trData[1]),
					Min: strings.TrimSpace(trData[2]),
					FgP: strings.TrimSpace(trData[3]),
					FtP: strings.TrimSpace(trData[4]),
					Reb: strings.TrimSpace(trData[5]),
					Ast: strings.TrimSpace(trData[6]),
					Blk: strings.TrimSpace(trData[7]),
					Stl: strings.TrimSpace(trData[8]),
					To:  strings.TrimSpace(trData[9]),
					Pf:  strings.TrimSpace(trData[10]),
					Pts: strings.TrimSpace(trData[11]),
				}
				allStat = append(allStat, statNbl)
			}
		})
	})

	fmt.Println(allStat)
	return allStat
}
