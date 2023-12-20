package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"log"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/utils"
	"strings"
)

func ScrapsBnxt(url string) []*models.MatchStatResponse {
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
	var allStat []*models.MatchStatResponse
	doc.Find("#match_stats_table").Each(func(i int, s *goquery.Selection) {
		allTr := s.Find("tr")
		allTr.Each(func(i int, s2 *goquery.Selection) {
			var trData []string
			allTd := s2.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})
			if len(trData) == 28 {
				statResp := &models.MatchStatResponse{
					Date:   utils.BnxtDate(trData[0]),
					Opp:    trData[1],
					Result: trData[2],
					Min:    trData[4],
					FGP:    trData[13],
					FTP:    trData[16],
					ThreeP: trData[10],
					REB:    trData[19],
					AST:    trData[22],
					BLK:    trData[25],
					STL:    trData[23],
					PF:     trData[20],
					TO:     trData[24],
					PTS:    trData[3],
				}
				allStat = append(allStat, statResp)
				fmt.Println(statResp)
			}
		})

	})
	fmt.Println(allStat)
	return allStat
}
