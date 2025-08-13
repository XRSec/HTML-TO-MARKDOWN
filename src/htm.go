package htm

import (
	md "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var err error

/*
Convert HTML to Markdown

func Convert(html, output string) string
- url: url
- html: html strings
- output: output file path (default: "")
- string: markdown strings
*/
func Convert(url, html, output string) string {
	opt := converter.WithDomain(url)
	if url == "" {
		opt = nil
	}

	markdown, err := md.ConvertString(html, opt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTML → Markdown 成功")

	if output != "" {
		if err := os.WriteFile(output, []byte(markdown), 0644); err != nil {
			log.Fatal(err)
		}
		log.Printf("Markdown 已导出到 %v", output)
	}
	return markdown
}

/*
Get html stings

	func Get(url, output string) string
	- url: url
	- output: output file path (default: "")
	- string: html strings
*/
func Get(url, output string) string {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("请求错误: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取错误: %v", err)
	}

	if output != "" {
		if err := os.WriteFile(output, body, 0644); err != nil {
			log.Fatalf("写入错误: %v", err)
		}
		log.Println("Successfully Download HTML!")
	}
	return string(body)
}
