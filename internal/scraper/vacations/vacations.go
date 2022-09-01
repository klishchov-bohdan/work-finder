package vacations

import (
	"fmt"
	"net/url"
	"strconv"

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
	fmt.Println(vs.URL.Query().Get("page"))
}

func (vs *VacationScraper) Parse() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{vs.URL.String()},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.card.job-link").Each(func(i int, s *goquery.Selection) {
				g.Exports <- map[string]interface{}{
					"text": s.Find("h2 a").Text(),
				}
			})
		},
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}
