package member

import (
	"log"

	"encore.dev/metrics"
	"github.com/gocolly/colly"
)

type Labels struct {
	Success bool
}

var CrawlRequests = metrics.NewCounterGroup[Labels, uint64]("crawl_requests", metrics.CounterConfig{})

const urlMembers = "https://thinkport.digital/thinkport-cloud-experten-uber-uns/"

// Returns an array of Members crawled from a website
func getMembers() []MemberStruct {
	var success bool

	c := colly.NewCollector()
	memberCount := 0

	m := make(map[string]MemberStruct)

	// Find all divs with class "lae-team-member " and print them
	c.OnHTML(".lae-team-member", func(e *colly.HTMLElement) {
		memberCount++

		if e.ChildText(".lae-title") != "" {
			mail := e.ChildAttr("a[href^=mailto]", "href")
			if len(mail) > 0 {
				mail = mail[7:]
			}
			// Append member to slice of members if not empty
			m[e.ChildText(".lae-title")] = MemberStruct{
				Name:     e.ChildText(".lae-title"),
				Surname:  extractSurnameFromEmail(mail),
				Position: e.ChildText(".lae-team-member-position"),
				Details:  e.ChildText(".lae-team-member-details"),
				Email:    mail,
				Linkedin: e.ChildAttr(".lae-linkedin", "href"),
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	err := c.Visit(urlMembers)

	if err != nil {
		log.Println("Error: ", err)
		success = false
	} else {
		log.Println("Members crawled: ", memberCount)
	}

	CrawlRequests.With(Labels{Success: success}).Increment()

	// Remove empty members
	for k, v := range m {
		if v.Name == "" {
			delete(m, k)
		}
	}

	// Convert members as array
	members := make([]MemberStruct, 0, len(m))
	for _, v := range m {
		members = append(members, v)
	}

	return members

}
