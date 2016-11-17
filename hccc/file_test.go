package hccc

import (
	"reflect"
	"testing"
)

func TestGeneratePaths(t *testing.T) {
	var (
		slice  []string
		expect []string
	)
	slice = generatePaths("testdata/html/1.html")
	expect = []string{"testdata/html/1.html"}

	if !reflect.DeepEqual(expect, slice) {
		t.Error("error")
	}

	slice = generatePaths("testdata/html")
	expect = []string{
		"testdata/html/1.html",
		"testdata/html/2.html",
	}

	if !reflect.DeepEqual(expect, slice) {
		t.Error("error")
	}
}

func TestIsDirectory(t *testing.T) {
	rd, _ := isDirectory("./testdata/html")
	if !rd {
		t.Error("error")
	}

	rf, _ := isDirectory("./testdata/html/1.html")
	if rf {
		t.Error("error")
	}
}
