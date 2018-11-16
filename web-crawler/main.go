package main

import (
	"fmt"
	"os"
	"web-crawler/parse"
)

func main() {
	chechArgs()
	url := os.Args[1]
	parse.Crawl(url)
}

func chechArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Insert URL argument")
		os.Exit(1)
	}
}
