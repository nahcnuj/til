package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls { // the url variable is re-used
		go func() {
			results[url] = wc(url)
		}()
	}
	time.Sleep(2 * time.Second)
	return results
}
