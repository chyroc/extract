package fetch

import (
	"context"
	"io/ioutil"
	"net/http"
)

var generalFetchIns = generalFetch{}

type generalFetch struct{}

func (generalFetch) Fetch(ctx context.Context, url string) (string, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

var httpClient = &http.Client{}
