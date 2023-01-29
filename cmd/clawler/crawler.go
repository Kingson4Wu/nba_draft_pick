package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	res, err := http.Get("https://nba.xxxx.com/teams")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	/*if b, err := io.ReadAll(res.Body); err == nil {
		context := string(b)
		fmt.Println(context)
	}*/

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var relations []string

	doc.Find("div.team").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			l, _ := s.Attr("href")
			ll := strings.Split(l, "/")
			n := ll[len(ll)-1]
			fmt.Println(n)

			img, _ := s.Find("div.img").Find("img").Attr("src")
			fmt.Println(img)

			name := s.Find("div.font").Find("h2").Text()
			fmt.Println(name)

			b := download(img)
			os.WriteFile("./cmd/photos/"+n+".png", b, os.ModePerm)

			relations = append(relations, n+"_"+name)

		})
	})

	os.WriteFile("./cmd/photos/result.txt", []byte(strings.Join(relations, "\r\n")), os.ModePerm)

	//fmt.Println(menu)
}

func download(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	if b, err := io.ReadAll(res.Body); err == nil {
		return b
	}
	return nil
}
