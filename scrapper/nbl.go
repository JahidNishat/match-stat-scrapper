package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/utils"
	"strings"
	"time"
)

func NblScrap(url string) []*models.MatchStatResponse {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.Close()

	page := browser.Timeout(1 * time.Minute).MustPage(url).MustWaitLoad()
	defer page.MustClose()

	// Get the HTML content after JavaScript execution
	pageStr := page.MustHTML()
	if pageStr == "" {
		fmt.Println("Page not found in NBL scrapping, please retry")
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	if err != nil {
		log.Fatal(err)
	}
	// Find elements by tag name or class
	var allStat []*models.MatchStatResponse
	doc.Find(".player-game-logs").Each(func(i int, s *goquery.Selection) {
		allTr := s.Find("tbody").Find("tr")
		allTr.Each(func(i int, s2 *goquery.Selection) {
			var trData []string
			allTd := s2.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})
			if len(trData) == 12 {
				statResp := &models.MatchStatResponse{
					Date: utils.NblDate(strings.TrimSpace(trData[0])),
					//Opp:    strings.TrimSpace(trData[1]),
					Result: "",
					Min:    strings.TrimSpace(trData[2]),
					FGP:    strings.TrimSpace(trData[3]),
					FTP:    strings.TrimSpace(trData[4]),
					ThreeP: "",
					REB:    strings.TrimSpace(trData[5]),
					AST:    strings.TrimSpace(trData[6]),
					BLK:    strings.TrimSpace(trData[7]),
					STL:    strings.TrimSpace(trData[8]),
					PF:     strings.TrimSpace(trData[10]),
					TO:     strings.TrimSpace(trData[9]),
					PTS:    strings.TrimSpace(trData[11]),
				}
				allStat = append(allStat, statResp)
			}
		})
	})

	fmt.Println("NBL Scrapper result : ", allStat)
	return allStat
}
