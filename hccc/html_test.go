package hccc

import (
	"reflect"
	"testing"
)

func TestSearchHTMLClass(t *testing.T) {
	list.classListHTML = []string{}

	searchHTMLClass("./testdata/list.html")

	s := []string{
		"main",
		"section",
		"section--first",
		"header",
		"content",
		"content__text",
	}

	if !reflect.DeepEqual(s, list.classListHTML) {
		t.Error("error")
	}
}
