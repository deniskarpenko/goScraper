package main

import "./scraper"

func main() {
	s := scraper.Scraper{}
	s.SetUrl("UTL")
	s.GetUrl()
}
