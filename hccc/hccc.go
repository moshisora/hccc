package hccc

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var p = Parser{}

// SearchHTMLClasses - search classes in given html files
func SearchHTMLClasses(path string) {
	fmt.Println("search html classes")
	fmt.Printf("html: %s\n", path)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal("[error] can't open html file", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("[error] can't read html file", err)
	}
}

// SearchCSSClasses - search classes in given style files
func SearchCSSClasses(path string) {
	fmt.Println("search css classes")
	fmt.Printf("css : %s\n", path)
}
