package main

import (
	. "./util"
)

func main() {
	parseConfig := ReadParseConfig()
	var endFlag bool
	var url string
	for !endFlag {
		//url, endFlag = parseConfig.GetNextSiteURL()
		/*html, err := web.GetHtmlSource(url)
		if err != nil {
			log.Println("Site parsing was stoped.")
			log.Println(html)
			log.Println(err)
			return
		}
		fmt.Print(html)*/
	}

}
