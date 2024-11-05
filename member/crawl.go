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
			// Append member to slice of members if not empty
			m[e.ChildText(".lae-title")+"."+e.ChildText(".lae-team-member-position")] = MemberStruct{
				Hash:     hashMember(MemberStruct{Name: e.ChildText(".lae-title"), Position: e.ChildText(".lae-team-member-position"), Avatar: e.ChildAttr(".lae-image", "src")}),
				Name:     e.ChildText(".lae-title"),
				Position: e.ChildText(".lae-team-member-position"),
				Avatar:   e.ChildAttr(".lae-image", "src"),
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
