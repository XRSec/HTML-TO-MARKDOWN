package main

import (
	"flag"
	"fmt"
	"github.com/JohannesKaufmann/html-to-markdown"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	url                                      string
	output                                   string
	help                                     bool
	version                                  bool
	err                                      error
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

func HtmlToMarkdown(html, output string, export bool) string {
	var (
		markdown string
	)
	converter := md.NewConverter("", true, nil)
	if markdown, err = converter.ConvertString(html); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully converted HTML to Markdown")
	}
	if export {
		if err = ioutil.WriteFile(output, []byte(markdown), 0644); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Successfully export Markdown to %v", output)
		}
	}
	return markdown
}

func GetSource(url, output string, export bool) string {
	var (
		req  *http.Response
		body []byte
	)
	http.TimeoutHandler(nil, time.Duration(10)*time.Second, "请求超时")
	if req, err = http.Get(url); err != nil {
		log.Fatalf("请求错误: Error: [%s]", err)
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			log.Fatalf("关闭请求错误: Error: [%s]", err)
		}
	}(req.Body)
	if body, err = ioutil.ReadAll(req.Body); err != nil {
		log.Fatalf("读取错误: Error: [%s]", err)
	}
	if export {
		if err = ioutil.WriteFile(output, body, 0644); err != nil {
			log.Fatalf("写入错误: Error: [%s]", err)
		} else {
			log.Println("Successfully Download HTML!")
		}
	}
	return string(body)
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

	body := GetSource(url, "", false)
	HtmlToMarkdown(body, output, true)
}
