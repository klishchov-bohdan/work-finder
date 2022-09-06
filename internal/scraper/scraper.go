package scraper

import (
	"errors"
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

func (scr *Scraper) GetVacationPages(pageCount uint64) (vacationPages []*models.VacationPage, err error) {
	if pageCount > scr.vs.GetMaxPageNum() {
		return nil, errors.New("Scraper.GetVacationPages(): the page number can`t be bigger then page count")
	}
	if pageCount == 0 {
		pageCount = scr.vs.GetMaxPageNum()
	}
	for scr.vs.CurrentPage = 1; scr.vs.CurrentPage <= pageCount; scr.vs.NextPage() {
		vacationPages = append(vacationPages, &models.VacationPage{
			Vacations: scr.vs.ParsePage(),
			PageNum:   scr.vs.CurrentPage,
		})
	}
	return
}
