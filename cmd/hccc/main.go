package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moshisora/hccc/hccc"
)

var (
	version   string
	buildDate string
)

func main() {
	var (
		help        bool
		showVersion bool
		htmlPath    string
		cssPath     string
	)

	flag.BoolVar(&help, "help", false, "show help message")
	flag.BoolVar(&showVersion, "v", false, "show varsion")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.StringVar(&htmlPath, "h", "", "html/template files to search classes")
	flag.StringVar(&htmlPath, "html", "", "html/template files to search classes")
	flag.StringVar(&cssPath, "c", "", "css/style files to search classes")
	flag.StringVar(&cssPath, "css", "", "css/style files to search classes")
	flag.Parse()

	if showVersion {
		printVersion()
	}
	if help {
		printUsage()
	}

	hccc.Run(htmlPath, cssPath)
}

func printVersion() {
	fmt.Println("version: ", version)
	fmt.Println("build date: ", buildDate)
	os.Exit(0)
}

func printUsage() {
	fmt.Println("usage")
	flag.PrintDefaults()
	os.Exit(1)
}
