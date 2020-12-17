package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)
	for _, u := range urls {
		go func(url string) {
			r := wc(url)
			resultChan <- result{url, r}
		}(u)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.string] = r.bool
	}
	return results
}
