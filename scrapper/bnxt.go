package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"log"
	"match_statistics_scrapper/models"
	"strings"
)

func ScrapsBnxt(url string) []models.BnxtStat {
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

	fmt.Println("Document after request ", doc.Text())
	// Find elements by tag name or class
	var allStat []models.BnxtStat
	doc.Find("#match_stats_table").Each(func(i int, s *goquery.Selection) {
		allTr := s.Find("tr")
		allTr.Each(func(i int, s2 *goquery.Selection) {
			var trData []string
			allTd := s2.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})
			if len(trData) == 28 {
				statBnxt := models.BnxtStat{
					GameDate: trData[0],
					Game:     trData[1],
					Result:   trData[2],
					PTS:      trData[3],
					Min:      trData[4],
					TwoP:     trData[7],
					ThreeP:   trData[10],
					FgP:      trData[13],
					FtP:      trData[16],
					Dr:       trData[17],
					Or:       trData[18],
					Tot:      trData[19],
					Fp:       trData[20],
					Df:       trData[21],
					Ast:      trData[22],
					St:       trData[23],
					To:       trData[24],
					Bs:       trData[25],
					Br:       trData[26],
					Eff:      trData[27],
				}
				allStat = append(allStat, statBnxt)
			}
		})

	})
	fmt.Println(allStat)
	return allStat
}
