package member

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Returns an array of Members crawled from a website
func getMembers() []MemberStruct {

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
				Position: e.ChildText(".lae-position"),
				Email:    mail,
				Linkedin: e.ChildAttr("a[href^=https://www.linkedin.com]", "href"),
			}

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://thinkport.digital/thinkport-cloud-experten-uber-uns/")

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
