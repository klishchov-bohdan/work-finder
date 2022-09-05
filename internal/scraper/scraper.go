package scraper

import (
	"workerFinder/internal/models"
	"workerFinder/internal/scraper/vacations"
)

type Scraper struct {
	vs vacations.VacationScraper
}

func NewScrapper() *Scraper {
	return &Scraper{
		vs: *vacations.NewVacationScraper(),
	}
}

func (scr *Scraper) GetVacationPages(pageNum uint64) *models.Vacation {
	return &models.Vacation{}
}
