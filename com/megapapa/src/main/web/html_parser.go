package web

import (
	"io"
	"golang.org/x/net/html"
	. "../entity"
	"fmt"
	"net/http"
)

func getHtml(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	return response.Body, err
}

func ParseHtml(url string, config SiteConfig) (error) {
	htmlBody, err := getHtml(url)
	if err != nil {
		return err
	}
	defer htmlBody.Close()
	tokenizer := html.NewTokenizer(htmlBody)
	for {
		tokenType := tokenizer.Next()
		parseToken(tokenType)
	}
	return nil
}

func parseToken(tokenType html.TokenType) {
	fmt.Println(tokenType.String())
}
