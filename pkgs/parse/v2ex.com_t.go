package parse

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	register(&v2exComT{})
}

type v2exComT struct{}

func (r v2exComT) SupportHost() []string {
	return []string{
		"v2ex.com",
		"www.v2ex.com",
	}
}

func (r v2exComT) Parse(url string, html string) (string, string, error) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	sel := doc.Find("#Main")
	// sel.Find(".highlight__panel").Remove()
	content, err := sel.Html()
	if err != nil {
		return "", "", err
	}
	title := strings.TrimSpace(doc.Find("title").Text())

	return title, content, nil
}
