package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls { // the url variable is re-used
		go func(u string) { // receives a copy of url to make sure the url is fixed for each iteration
			results[u] = wc(u)
		}(url)
	}
	time.Sleep(2 * time.Second)
	return results
}
