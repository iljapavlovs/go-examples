package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group) //1. Create Error Wait Group instead of "var wg sync.WaitGroup"
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somename.com/",
	}

	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines

		//2. Next, instead of adding every goroutine to the group, call g.Go with the function to be a goroutine. The only requirement is it must have the following signature: func() error. Also, since the ErrGroup will handle when goroutines are completed, there is no need to call wg.Done().
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
