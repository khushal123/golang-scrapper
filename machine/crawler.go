package machine

import (
	"net/url"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

// RunCrawler utilizes the colly package to crawl the given rootURL looking for <a href=""> tags
func RunCrawler(rootURL string, maxDepth int) ([]string, error) {
	if maxDepth == 0 {
		return nil, nil
	}
	rootHost := hostName(rootURL)
	c := colly.NewCollector(
		colly.MaxDepth(maxDepth),
		colly.AllowedDomains(rootHost),
	)

	var linksMap = make(map[string]bool)
	linksMap[rootURL] = true

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link == "" || link == "/" {
			return
		}
		if relativeURL(link) {
			linksMap[rootURL+link] = true
		} else if sameHost(link, rootHost) {
			linksMap[link] = true
		}
		e.Request.Visit(link)
	})

	err := c.Visit(rootURL)
	if err != nil {
		return nil, err
	}
	var links []string
	for k := range linksMap {
		links = append(links, k)
	}
	sort.Strings(links)
	return links, nil
}

// check if the link is a relative url
func relativeURL(url string) bool {
	return !strings.HasPrefix(url, "http")
}

// check if the link is from the same host as the rootURL
func sameHost(link, rootHost string) bool {
	return hostName(link) == rootHost
}

// get the hostname with port from a full url
func hostName(fullURL string) string {
	u, err := url.Parse(fullURL)
	if err != nil {
		panic(err)
	}
	return u.Host
}
