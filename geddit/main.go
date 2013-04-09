package main

import (
	"fmt"
	"github.com/matalangilbert/reddit"
	"log"
	"os"
)

// Takes a command line argument naming the subreddit to fetch
func main() {
	var subreddit string

	if len(os.Args) > 1 {
		subreddit = os.Args[1]
	} else {
		subreddit = "hot"
	}

	items, err := reddit.Get(subreddit)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
