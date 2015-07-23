package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "yomikata [word]")
		os.Exit(1)
	}
	word := os.Args[1]
	resp, err := http.Get("http://yomikata.org/word/" + word)
	if err != nil {
		fmt.Fprintln(os.Stderr, os.Args[0]+":", err)
		os.Exit(1)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, os.Args[0]+":", err)
		os.Exit(1)
	}
	if doc.Find("#word").Text() == "" {
		fmt.Fprintln(os.Stderr, "わかりません")
		os.Exit(1)
	}
	doc.Find(".spAns").Each(func(_ int, s *goquery.Selection) {
		fmt.Printf("%s (%s)\n", s.Find(".psAns").Text(), s.Find(".psPt").Text())
	})
}
