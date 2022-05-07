package fetch

import (
	"context"
	"net/url"
)

type Fetcher interface {
	Fetch(ctx context.Context, url string) (string, error)
}

type FetcherByHost interface {
	SupportHost() []string
}

func Get(url *url.URL) (Fetcher, error) {
	fetcher, ok := fetchByHosts[url.Host]
	if ok {
		return fetcher, nil
	}
	return generalFetchIns, nil
}

func Fetch(ctx context.Context, url *url.URL) (string, error) {
	fetcher, err := Get(url)
	if err != nil {
		return "", err
	}
	return fetcher.Fetch(ctx, url.String())
}

var fetchByHosts = map[string]Fetcher{}
