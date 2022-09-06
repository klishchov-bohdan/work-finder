package main

import (
	"fmt"
	"log"
	"workerFinder/internal/scraper"
)

func main() {
	scr := scraper.NewScrapper()
	vpgs, err := scr.GetVacationPages(10)
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
