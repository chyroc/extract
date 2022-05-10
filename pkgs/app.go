package pkgs

import (
	"context"
	"net/url"

	"github.com/chyroc/extract/pkgs/fetch"
	"github.com/chyroc/extract/pkgs/parse"
)

func Run(ctx context.Context, uri string) (string, string, error) {
	// 解析 url
	urlParsed, err := url.Parse(uri)
	if err != nil {
		return "", "", err
	}

	// 获取内容
	html, err := fetch.Fetch(ctx, urlParsed)
	if err != nil {
		return "", "", err
	}

	// 解析
	return parse.Parse(uri, urlParsed, html)
}
