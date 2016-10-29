package hccc

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
}
