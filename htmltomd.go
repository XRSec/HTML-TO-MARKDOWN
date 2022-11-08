/*
	Package main converts html to markdown.

		import htm "github.com/XRSec/HTML-TO-MARKDOWN/src"

		body := htm.Get("https://www.baidu.com/", "")
		htm.Convert(body, "baidu.md")
*/
package main

import (
	"flag"
	"fmt"
	htm "github.com/XRSec/HTML-TO-MARKDOWN/src"
	"log"
)

var (
	url                                      string
	output                                   string
	help                                     bool
	version                                  bool
	buildTime, commitId, versionData, author string
)

func init() {
	log.SetPrefix("[HTML-TO-MARKDOWN] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.StringVar(&url, "url", "https://www.baidu.com", "The specified url")
	flag.StringVar(&url, "u", "https://www.baidu.com", "The specified url")
	flag.StringVar(&output, "output", "index.md", "File import Path")
	flag.StringVar(&output, "o", "index.md", "File import Path")
	flag.BoolVar(&help, "help", false, "Display help information")
	flag.BoolVar(&help, "h", false, "Display help information")
	flag.BoolVar(&version, "version", false, "HTML-TO-MARKDOWN version")
	flag.BoolVar(&version, "v", false, "HTML-TO-MARKDOWN version")
}

func main() {
	flag.Parse()
	if version {
		fmt.Printf("\n ╷──────────────────────────────────────────────────────────────────────────────╷ \n")
		fmt.Printf(" │                                                                              │\n")
		fmt.Printf(" │  HTML-TO-MARKDOWN                                                            │\n")
		fmt.Printf(" │  Version: %6v\t | BuildTime: %18v                       │\n", versionData, buildTime)
		fmt.Printf(" │  Author: %7v\t | CommitId: %41v  │\n", author, commitId)
		fmt.Printf(" │                                                                              │\n")
		fmt.Printf(" ╵──────────────────────────────────────────────────────────────────────────────╵ \n\n")
		return
	}
	if help || url == "" {
		flag.Usage()
		return
	}

	fmt.Printf("\n ╷──────────────────────────────────────────────────────────────╷ \n")
	fmt.Printf(" │                                                              │\n")
	fmt.Printf(" │  HTML-TO-MARKDOWN                                            │\n")
	fmt.Printf(" │  URL: %50v     │\n", url)
	fmt.Printf(" │  Output: %10v                                          │\n", output)
	fmt.Printf(" │                                                              │\n")
	fmt.Printf(" ╵──────────────────────────────────────────────────────────────╵ \n\n")

	body := htm.Get(url, "")
	htm.Convert(body, output)
}
