package extract

import (
	"context"
)

type Meta struct {
	Title   string
	Content string
}

func GetMeta(ctx context.Context, url string) (*Meta, error) {
	panic("impl")
}
