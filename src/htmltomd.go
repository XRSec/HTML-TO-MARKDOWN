package htmltomd

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var err error

func Convert(html, output string, export bool) string {
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

func Get(url, output string, export bool) string {
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
