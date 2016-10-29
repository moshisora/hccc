package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moshisora/hccc/hccc"
)

var version string

func main() {
	var (
		help     bool
		version  bool
		htmlPath string
		cssPath  string
	)

	flag.BoolVar(&help, "help", false, "show help message")
	flag.BoolVar(&version, "v", false, "show varsion")
	flag.BoolVar(&version, "version", false, "show version")
	flag.StringVar(&htmlPath, "h", "", "html/template files to search classes")
	flag.StringVar(&htmlPath, "html", "", "html/template files to search classes")
	flag.StringVar(&cssPath, "c", "", "css/style files to search classes")
	flag.StringVar(&cssPath, "css", "", "css/style files to search classes")
	flag.Parse()

	if version {
		showVersion()
	}
	if help {
		showUsage()
	}

	hccc.Run(htmlPath, cssPath)

	fmt.Println("finish")
}

func showVersion() {
	fmt.Println("version:", version)
	os.Exit(0)
}

func showUsage() {
	fmt.Println("usage")
	flag.PrintDefaults()
	os.Exit(1)
}
