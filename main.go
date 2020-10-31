package main

import (
	"./scraper"
	"./server"
)

func main() {
	s := scraper.Scraper{}
	s.SetUrl("UTL")
	serv := server.Server{}
	serv.Start()
}
