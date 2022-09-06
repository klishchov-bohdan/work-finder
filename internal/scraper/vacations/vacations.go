package vacations

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"workerFinder/internal/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type VacationScraper struct {
	URL         *url.URL
	CurrentPage uint64
}

func NewVacationScraper() *VacationScraper {
	url, err := url.Parse("https://www.work.ua/ru/jobs/?advs=1&category=20+22+14+23+4+2+1+8+24+10+12+3+9+15+19+6+6792+26+5+17+25+21+30+13+27+11+7+18&page=1")
	if err != nil {
		fmt.Println("VacationScraper.NewVacationScraper():", err)
		return nil
	}
	return &VacationScraper{
		URL:         url,
		CurrentPage: 1,
	}
}

func (vs *VacationScraper) NextPage() {
	vs.CurrentPage++
	values, _ := url.ParseQuery(vs.URL.RawQuery)
	values.Set("page", strconv.FormatUint(vs.CurrentPage, 10))
	vs.URL.RawQuery = values.Encode()
}

func (vs *VacationScraper) ParsePage() (vacations []*models.Vacation) {
	geziyor.NewGeziyor(&geziyor.Options{
		RequestDelay: 100 * time.Millisecond,
		StartURLs:    []string{vs.URL.String()},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.card.job-link").Each(func(i int, s *goquery.Selection) {
				if href, ok := s.Find("h2 a").Attr("href"); ok {
					g.Get(r.JoinURL(href), func(_g *geziyor.Geziyor, _r *client.Response) {
						customer, address, conditions, customerDetail := "", "", "", ""
						_r.HTMLDoc.Find("div.card p").Each(func(_i int, _s *goquery.Selection) {
							if attr, exists := _s.Find("span").Attr("title"); exists && attr == "Данные о компании" {
								customer = _s.Find("a b").Text()
								customerDetail = _s.Find("span.add-top-xs").Text()
							}
							if attr, exists := _s.Find("span").Attr("title"); exists && attr == "Адрес работы" {
								_s.Children().RemoveFiltered("span")
								address = _s.Text()
							}
							if attr, exists := _s.Find("span").Attr("title"); exists && attr == "Условия и требования" {
								conditions = _s.Text()
							}
						})
						logo, _ := _r.HTMLDoc.Find("div.card p.logo-job-container a img").Attr("src")
						space := regexp.MustCompile(`\s+`)
						vacation := &models.Vacation{
							Title:          s.Find("h2 a").Text(),
							Salary:         _r.HTMLDoc.Find("div.card p.text-indent.text-muted.add-top-sm b.text-black").Text(),
							SalaryDetail:   _r.HTMLDoc.Find("div.card p.text-indent.text-muted.add-top-sm span.text-muted").Text(),
							Customer:       customer,
							CustomerDetail: space.ReplaceAllString(strings.TrimSpace(strings.ReplaceAll(customerDetail, "\n", "")), " "),
							Address:        strings.TrimSpace(strings.ReplaceAll(address, "\n", "")),
							Conditions:     strings.TrimSpace(strings.ReplaceAll(conditions, "\n", "")),
							Logo:           logo,
						}
						vacations = append(vacations, vacation)
					})
				}
			})
		},
	}).Start()
	return
}

func (vs *VacationScraper) GetMaxPageNum() uint64 {
	var maxPageNum uint64
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{vs.URL.String()},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			var err error
			maxPageNum, err = strconv.ParseUint(r.HTMLDoc.Find("div#pjax-job-list nav ul li:nth-child(6)").Text(), 10, 64)
			if err != nil {
				log.Fatal("VacationScraper.GetMaxPageNum():", err)
			}
		},
	}).Start()
	return maxPageNum
}

func clearOut() {
	err := os.Remove("out.json")
	if err != nil {
		fmt.Println("VacationScraper.clearOut():", err)
		return
	}
}
