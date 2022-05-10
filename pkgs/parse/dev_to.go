package parse

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	register(&devTo{})
}

type devTo struct{}

func (r devTo) SupportHost() []string {
	return []string{
		"dev.to",
	}
}

func (r devTo) Parse(url string, html string) (string, string, error) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	sel := doc.Find(".crayons-article__body")
	sel.Find(".highlight__panel").Remove()
	content, err := sel.Html()
	if err != nil {
		return "", "", err
	}
	title := strings.TrimSpace(doc.Find("title").Text())

	return title, content, nil
}
