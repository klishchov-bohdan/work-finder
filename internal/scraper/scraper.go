package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Scraper struct {
	url string
	c   *colly.Collector
}

func NewScrapper(url string) *Scraper {
	return &Scraper{
		url: url,
		c: colly.NewCollector(
			colly.AllowedDomains("www.work.ua"),
			colly.MaxDepth(30),
		),
	}
}

func (scr *Scraper) GetVacationNames() []string {
	var vNames []string
	scr.c.OnHTML("div.card h2 a", func(e *colly.HTMLElement) {
		vNames = append(vNames, e.Text)
	})
	scr.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})
	scr.c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	scr.c.OnHTML("div#pjax-job-list nav ul.pagination li.no-style", func(e *colly.HTMLElement) {
		t := e.ChildAttr("a", "href")
		e.Request.Visit("https://www.work.ua" + t)
	})
	scr.c.Visit(scr.url)
	return vNames
}
