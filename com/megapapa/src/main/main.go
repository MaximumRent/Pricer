package main

import (
	"./web"
	"fmt"
)

func main() {
	url := "https://www.kufar.by/минск_город/Квартиры?cu=BYR&o=1"
	html, err := web.GetHtml(url)
	if err != nil {
		return
	}
	fmt.Print(html)
}
