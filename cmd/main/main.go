package main

import (
	"fmt"
	"workerFinder/internal/scraper"
)

func main() {
	url := "https://www.work.ua/ru/jobs/?advs=1&category=20+22+14+23+4+2+1+8+24+10+12+3+9+15+19+6+6792+26+5+17+25+21+30+13+27+11+7+18"
	scr := scraper.NewScrapper(url)
	fmt.Println(len(scr.GetVacationNames()))
}
