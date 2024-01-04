package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

// CrawlWebpage craws the given rootURL looking for <a href=""> tags
// that are targeting the current web page, either via an absolute url like http://mysite.com/mypath or by a relative url like /mypath
// and returns a sorted list of absolute urls  (eg: []string{"http://mysite.com/1","http://mysite.com/2"})
func CrawlWebpage(rootURL string, maxDepth int) ([]string, error) {
	fmt.Println(parseMasterDomain(rootURL))
	c := colly.NewCollector(
		colly.MaxDepth(maxDepth),
		colly.AllowedDomains(parseMasterDomain(rootURL)),
	)

	var links []string
	var linksMap = make(map[string]bool)
	linksMap[rootURL] = true

	links = append(links, rootURL)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Check for relative URLs or links belonging to the same master domain
		if maxDepth == 0 || link == "" || link == "/" {
			return
		}

		if relativeURL(link) {
			links = append(links, rootURL+link)
		} else {
			links = append(links, link)
		}
		maxDepth = maxDepth - 1
		e.Request.Visit(link)
	})

	err := c.Visit(rootURL)
	maxDepth = maxDepth - 1
	if err != nil {
		return nil, err
	}

	return links, nil
}

func relativeURL(url string) bool {
	return !strings.HasPrefix(url, "http")
}

func parseMasterDomain(fullURL string) string {
	u, err := url.Parse(fullURL)
	if err != nil {
		panic(err)
	}
	return u.Hostname()
}

// --- DO NOT MODIFY BELOW ---

func main() {
	const (
		defaultURL      = "https://cube.dev"
		defaultMaxDepth = 1
	)
	urlFlag := flag.String("url", defaultURL, "the url that you want to crawl")
	maxDepth := flag.Int("depth", defaultMaxDepth, "the maximum number of links deep to traverse")
	flag.Parse()

	links, err := CrawlWebpage(*urlFlag, *maxDepth)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	fmt.Println("Links")
	fmt.Println("-----")
	for i, l := range links {
		fmt.Printf("%03d. %s\n", i+1, l)
	}
	fmt.Println()
}
