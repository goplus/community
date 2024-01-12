package core

import "context"

// PutMedia uploads media.
func (p *Community) PutMedia(ctx context.Context, uid string, media []byte) (id string, err error) {
	return
}

func (p *Community) DeleteMedia(ctx context.Context, uid, id string) (err error) {
	return
}

func (p *Community) MediaURL(id string) (url string) {
	return
}
