package scraper

import (
	"fmt"
)

type Scraper struct {
	url string
}

func (s *Scraper) GetUrl() {
	fmt.Println(s.url)
}

func (s *Scraper) SetUrl(url string) {
	s.url = url
}

