package main

import (
	"fmt"
	"log"
	"strings"
	"regexp"
	"encoding/json"
	// "io/ioutil"
	"os"
	"github.com/gocolly/colly"
)

type Article struct {
	Title	string
	Link	string
	LinkType string
}

// type StoryArticle struct {
// 	Title	string
// 	Link	string
// }
// type websiteLinks struct {
// 	Title	string
// 	Link	string
// }
// type AdLinks struct {
// 	Title	string
// 	Link	string
// }

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
		artType := "story"

		if strings.Contains(link, "https://www.marketwatch.com/story/"){
			// article := websiteLinks{
			// 	Title:       title,
			// 	Link:         link,
			// }
			//
			// fmt.Printf("Story Link found: %q -> %s\n", article.Title, article.Link)
			artType = "story"

		} else if strings.Contains(link, "marketwatch.com"){
			// article := StoryArticle{
			// 	Title:       title,
			// 	Link:         link,
			// }
			//
			// fmt.Printf("Website Link found: %q -> %s\n", article.Title, article.Link)
			artType = "other website link"
		} else {
				// article := AdLinks{
				// 	Title:       title,
				// 	Link:         link,
				// }

				// fmt.Printf("Ad Link found: %q -> %s\n", article.Title, article.Link)
				artType = "Ad"
		}

		article := Article{
			Title:       title,
			Link:         link,
			LinkType:	artType,
		}

		JSONarticle, _ := json.MarshalIndent(article, "", " ")
		fmt.Println(string(JSONarticle))
		// err = ioutil.WriteFile("output.json", JSONarticle, 0644)

		f, err := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0660)
		if err != nil {
		    panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(string(JSONarticle)); err != nil {
		    panic(err)
		}
})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.marketwatch.com/")
}
