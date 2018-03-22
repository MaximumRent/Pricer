package web

import (
	"io"
	"golang.org/x/net/html"
	. "../entity"
)

func ParseHtml(htmlBody io.ReadCloser, config SiteConfig) (error) {
	defer htmlBody.Close()
	tokenizer := html.NewTokenizer(htmlBody)

	return nil
}
