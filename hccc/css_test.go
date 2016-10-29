package hccc

import (
	"reflect"
	"testing"
)

func TestSearchCSSClass(t *testing.T) {
	searchCSSClass("./testdata/test.css")

	s := []string{
		"class-single",
		"class-single-1",
		"class-single-2",
		"class-multi",
		"class-multi-1",
		"class-multi-2",
		"class-multi-3",
		"class-complex-1",
		"class-complex-2",
		"class-tag-before",
		"class-tag-after",
		"class-before",
		"class-after",
		"class-id",
	}

	if !reflect.DeepEqual(s, list.classListCSS) {
		t.Error("error")
	}
}
