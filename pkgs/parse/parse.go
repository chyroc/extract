package parse

import (
	"fmt"
	"net/url"
)

type Parser interface {
	Parse(url string, html string) (string, string, error)
}

type HostMatcher interface {
	SupportHosts() []string
}

type URLMatcher interface {
	MatchURL(url *url.URL) bool
}

func Parse(url2 string, url *url.URL, html string) (string, string, error) {
	parser, ok := parsersByHost[url.Host]
	if ok {
		return parser.Parse(url2, html)
	}
	for _, v := range parserByURL {
		if v.match(url) {
			return v.parser.Parse(url2, html)
		}
	}
	return "", "", fmt.Errorf("no parser for %s", url.Host)
}

var (
	parsersByHost = make(map[string]Parser)
	parserByURL   = []x{}
)

func register(parser Parser) {
	if m, ok := parser.(HostMatcher); ok {
		for _, v := range m.SupportHosts() {
			if _, ok := parsersByHost[v]; ok {
				panic("duplicate register host:" + v)
			}
			parsersByHost[v] = parser
		}
	}

	if m, ok := parser.(URLMatcher); ok {
		parserByURL = append(parserByURL, x{parser: parser, match: m.MatchURL})
	}
}

type x struct {
	parser Parser
	match  func(*url.URL) bool
}
