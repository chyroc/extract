package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/chyroc/extract/pkgs"
	"github.com/codesoap/rss2"
	"github.com/google/uuid"
)

func main() {
	title, content, err := pkgs.Run(context.Background(), os.Args[1])
	if err != nil {
		panic(err)
	}
	generateTestRssXml("/tmp/rss.xml", title, content)
}

func generateTestRssXml(xmlFile string, title, content string) {
	err := writeRss(xmlFile, &Feed{
		Title:       "测试 Rss",
		Link:        "http://127.0.0.1:8080",
		Description: "测试 Rss",
		Items: []*Item{
			{
				Title:       title,
				Link:        "http://127.0.0.1:8080",
				Description: content,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func writeRss(xmlFile string, feed *Feed) error {
	ch, err := NewChannel(feed.Title, feed.Link, feed.Description)
	if err != nil {
		return err
	}
	for _, v := range feed.Items {
		item, err := NewItem(v.Title, v.Description)
		if err != nil {
			return err
		}
		item.Link = v.Link
		item.Author = v.Author
		if !(v.PubDate.IsZero() || v.PubDate.Year() == 0 || v.PubDate.Year() == 1) {
			item.PubDate = &rss2.RSSTime{Time: v.PubDate}
		}
		item.GUID, _ = rss2.NewGUID(uuid.New().String())
		ch.Items = append(ch.Items, item)
	}
	rss := rss2.NewRSS(ch)
	bs, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(xmlFile, bs, 0o666)
}

type Feed struct {
	Title       string
	Link        string
	Description string
	Items       []*Item
}

type Item struct {
	Title       string
	Link        string
	Description string
	Author      string
	PubDate     time.Time
}

func NewChannel(title, link, description string) (*rss2.Channel, error) {
	if len(title) == 0 || len(link) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewChannel()`)
	}
	return &rss2.Channel{
		XMLName:     xml.Name{Local: `channel`},
		Title:       title,
		Link:        link,
		Description: description,
	}, nil
}

func NewItem(title, description string) (*rss2.Item, error) {
	if len(title) == 0 {
		return nil, fmt.Errorf(`cannot create item with empty title and description`)
	}
	return &rss2.Item{
		XMLName:     xml.Name{Local: `item`},
		Title:       title,
		Description: description,
	}, nil
}
