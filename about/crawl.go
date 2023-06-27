package about

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const urlAboutLocations = "https://thinkport.digital/kontaktieren-3/"

// body > div.elementor.elementor-7608 > div > div > section.elementor-section.elementor-top-section.elementor-element.elementor-element-6634598.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div > div > section > div > div > div.elementor-column.elementor-col-25.elementor-inner-column.elementor-element.elementor-element-05ba45e > div > div > div.elementor-element.elementor-element-055f9c6.elementor-widget.elementor-widget-heading > div > h2
// Returns an array of Locations crawled from a website
func getLocations() []LocationStruct {
	c := colly.NewCollector()
	m := make(map[string]LocationStruct)

	c.OnHTML(".elementor-element-4b29ef2  .elementor-column-wrap", func(e *colly.HTMLElement) {

		if e.ChildText(".elementor-heading-title") != "" {
			// Get all children as a array with css class  "elementor-heading-title"
			titles := []string{}
			e.ForEach(".elementor-heading-title", func(i int, el *colly.HTMLElement) {
				titles = append(titles, el.Text)
			})

			// Split the address which is the second element of the titles array by <br>
			data := strings.Split(titles[1], ")")
			comment := data[0] + ")"
			address := data[1]

			// Append location to slice of locations if not empty
			m[e.ChildText(".elementor-heading-title")] = LocationStruct{
				Name:    titles[0],
				Comment: comment,
				Address: address,
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	var err = c.Visit(urlAboutLocations)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Convert locations as array
	locations := make([]LocationStruct, 0, len(m))
	for _, v := range m {
		locations = append(locations, v)
	}

	return locations
}
