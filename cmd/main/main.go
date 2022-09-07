package main

import (
	"fmt"
	"log"
	"workerFinder/internal/scraper"
	"workerFinder/internal/scraper/vacations"
)

func main() {
	scr := scraper.NewScrapper()
	codes := vacations.NewCategoryCodes()
	vpgs, err := scr.GetVacationPages(2, codes.Accounting, codes.BankingFinance)
	if err != nil {
		log.Fatal(err)
	}
	for _, vp := range vpgs {
		fmt.Println(vp.PageNum)
		for _, v := range vp.Vacations {
			fmt.Println(*v)
		}
	}
}
