package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	// WebSite スクレイピングするサイトのURL
	WebSite = "http://localhost:8080/"
	// ScrapeTag スクレイピングするタグ
	ScrapeTag = "body #conL a"
)

func exampleScrape() {
	// Request the HTML page.
	res, err := http.Get(WebSite)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // 終了時にclose
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	baseURL, _ := url.Parse(WebSite)

	// Find the review items.
	var titles []string
	var urls []string
	// bodyタグ内のid=conL配下のaタグのテキスト、hrefを取得
	doc.Find(ScrapeTag).Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		href, _ := s.Attr("href")
		absURL := convertURL(baseURL, href)
		titles = append(titles, title)
		urls = append(urls, absURL)
	})

	// Output.
	fmt.Println("Title:\n", titles)
	fmt.Println("URL:\n", urls)
}

// 相対URL -> 絶対URL変換
func convertURL(baseURL *url.URL, webURL string) string {
	targetURL, err := url.Parse(webURL)
	if err != nil {
		return ""
	}

	// 相対URL -> 絶対URLに変換
	absURL := baseURL.ResolveReference(targetURL)
	return absURL.String()
}

func main() {
	exampleScrape()
}
