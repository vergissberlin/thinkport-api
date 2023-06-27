package product

import (
	"fmt"
	"strings"

	"encore.dev/metrics"
	"github.com/gocolly/colly"
)

const urlProductTrainings = "https://thinkport.digital/cloud-trainings-workshops/"

var RequestsMade = metrics.NewCounter[uint64]("request_member", metrics.CounterConfig{})

// Returns an array of Trainings crawled from a website
func getTrainings() []TrainingStruct {
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
		RequestsMade.Increment()
	})

	var err = c.Visit(urlProductTrainings)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Convert trainings as array
	trainings := make([]TrainingStruct, 0, len(m))
	for _, v := range m {
		trainings = append(trainings, v)
	}

	return trainings

}
