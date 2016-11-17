package hccc

import (
	"fmt"
	"strings"
)

// ClassList : class lists
type ClassList struct {
	html []string
	css  []string
}

var list = ClassList{}

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
	hc := strings.Join(list.html, " ")
	cc := strings.Join(list.css, " ")

	var (
		hcf []string
		ccf []string
	)

	for _, h := range list.html {
		if !strings.Contains(cc, h) {
			hcf = append(hcf, h)
		}
	}

	for _, c := range list.css {
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
