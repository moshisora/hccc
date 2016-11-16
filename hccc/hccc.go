package hccc

import (
	"fmt"
	"strings"
)

// List : class lists
type List struct {
	classListHTML []string
	classListCSS  []string
}

var list = List{}

// Run - Main Process
func Run(htmlPath, cssPath string) {
	searchHTMLClass(htmlPath)
	searchCSSClass(cssPath)

	hc, cc := compare()

	fmt.Println(color("classes used in html but undefined in css", "yellow"))
	for _, h := range hc {
		fmt.Println(h)
	}

	fmt.Println(color("classes defined in css but not used in html", "yellow"))
	for _, c := range cc {
		fmt.Println(c)
	}
}

func compare() ([]string, []string) {
	hc := strings.Join(list.classListHTML, " ")
	cc := strings.Join(list.classListCSS, " ")

	var (
		hcf []string
		ccf []string
	)

	for _, h := range list.classListHTML {
		if !strings.Contains(cc, h) {
			hcf = append(hcf, h)
		}
	}

	for _, c := range list.classListCSS {
		if !strings.Contains(hc, c) {
			ccf = append(ccf, c)
		}
	}

	return hcf, ccf
}

func color(str, color string) string {
	switch color {
	case "yellow":
		return "\x1b[33m" + str + "\x1b[0m"
	}
	return str
}
