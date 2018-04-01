package web

import (
	"io"
	//"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
	. "../entity"
	"net/http"
	"fmt"
	"bytes"
)

const _HTML_ELEMENT_CONCATENATOR = "."
const _HREF_TAG = "href_tag"

func getHtml(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	return response.Body, err
}

func StartParsing(config SiteConfig) (error) {
	for i := 1; i <= config.PagesToParsing; i++ {
		htmlBody, err := getHtml(getNextUrlFromPattern(config.SiteURL, i))
		defer htmlBody.Close()
		if err != nil {
			return nil
		}
		document, err := goquery.NewDocumentFromReader(htmlBody)
		document.Find(getSelector(config.Tags[_HREF_TAG])).Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
		})
	}
	/*htmlBody, err := getHtml(config.SiteURL)
	if err != nil {
		return err
	}
	defer htmlBody.Close()
	document, err := goquery.NewDocumentFromReader(htmlBody)
	document.Find("span.adview_spec__info").Each(func (i int, selection *goquery.Selection) {
		*//*fmt.Println(i)
		fmt.Println(selection.Text())*//*
	})*/
	/*tokenizer := html.NewTokenizer(htmlBody)
	for  {
		fmt.Println(string(tokenizer.Text()))
		tokenizer.Next()
		//parseToken(tokenType)
	}*/
	return nil
}

func getSelector(tag ParseableTag) (string) {
	var buffer bytes.Buffer
	buffer.WriteString(tag.Tag)
	buffer.WriteString(_HTML_ELEMENT_CONCATENATOR)
	buffer.WriteString(tag.Class)
	return buffer.String()
}

func getNextUrlFromPattern(urlPattern string, elementsForPattern interface{}) (string) {
	return fmt.Sprintf(urlPattern, elementsForPattern)
}
