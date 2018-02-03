package main

import (
	"./web"
	"fmt"
	"log"
	. "./util"
)

func main() {
	parseConfig := ReadParseConfig()
	html, err := web.GetHtmlSource(parseConfig.GetNextSiteURL())
	if err != nil {
		log.Println("Site parsing was stoped.")
		log.Println(html)
		return
	}
	fmt.Print(html)
}
