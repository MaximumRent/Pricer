package web

import (
	"io"
	"github.com/PuerkitoBio/goquery"
	. "../entity"
	"net/http"
	"fmt"
	"bytes"
	"log"
	"../util"
	"../unique_parsers"
	"../db"
)

const _HTML_ELEMENT_CONCATENATOR = "."
const _HREF_TAG = "href_tag"

func getHtml(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	return response.Body, err
}

/* Start parsing site from configuration */
func StartParsing(config SiteConfig) (error) {
	log.Printf("Start parsing: %s", config.SiteName)
	for i := 1; i <= config.PagesToParsing; i++ {
		url := GetNextUrlFromPattern(config.SiteURL, i)
		htmlBody, err := getHtml(url)
		defer htmlBody.Close()
		if err != nil {
			return nil
		}
		document, err := goquery.NewDocumentFromReader(htmlBody)
		document.Find(getSelector(config.Tags[_HREF_TAG])).Each(func(i int, selection *goquery.Selection) {
			href, exists := selection.Attr("href")
			if exists {
				parseResult, err := parseApartmentPage(href, config)
				if err == nil {
					db.WriteParseResultToDb(parseResult)
				}
			}
		})
	}
	// Make normal return
	return nil
}

/* Called for each apartment page */
func parseApartmentPage(url string, config SiteConfig) (*ParseResult, error) {
	log.Printf("Start parsing apartment... - %s", url)
	htmlBody, err := getHtml(url)
	defer htmlBody.Close()
	if err != nil {
		return nil, err
	}
	document, err := goquery.NewDocumentFromReader(htmlBody)
	result := new(ParseResult)
	for _, tag := range config.Tags {
		// Now, non-repeatable tag is price_tag
		if !tag.Repeatable {
			document.Find(getSelector(tag)).Each(func(i int, selection *goquery.Selection) {
				if util.IsNonTrashString(selection.Text()) {
					// Add price extractor
					value := util.PickoutLettersFromString(selection.Text())
					callPriceExtractor(config.SiteName, value, result)
				}
			})
		} else {
			document.Find(getSelector(tag)).Each(func(i int, selection *goquery.Selection) {
				if util.IsNonTrashString(selection.Text()) {
					value := util.PickoutLettersFromString(selection.Text())
					callSiteParser(config.SiteName, tag, value, result)
					//fmt.Println(util.PickoutLettersFromString(selection.Text()))
				}
			})
		}
	}
	return result, nil
}

func callPriceExtractor(siteName string, value string, result *ParseResult) {
	switch siteName {
	case "kufar.by":
		unique_parsers.ExtractKufarPrice(value, result)
	}
}

// For each site need to create unique parser for non-repeatable tags
func callSiteParser(siteName string, tag ParseableTag, value string, result *ParseResult) {
	switch siteName {
		case "kufar.by":
			unique_parsers.ParseKufar(tag, value, result)
	}
}

func getSelector(tag ParseableTag) (string) {
	var buffer bytes.Buffer
	buffer.WriteString(tag.Tag)
	buffer.WriteString(_HTML_ELEMENT_CONCATENATOR)
	buffer.WriteString(tag.Class)
	return buffer.String()
}

func GetNextUrlFromPattern(urlPattern string, elementsForPattern interface{}) (string) {
	return fmt.Sprintf(urlPattern, elementsForPattern)
}
