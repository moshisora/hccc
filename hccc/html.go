package hccc

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// searchHTMLClass - search classes in given html files
func searchHTMLClass(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("[error] can't open html file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal("[error] can't read html file", err)
	}

	for scanner.Scan() {
		t := scanner.Text()

		if strings.Index(t, " class=\"") == -1 {
			continue
		}

		s := strings.Split(t, " class=\"")

		i := strings.Index(s[1], "\"")
		classes := s[1][:i]

		for _, c := range strings.Split(classes, " ") {
			list.html = append(list.html, c)
		}
	}
}
