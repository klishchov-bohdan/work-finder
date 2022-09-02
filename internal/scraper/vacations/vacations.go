package vacations

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
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

func (vs *VacationScraper) ParsePage() {
	clearOut()
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{vs.URL.String()},
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
						g.Exports <- map[string]interface{}{
							"title":           s.Find("h2 a").Text(),
							"sallary":         _r.HTMLDoc.Find("div.card p.text-indent.text-muted.add-top-sm b.text-black").Text(),
							"sallary_detail":  _r.HTMLDoc.Find("div.card p.text-indent.text-muted.add-top-sm span.text-muted").Text(),
							"customer":        customer,
							"address":         strings.TrimSpace(strings.ReplaceAll(address, "\n", "")),
							"logo":            logo,
							"conditions":      strings.TrimSpace(strings.ReplaceAll(conditions, "\n", "")),
							"customer_detail": space.ReplaceAllString(strings.TrimSpace(strings.ReplaceAll(customerDetail, "\n", "")), " "),
						}
					})
				}
			})
		},
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func (vs *VacationScraper) ParsePages() {

}

func clearOut() {
	err := os.Remove("out.json")
	if err != nil {
		fmt.Println("VacationScraper.clearOut():", err)
		return
	}
}
