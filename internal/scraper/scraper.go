package scraper

import (
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

type Scraper struct {
	url string
}

func NewScrapper(url string) *Scraper {
	return &Scraper{
		url: "https://www.work.ua/ru/jobs/?advs=1&category=20+22+14+23+4+2+1+8+24+10+12+3+9+15+19+6+6792+26+5+17+25+21+30+13+27+11+7+18",
	}
}

func (scr *Scraper) GetVacationNames() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{scr.url},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {

		},
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}
