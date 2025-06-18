package membot

import (
	"go-core-4/GoSearch/pkg/crawler"
)

// Service - имитация служба поискового робота.
type Service struct{}

// New - констрктор имитации службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(url string, depth int) ([]crawler.Document, error) {

	data := []crawler.Document{
		{
			ID:    0,
			URL:   "https://go.dev",
			Title: "go",
		},
		{
			ID:    1,
			URL:   "https://golang.org",
			Title: "golang",
		},
	}

	return data, nil
}
