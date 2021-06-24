package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func main() {
	resp, err := http.Get("https://www.xcar.com.cn/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("http resp err: %v", err)
		return
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

// determineEncoding 确定网页的编码格式，默认是utf-8
func determineEncoding(bufReader *bufio.Reader) encoding.Encoding {
	bytes, err := bufReader.Peek(1024)
	if err != nil {
		log.Printf("fetcher err: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
