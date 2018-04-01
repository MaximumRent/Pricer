package test

import (
	"testing"
	"../main/web"
	)

func TestGetNextUrlFromPattern(t *testing.T) {
	expected := "https://www.example.com/minsk/Квартиры?cu=10"
	actual := web.GetNextUrlFromPattern("https://www.example.com/minsk/Квартиры?cu=%v", 10)
	if (actual != expected) {
		t.Fail()
	}
}
