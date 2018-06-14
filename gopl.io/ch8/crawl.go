package ch8

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/url"
	"os"
)

func Extract(uri string) ([]string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s : %s", uri, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", uri, err)
	}

	parseUrl, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("parsing url %s as %v", uri, err)
	}
	host := parseUrl.Host

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)

				if err != nil {
					continue
				}

				if host != link.Host {
					continue
				}

				links = append(links, link.String())
			}

		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func crawl()  {
	//worklist := make(chan []string)
	//
	//go func() { worklist <- os.Args[1:]}()
	//
	//seen := make(map[string]bool)
	//for list := range worklist {
	//	for _, link := range list {
	//		if !seen[link] {
	//			seen[link] = true
	//			go func(link string) {
	//				worklist <- ch8.Crawl(link)
	//			}(link)
	//		}
	//	}
	//}


	// buffer
	worklist := make(chan []string, 20)

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true

				go func(link string) {
					worklist <- Crawl(link)
				}(link)
			}
		}
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node))  {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func test() {
	req, err := http.NewRequest("GET", "http://www.catjc.com", nil)

}