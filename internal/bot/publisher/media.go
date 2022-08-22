package publisher

import (
	"context"
	"fmt"

	"github.com/haski007/insta-bot/internal/clients/instapi"
)

type MediaType int

const (
	Photo MediaType = 1
	Video MediaType = 2
	Album MediaType = 8
)

func (rcv *Publisher) GetMediaInfoFromURL(ctx context.Context, url string) (*instapi.Media, error) {
	pk, err := rcv.getMediaPkFromURL(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get pk from url err: %w", err)
	}

	return rcv.GetMediaInfoFromPk(ctx, pk)
}

func (rcv *Publisher) GetMediaInfoFromPk(ctx context.Context, pk int) (*instapi.Media, error) {
	return rcv.getMediaInfoFromPk(ctx, pk)
}

func (rcv *Publisher) getMediaInfoFromPk(ctx context.Context, pk int) (*instapi.Media, error) {
	media, rsp, err := rcv.apiCli.MediaApi.MediaInfoMediaInfoPost(ctx).
		Sessionid(rcv.authToken).
		Pk(pk).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("MediaInfoMediaInfoPost err: %w", err)
	} else if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("MediaInfoMediaInfoPost status: %s", rsp.Status)
	}

	return media, nil
}
