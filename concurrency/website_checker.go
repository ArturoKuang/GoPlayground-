package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result, 3)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
			// fmt.Printf("send %q\n", u)
		}(url)
	}

	// fmt.Println("start for")
	for i := 0; i < len(urls); i++ {
		// fmt.Printf("hit %d\n", i)
		r := <-resultChannel
		// fmt.Printf("recieve %q \n", r.string)
		results[r.string] = r.bool
	}

	return results
}

