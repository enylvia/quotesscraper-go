package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Quotes struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	allQuotes := make([]Quotes, 0)

	collector := colly.NewCollector(colly.AllowedDomains("jagokata.com", "https://jagokata.com"))
	collector.OnHTML(".citatenlijst-auteurs li", func(h *colly.HTMLElement) {

		quoteText := h.ChildText("q")
		quoteAuthor := h.ChildText("h5")
		quotes := Quotes{
			Quote:  quoteText,
			Author: quoteAuthor,
		}
		allQuotes = append(allQuotes, quotes)
	})
	for i := 0; i <= 42; i++ {
		collector.Visit("https://jagokata.com/kata-bijak/dari-tere_liye.html?page=" + strconv.Itoa(i))
	}
	writeJson(allQuotes)

	log.Printf("Scraping Complete \n")
}

func writeJson(data []Quotes) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create Json File")
		return
	}

	_ = ioutil.WriteFile("quotestereliye.json", file, 0644)
}
