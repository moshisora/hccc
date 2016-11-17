package hccc

import (
	"reflect"
	"testing"
)

func TestSetClasses(t *testing.T) {
	list.html = []string{}
	list.css = []string{}

	setClasses("./testdata/html", "./testdata/css")

	expect := []string{
		"class-1",
		"class-2",
		"class-3",
		"class-4",
	}

	if !reflect.DeepEqual(expect, list.html) {
		t.Error("error")
	}
	if !reflect.DeepEqual(expect, list.css) {
		t.Error("error")
	}
}

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
