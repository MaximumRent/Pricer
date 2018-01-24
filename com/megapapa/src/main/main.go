package main

import (
	"./util"
	"fmt"
)

func main() {
	url := "https://www.kufar.by/минск_город/Квартиры?cu=BYR&o=1"
	html, err := util.GetHtml(url)
	if err != nil {
		return
	}
	fmt.Print(html)
}
