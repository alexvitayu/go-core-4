package spider

import (
	"go-core-4/GoSearch/pkg/crawler"
	"testing"
)

func TestNew(t *testing.T) {
	//Arrange
	url := "https://golang.org"
	expect := []crawler.Document{}

	//Act
	mySearch := New()
	data, err := mySearch.Scan(url, 1)

	//Assert
	if err != nil {
		t.Errorf("Ожидалось %v получили %v", expect, data)
	}
}
