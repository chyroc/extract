package parse

import (
	"fmt"
	"net/url"
)

type Parser interface {
	Host() []string
	Parse(url *url.URL, html string) (string, string, error)
}

func Parse(url *url.URL, html string) (string, string, error) {
	parser, ok := parsers[url.Host]
	if !ok {
		return "", "", fmt.Errorf("no parser for %s", url.Host)
	}
	return parser.Parse(url, html)
}

var parsers = make(map[string]Parser)

func register(parser Parser) {
	for _, v := range parser.Host() {
		if _, ok := parsers[v]; ok {
			panic("duplicate register host:" + v)
		}
		parsers[v] = parser
	}
}
