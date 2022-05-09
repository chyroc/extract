package parse

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	as := assert.New(t)
	title, content, err := testPage("https://www.v2ex.com/t/851390")
	as.Nil(err)
	as.Equal("求推荐一个离线下载的服务 - V2EX", title)
	fmt.Println(content)
}

func testPage(uri string) (string, string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", "", err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	urlParsed, err := url.Parse(uri)
	if err != nil {
		return "", "", err
	}

	return Parse(urlParsed, string(bs))
}
