package main

import (
	"fmt"
	"github.com/matalangilbert/reddit"
	"log"
)

func main() {
	items, err := reddit.Get("hot")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
