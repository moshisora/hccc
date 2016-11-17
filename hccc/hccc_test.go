package hccc

import (
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	list.html = []string{}
	list.css = []string{}

	searchHTMLClass("./testdata/compare.html")
	searchCSSClass("./testdata/compare.css")

	hc, cc := compare()

	hce := []string{"class-1"}
	cce := []string{"class-4"}

	if !reflect.DeepEqual(hce, hc) {
		t.Error("error")
	}
	if !reflect.DeepEqual(cce, cc) {
		t.Error("error")
	}
}

func TestColor(t *testing.T) {
	str := color("str", "yellow")
	result := "\x1b[33mstr\x1b[0m"

	if str != result {
		t.Error("error")
	}
}
