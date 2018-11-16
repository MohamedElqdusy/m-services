package parse

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestGetSitemap(t *testing.T) {
	page1 := strings.NewReader(`<!DOCTYPE html>
    <head>
        <title>Web Crawler Test</title>
        <link href="/styles.css"/>
        <script src="/scripts.js"></script>
    </head>
    <body>
        <a href="/page-1.html">page 1</a>
        <a href="/page-2.html">page 2</a>

        <img alt="" src="/image.jpg"/>
        <a href="http://facebook.com">facebook</a>
        <a href="http://twitter.com">Twitter</a>

    </body>
	</html>`)
	expectedSitemap1 := &map[string]bool{
		"/page-1.html": true,
		"/page-2.html": true,
	}
	page2 := strings.NewReader(`<!DOCTYPE html>
    <head>
        <title>Web Crawler Test</title>
        <link href="/styles.css"/>
        <script src="/scripts.js"></script>
    </head>
    <body>
       hhhhhhhh
    </body>
</html>`)

	expectedSitemap2 := &map[string]bool{}

	tt := []struct {
		Name            string
		Page            *strings.Reader
		ExpectedSitemap *map[string]bool
	}{{Name: "page with links", Page: page1, ExpectedSitemap: expectedSitemap1},
		{Name: "page without links", Page: page2, ExpectedSitemap: expectedSitemap2}}
	for _, tc := range tt {
		b, err := html.Parse(tc.Page)
		if err != nil {
			t.Fatalf("Test '%s' failed with %v", tc.Name, err)
		}
		sitemap := make(map[string]bool)
		getSitemap(&sitemap, b)
		for link := range sitemap {
			expectedLinks := *tc.ExpectedSitemap
			if !expectedLinks[link] {
				t.Fatalf("Test '%s' failed with wrong link %s", tc.Name, link)
			}
		}
	}
}

func TestGetAssets(t *testing.T) {
	page1 := strings.NewReader(`<!DOCTYPE html>
    <head>
        <title>Web Crawler Test</title>
        <link href="/styles.css"/>
        <script src="/scripts.js"></script>
    </head>
    <body>
        <a href="/page-1.html">page 1</a>
        <a href="/page-2.html">page 2</a>

        <img alt="" src="/image.jpg"/>
        <a href="http://facebook.com">facebook</a>
        <a href="http://twitter.com">Twitter</a>

    </body>
</html>`)
	expectedAssets1 := &map[string]bool{
		"/page-1.html":        true,
		"/page-2.html":        true,
		"http://facebook.com": true,
		"http://twitter.com":  true,
		"/image.jpg":          true,
		"/scripts.js":         true,
		"/styles.css":         true,
	}
	page2 := strings.NewReader(`<!DOCTYPE html>
	<head>
		<title>Web Crawler Test</title>
	</head>
	<body>
	Anything
	</body>
	</html>`)
	expectedAssets2 := &map[string]bool{}
	tt := []struct {
		Name           string
		Page           *strings.Reader
		ExpectedAssets *map[string]bool
	}{{Name: "page with assets", Page: page1, ExpectedAssets: expectedAssets1},
		{Name: "page without assets", Page: page2, ExpectedAssets: expectedAssets2}}
	for _, tc := range tt {
		b, err := html.Parse(tc.Page)
		if err != nil {
			t.Fatalf("Test '%s' failed with %v", tc.Name, err)
		}
		assets := make(map[string]bool)
		getAssets(&assets, b)
		for link := range assets {
			expectedLinks := *tc.ExpectedAssets
			if !expectedLinks[link] {
				t.Fatalf("Test '%s' failed, cannot find link %s", tc.Name, link)
			}
		}
	}
}

func TestGetAbsoluteUrl(t *testing.T) {
	tt := []struct {
		Name        string
		Url         string
		Base        string
		ExpectedUrl string
	}{{Name: "first test", Base: "http://example.com", Url: "/page", ExpectedUrl: "http://example.com/page"},
		{Name: "second test", Base: "http://example.com", Url: "http://test.com/page", ExpectedUrl: "http://test.com/page"},
		{Name: "third test", Base: "", Url: "/page", ExpectedUrl: "/page"}}
	for _, tc := range tt {
		absoluteUrl := getAbsoluteUrl(tc.Url, tc.Base)
		if absoluteUrl != tc.ExpectedUrl {
			t.Fatalf("Test '%s' failed, Expected %s Found %s", tc.Name, tc.ExpectedUrl, absoluteUrl)
		}
	}
}
