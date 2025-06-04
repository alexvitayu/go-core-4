package main

import (
	"flag"
	"fmt"
	"go-core-4/GoSearch/pkg/crawler/spider"
	"log"
	"net/url"
)

/*type Docs struct {
	docs []crawler.Document
}

/*func NewDocs(craw crawler.Interface, url string, depth int) (*Docs, error) {
	docs, err := craw.Scan(url, depth)
	if err != nil {
		log.Fatal(err)
	}
	return &Docs{
		docs: docs,
	}, nil
}*/

/*func NewDocs(s spider.Service, url string, depth int) (*Docs, error) {
	docs, err := s.Scan(url, depth)
	if err != nil {
		log.Fatal(err)
	}
	return &Docs{
		docs: docs,
	}, nil
}*/

func main() {
	s := flag.String("s", "", "for searching")

	flag.Parse()

	baseUrl, err := url.Parse("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("s", *s)
	baseUrl.RawQuery = params.Encode()

	search := spider.New()
	myDocs, err := search.Scan(baseUrl.String(), 1)
	if err != nil {
		log.Fatal(err)
	}
	/*for _, doc := range myDocs {
		if strings.Contains(doc.URL, *s) {
			fmt.Printf(doc.URL)
		}
	}*/
	for _, doc := range myDocs {
		fmt.Println(doc.Title)
		fmt.Println(doc.ID)
		fmt.Println(doc.URL)
		fmt.Println(doc.Body)
	}
}
