package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var (
	exact = flag.Bool("exact", false, "exact")
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "yomikata [word]")
		os.Exit(1)
	}
	word := flag.Arg(0)
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
	if *exact {
		fmt.Println(doc.Find(".spAns .psAns").First().Text())
	} else {
		doc.Find(".spAns").Each(func(_ int, s *goquery.Selection) {
			fmt.Printf("%s (%s)\n", s.Find(".psAns").Text(), s.Find(".psPt").Text())
		})
	}
}
