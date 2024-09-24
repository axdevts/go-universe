package scrap

import (
	"fmt"

	"github.com/anaskhan96/soup"
	"github.com/go-rod/rod"
	"github.com/gocolly/colly"
)

func ScrapDataBySoup() {
	fmt.Println("Scrap by Soup: ... ")
	// Make a GET request
	resp, err := soup.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse the HTML
	doc := soup.HTMLParse(resp)

	// Find and print the title
	fmt.Println("Title:", doc.Find("title").Text())

	// Find all links
	links := doc.FindAll("a")
	for _, link := range links {
		fmt.Println("Link:", link.Text(), link.Attrs()["href"])
	}
}

func ScrapByColly() {
	fmt.Println("Scrap by Colly: ... ")
	// Create a new collector
	c := colly.NewCollector()

	// On every <a> element which has href attribute
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q\n", link)
		// Visit the link found on the page
		e.Request.Visit(link)
	})

	// On every visited page
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// Start scraping the website
	c.Visit("https://google.com")
}

func ScrapByRod() {

	fmt.Println("Scrap by Rod: ... ")
	// Launch browser
	browser := rod.New().MustConnect()

	// Open the page
	page := browser.MustPage("https://google.com")

	// Wait for the page to load
	page.MustWaitLoad()

	// Extract the title
	title := page.MustEval(`document.title`).String()
	println("Page Title: ", title)

}
