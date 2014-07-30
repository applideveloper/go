package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	fetched := map[string]bool{url: true}

	if depth <= 0 {
		return
	}

	c := make(chan []string)
	urls := []string{url}
	for i := 0; i < depth; i++ {
		var next []string

		for _, u := range urls {
			go _crawl(u, fetcher, c)
		}

		for j := 0; j < len(urls); j++ {
			res := <-c
			for _, r := range res {
				if !fetched[r] {
					fetched[r] = true
					next = append(next, r)
				}
			}
		}
		urls = next
	}
}

func _crawl(url string, fetcher Fetcher, c chan []string) {
	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		c <- []string{}
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	c <- urls
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

/*
Exercise: Equivalent Binary Trees

1. Walk 関数を実装してみてください。

2. 関数をテストしてみてください。

関数 tree.New(k) は、値( k, 2k, 3k, ..., 10k )をもつ、ランダムに構造化されたバイナリツリーを生成します。

新しいチャネル ch を生成し、 Walk をはじめてみましょう。

go Walk(tree.New(1), ch)
そして、そのチャネルから読み出し、10個の値を表示してみてください。 それは、 1, 2, 3, ..., 10 という表示になるでしょう。

3. t1 と t2 が同じ値を保存しているどうかを判断するため、 Walk を使って、 Same 関数を実装してみてください。

4. Same 関数をテストしてみてください。

Same(tree.New(1), tree.New(1)) は、 true を返す必要があります。 また、 Same(tree.New(1), tree.New(2)) は、 false を返す必要があります。
*/
