package main

import (
	"fmt"
	"workerFinder/internal/scraper/vacations"
)

func main() {
	scr := vacations.NewVacationScraper()
	fmt.Println(scr.GetMaxPageNum())
}
