package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"log"
	"match_statistics_scrapper/models"
	"strings"
)

func ScrapsBritishBasketBall(url string) {
	page := rod.New().MustConnect().MustPage(url).MustWaitLoad()

	// Get the HTML content after JavaScript execution
	pageStr := page.MustHTML()
	page.MustClose()
	if pageStr == "" {
		fmt.Println("Page not found in British Basket Ball scrapping, please retry")
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	if err != nil {
		log.Fatal(err)
	}

	// Find elements by tag name or class
	var allStat []models.BritishBasketBallStat
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		allTr := s.Find("tbody").Find("tr")
		allTr.Each(func(k int, s2 *goquery.Selection) {
			var trData []string
			allTd := s2.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})
			statBnxt := models.BritishBasketBallStat{
				Team:    strings.TrimSpace(trData[0]),
				Date:    strings.TrimSpace(trData[1]),
				Min:     strings.TrimSpace(trData[2]),
				FgP:     strings.TrimSpace(trData[5]),
				ThreePP: strings.TrimSpace(trData[8]),
				FtP:     strings.TrimSpace(trData[11]),
				Off:     strings.TrimSpace(trData[12]),
				Def:     strings.TrimSpace(trData[13]),
				Reb:     strings.TrimSpace(trData[14]),
				Ast:     strings.TrimSpace(trData[15]),
				Stl:     strings.TrimSpace(trData[16]),
				Blk:     strings.TrimSpace(trData[17]),
				Pf:      strings.TrimSpace(trData[18]),
				To:      strings.TrimSpace(trData[19]),
				Pts:     strings.TrimSpace(trData[20]),
			}
			allStat = append(allStat, statBnxt)
		})

	})
	fmt.Println(allStat)
	return
}
