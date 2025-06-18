package index

import (
	"fmt"
	"go-core-4/GoSearch/pkg/crawler"
	"strings"
)

func IndDocumet(docs []crawler.Document) map[string][]int {
	invertedIndex := make(map[string][]int)
	//перевожу массив структур в текст
	var b strings.Builder
	for _, doc := range docs {
		b.WriteString(fmt.Sprintf("URL: %[1]v\nTitle: %[2]v\n\n", doc.URL, doc.Title))

		text := b.String()

		//токенизация и нормализация
		tokens := strings.Fields(strings.ToLower(text))

		//уникальные токены
		uniqueTokens := make(map[string]struct{})
		for _, token := range tokens {
			uniqueTokens[token] = struct{}{}
		}

		// Заполнение индекса
		for token := range uniqueTokens {
			invertedIndex[token] = append(invertedIndex[token], doc.ID)
		}
	}
	return invertedIndex
}
