package main

import (
	"fmt"
	"log"
	"strings"
	"regexp"
	"github.com/gocolly/colly"
)

type StoryArticle struct {
	Title	string
	Link	string
}
type websiteLinks struct {
	Title	string
	Link	string
}
type AdLinks struct {
	Title	string
	Link	string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		title := e.Text
        link := e.Attr("href")

		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}

		title = reg.ReplaceAllString(title, " ")

		if strings.Contains(link, "https://www.marketwatch.com/story/"){
			article := websiteLinks{
				Title:       title,
				Link:         link,
			}

			fmt.Printf("Story Link found: %q -> %s\n", article.Title, article.Link)

		} else if strings.Contains(link, "marketwatch.com"){
			article := StoryArticle{
				Title:       title,
				Link:         link,
			}

			fmt.Printf("Website Link found: %q -> %s\n", article.Title, article.Link)
		}else {
				article := AdLinks{
					Title:       title,
					Link:         link,
				}

				fmt.Printf("Ad Link found: %q -> %s\n", article.Title, article.Link)
		}
})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.marketwatch.com/")
}
