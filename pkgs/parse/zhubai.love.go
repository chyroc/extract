package parse

import (
	"context"
	"net/url"
	"strings"

	"github.com/chromedp/chromedp"
)

func init() {
	register(zhubaiLove{})
}

type zhubaiLove struct{}

func (r zhubaiLove) MatchURL(url *url.URL) bool {
	return strings.HasSuffix(url.Host, ".zhubai.love")
}

func (r zhubaiLove) Parse(url string, html string) (string, string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var content string
	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML(`div[class^="PostPage_content_"]`, &content, chromedp.NodeVisible),
		chromedp.Text(`h1[class^="PostPage_title_"]`, &title, chromedp.NodeVisible),
	)
	if err != nil {
		return "", "", err
	}

	return title, strings.TrimSpace(content), nil
}
