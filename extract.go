package extract

import (
	"context"

	"github.com/chyroc/extract/pkgs"
)

type Meta struct {
	Title   string
	Content string
}

func GetMeta(ctx context.Context, url string) (*Meta, error) {
	title, content, err := pkgs.Run(ctx, url)
	if err != nil {
		return nil, err
	}
	return &Meta{
		Title:   title,
		Content: content,
	}, nil
}
