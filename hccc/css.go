package hccc

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// searchCSSClass - search classes in given style files
func searchCSSClass(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("[error] can't open css file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal("[error] can't read css file", err)
	}

	var tmp string
	var str string

	for scanner.Scan() {
		t := scanner.Text()
		tmp = tmp + t
	}

	for {
		fp := strings.Index(tmp, "{")
		lp := strings.Index(tmp, "}")

		if fp >= 0 && lp >= 0 && fp < lp {
			str = str + tmp[:fp]
			tmp = tmp[fp+1:]

		} else if lp >= 0 {
			tmp = tmp[lp+1:]
		} else {
			break
		}
	}

	str = strings.Replace(str, "::before", "", -1)
	str = strings.Replace(str, "::after", "", -1)

	slice := []string{}
	for _, s := range strings.Split(str, " ") {
		if strings.Index(s, ".") >= 0 {
			slice = append(slice, s)
		}
	}
	for _, s := range slice {
		fp := strings.Index(s, ".")
		s = s[fp:]

		fp = strings.Index(s, "#")
		if fp > 0 {
			s = s[:fp]
		}

		fp = strings.Index(s, "+")
		if fp > 0 {
			s = s[:fp]
		}

		for _, as := range strings.Split(s, ".") {
			if len(as) == 0 {
				continue
			}
			list.css = append(list.css, as)
		}
	}
}
