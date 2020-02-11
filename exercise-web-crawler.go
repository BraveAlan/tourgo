package main

import (
	"fmt"
	"sync"
)

// fetcher 是填充后的 fakeFetcher。
var (
	urlRecord = URLRecord{v: make(map[string]bool)}
	fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
})

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type URLRecord struct {
	v map[string]bool
	mux sync.Mutex
	wg sync.WaitGroup
}

func (c *URLRecord) setRecord(key string, state bool) {
	c.mux.Lock()
	c.v[key] = state
	c.mux.Unlock()
}

func (c *URLRecord) value(key string) (bool, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	v, ok := c.v[key]
	return v, ok
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。

	defer urlRecord.wg.Done()

	if depth <= 0 {
		return
	}

	if _, ok := urlRecord.value(url); ok {
		return
	}
	urlRecord.setRecord(url, true)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		urlRecord.mux.Lock()
		if _, ok := urlRecord.v[u]; !ok {
			urlRecord.wg.Add(1)
			go Crawl(u, depth-1, fetcher)
		}
		urlRecord.mux.Unlock()
	}
	return
}

func main() {
	urlRecord.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	urlRecord.wg.Wait()
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult



func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}


