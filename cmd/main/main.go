package main

import (
	"workerFinder/internal/scraper/vacations"
)

func main() {
	scr := vacations.NewVacationScraper()
	scr.Parse()
	scr.NextPage()
	scr.Parse()
}
