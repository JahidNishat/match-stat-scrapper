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

func ScrapsBritishBasketBall(url string) []*models.MatchStatResponse {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.Close()

	page := browser.Timeout(1 * time.Minute).MustPage(url).MustWaitLoad()
	defer page.MustClose()

	// Get the HTML content after JavaScript execution
	pageStr := page.MustHTML()
	if pageStr == "" {
		fmt.Println("Page not found in British Basket Ball scrapping, please retry")
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageStr))
	if err != nil {
		log.Fatal(err)
	}

	// Find elements by tag name or class
	var allStat []*models.MatchStatResponse
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		allTr := s.Find("tbody").Find("tr")
		allTr.Each(func(k int, s2 *goquery.Selection) {
			var trData []string
			allTd := s2.Find("td")
			allTd.Each(func(j int, element2 *goquery.Selection) {
				trData = append(trData, element2.Text())
			})

			statResp := &models.MatchStatResponse{
				Date:   utils.BritishBasketBallDate(strings.TrimSpace(trData[1])),
				Opp:    strings.TrimSpace(trData[0]),
				Result: "",
				Min:    strings.TrimSpace(trData[2]),
				FGP:    strings.TrimSpace(trData[5]),
				FTP:    strings.TrimSpace(trData[11]),
				ThreeP: strings.TrimSpace(trData[8]),
				REB:    strings.TrimSpace(trData[14]),
				AST:    strings.TrimSpace(trData[15]),
				BLK:    strings.TrimSpace(trData[17]),
				STL:    strings.TrimSpace(trData[16]),
				PF:     strings.TrimSpace(trData[18]),
				TO:     strings.TrimSpace(trData[19]),
				PTS:    strings.TrimSpace(trData[20]),
			}
			allStat = append(allStat, statResp)
			//fmt.Println(*statResp)
		})

	})
	fmt.Println("British Basketball Scrapper result : ", allStat)
	return allStat
}
