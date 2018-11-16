package parse

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"web-crawler/utils"

	"golang.org/x/net/html"
)

func Crawl(url string) {
	p := parse(url)
	sitemap := make(map[string]bool)
	getSitemap(&sitemap, p)
	printSitemap(&sitemap)
	printAssetsForAllSitemap(&sitemap, url)
}

// visite URL and parse its HTML
func parse(url string) *html.Node {
	r, err := http.Get(url)
	utils.HandleError(err)
	b, err := html.Parse(r.Body)
	utils.HandleError(err)
	return b
}

// traverse HTML recursively and collect sitemap relative links
func getSitemap(linksMap *map[string]bool, n *html.Node) {
	if n == nil {
		return
	}
	links := *linksMap
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			// links within the same domain
			if a.Key == "href" && strings.HasPrefix(a.Val, "/") {
				if !links[a.Val] {
					links[a.Val] = true
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getSitemap(&links, c)
	}
}

func printSitemap(linksMap *map[string]bool) {
	fmt.Println("Sitemap: ")
	printLinks(linksMap, "")
}

func printAssetsForAllSitemap(sitemap *map[string]bool, baseUrl string) {
	links := *sitemap
	var wg sync.WaitGroup
	for link := range links {
		go printAssetsForLink(link, baseUrl, &wg)
		wg.Add(1)
	}
	wg.Wait()
}

func printAssetsForLink(url, base string, wg *sync.WaitGroup) {
	defer wg.Done()
	absoluteUrl := getAbsoluteUrl(url, base)
	p := parse(absoluteUrl)
	assets := make(map[string]bool)
	getAssets(&assets, p)
	fmt.Printf("Assets on %s \n", url)
	printLinks(&assets, base)
}

// traverse HTML recursively and collect assets links
func getAssets(linksMap *map[string]bool, n *html.Node) {
	if n == nil {
		return
	}
	links := *linksMap
	isValidNodeData := (n.Data == "a" || n.Data == "link" || n.Data == "img" || n.Data == "script")
	if n.Type == html.ElementNode && isValidNodeData {
		for _, a := range n.Attr {
			isValidAttributeKey := (a.Key == "href" || a.Key == "rel" || a.Key == "src")
			if isValidAttributeKey {
				if !links[a.Val] {
					links[a.Val] = true
				}

			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getAssets(&links, c)
	}
}

func printLinks(linksMap *map[string]bool, base string) {
	var lineNumber int
	for link := range *linksMap {
		fmt.Printf("%v - ", lineNumber)
		if base != "" {
			fmt.Println(getAbsoluteUrl(link, base))
		} else {
			fmt.Println(link)
		}
		lineNumber++
	}
}

func getAbsoluteUrl(relativeUrl, baseUrl string) string {
	url, err := url.Parse(relativeUrl)
	if err != nil {
		return ""
	}
	parsedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	url = parsedBaseUrl.ResolveReference(url)
	return url.String()
}
