package web

import (
	"net/http"
	"io/ioutil"
)

const ERROR_GET_MSG = "Can't do GET request to current URL: "

func GetHtml(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return ERROR_GET_MSG + url, err;
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ERROR_GET_MSG + url, err;
	}
	html := string(bytes)
	return html, nil;
}
