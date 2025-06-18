package main

import (
	"flag"
	"fmt"
	"go-core-4/GoSearch/pkg/crawler"
	"go-core-4/GoSearch/pkg/crawler/spider"
	"go-core-4/GoSearch/pkg/index"
	"log"
	"sort"
	"strings"
)

type SortByID []crawler.Document

func (st *SortByID) Sort() {

}

func main() {
	s := flag.String("s", "", "for searching")

	flag.Parse()
	_ = s
	search := spider.New()
	myDocs, err := search.Scan("https://golang.org", 2)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(myDocs, func(i, j int) bool { return myDocs[i].ID < myDocs[j].ID })
	/*for _, doc := range myDocs {
		if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(*s)) {
			fmt.Printf("Искомый ID: %[1]v\nИскомый URL: %[2]v\nИскомый Title: %[3]v\nИскомый Body: %[4]v\n\n", doc.ID, doc.URL, doc.Title, doc.Body)
		}
	}*/

	m := index.IndDocumet(myDocs)
	res := find(m, *s)
	sort.Slice(res,func(i, j int) bool { return res[i]< res[j] })
	fmt.Println(res)
}

// Функция для поиска документов по запросу
func find(index map[string][]int, query string) []int {
	tokens := strings.Fields(strings.ToLower(query))
	results := make(map[int]struct{})

	for _, token := range tokens {
		if docIDs, exists := index[token]; exists {
			for _, id := range docIDs {
				results[id] = struct{}{}
			}
		}
	}

	var resultIDs []int
	for id := range results {
		resultIDs = append(resultIDs, id)
	}

	return resultIDs
}
