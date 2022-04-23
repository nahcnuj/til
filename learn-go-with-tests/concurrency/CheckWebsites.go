package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string // url
	bool   // result
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultChannel := make(chan result)
	for _, url := range urls { // the url variable is re-used
		go func(u string) { // receives a copy of url to make sure the url is fixed for each iteration
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
