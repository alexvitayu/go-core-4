package main

import (
	"flag"
	"fmt"
	"go-core-4/GoSearch/pkg/crawler/spider"
	"log"
	"strings"
)

func main() {
	s := flag.String("s", "", "for searching")

	flag.Parse()

	search := spider.New()
	myDocs, err := search.Scan("https://golang.org", 2)
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range myDocs {
		if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(*s)) {
			fmt.Printf("Искомый титул: %v\n", doc.Title)
		}
	}

}
