package main

import (
	"fmt"
	"sync" // sync package provides synchronization primitives (Mutex, WaitGroup)
)

// Fetcher interface defines something that can fetch URLs.
// Any type that implements Fetch() method satisfies this interface.
type Fetcher interface {
	// Fetch takes a URL and returns: page body, links found on that page, and any error
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCache is a concurrency-safe map to track visited URLs.
// Without the mutex, multiple goroutines could read/write the map simultaneously causing a race condition.
type SafeCache struct {
	mu      sync.Mutex          // Mutex to protect concurrent access to the map
	visited map[string]bool     // Map storing URLs we've already crawled
}

// Visit marks a URL as visited and returns false if it was already visited.
// This is a method on SafeCache (c is the receiver).
func (c *SafeCache) Visit(url string) bool {
	c.mu.Lock()         // Acquire the lock - only one goroutine can hold it at a time
	defer c.mu.Unlock() // Release the lock when function returns (defer runs at end of function)
	
	if c.visited[url] { // Check if we've already visited this URL
		return false    // Already visited, return false to skip it
	}
	c.visited[url] = true // Mark as visited
	return true           // First time seeing this URL, return true to process it
}

// Crawl recursively crawls pages starting with url, up to a maximum depth.
// It fetches URLs in parallel using goroutines and avoids fetching the same URL twice.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement WaitGroup counter when this function completes

	if depth <= 0 { // Base case: stop if we've reached max depth
		return
	}

	if !cache.Visit(url) { // Try to mark URL as visited; if already visited, skip
		return
	}

	body, urls, err := fetcher.Fetch(url) // Fetch the page content and links
	if err != nil {                       // Handle fetch errors (e.g., URL not found)
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body) // Print the URL and its content

	for _, u := range urls {                           // For each link found on the page
		wg.Add(1)                                      // Increment WaitGroup counter BEFORE launching goroutine
		go Crawl(u, depth-1, fetcher, cache, wg)       // Launch new goroutine to crawl that link (parallel!)
	}
}

func main() {
	cache := &SafeCache{visited: make(map[string]bool)} // Create shared cache with initialized map
	var wg sync.WaitGroup                                // WaitGroup to wait for all goroutines to finish

	wg.Add(1)                                            // Add 1 to counter for the initial Crawl call
	Crawl("https://golang.org/", 4, fetcher, cache, &wg) // Start crawling from root URL with depth 4

	wg.Wait()
}

// ---------------------------
// Fake fetcher implementation (simulates web requests for testing)
// ---------------------------

// fakeFetcher is a map from URL string to fakeResult pointer
type fakeFetcher map[string]*fakeResult

// fakeResult holds simulated page data
type fakeResult struct {
	body string   // The page content
	urls []string // Links found on the page
}

// Fetch implements the Fetcher interface for fakeFetcher
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {                      // Check if URL exists in our fake data
		return res.body, res.urls, nil              // Return the fake body and links
	}
	return "", nil, fmt.Errorf("not found: %s", url) // URL not in fake data, return error
}

// Pre-populated fake fetcher - simulates a small website structure
var fetcher = fakeFetcher{
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
}
