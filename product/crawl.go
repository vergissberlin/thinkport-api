package product

import (
	"fmt"
	"log"
	"strings"

	"encore.dev/metrics"
	"github.com/gocolly/colly"
)

type Labels struct {
	Success bool
}

var CrawlRequests = metrics.NewCounterGroup[Labels, uint64]("crawl_requests", metrics.CounterConfig{})

const urlProductTrainings = "https://thinkport.digital/cloud-trainings-workshops/"

// Returns an array of Trainings crawled from a website
func getTrainings() []TrainingStruct {
	var success bool
	c := colly.NewCollector()
	m := make(map[string]TrainingStruct)

	// Find all divs with class "lae-team-training " and print them
	c.OnHTML(".eael-filterable-gallery-item-wrap", func(e *colly.HTMLElement) {
		if e.ChildText(".fg-item-title") != "" {
			// Get the other css classes beside from "eael-filterable-gallery-item-wrap"
			classes := e.Attr("class")
			// Split the classes by space
			classesSplit := strings.Split(classes, " ")
			// Get the last element of the classesSplit array
			// The last element devidided by "-" is the training category
			trainingCategory := strings.Split(classesSplit[len(classesSplit)-1], "-")[2]

			// Append training to slice of trainings if not empty
			m[e.ChildText(".fg-item-title")] = TrainingStruct{
				Name:        e.ChildText(".fg-item-title"),
				Description: e.ChildText(".fg-item-content"),
				Category:    trainingCategory,
				Link:        e.ChildAttr("a[href^=https]", "href"),
			}

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(urlProductTrainings)
	if err != nil {
		log.Println("Error: ", err)
		success = false
	} else {
		log.Println("Trainings crawled: ", len(m))
		success = true
	}

	CrawlRequests.With(Labels{Success: success}).Increment()

	// Convert trainings as array
	trainings := make([]TrainingStruct, 0, len(m))
	for _, v := range m {
		trainings = append(trainings, v)
	}

	return trainings

}
