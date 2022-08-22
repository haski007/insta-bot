package publisher

import (
	"context"
	"fmt"

	"github.com/haski007/insta-bot/internal/clients/instapi"
)

func (rcv *Publisher) GetStoryInfoFromUrl(ctx context.Context, url string) (*instapi.Story, error) {
	pk, err := rcv.getStoryPkFromURL(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get pk from url err: %w", err)
	}

	return rcv.GetStoryInfoFromPk(ctx, pk)
}

func (rcv *Publisher) GetStoryInfoFromPk(ctx context.Context, pk string) (*instapi.Story, error) {
	return rcv.getStoryInfoFromPk(ctx, pk)
}

func (rcv *Publisher) getStoryInfoFromPk(ctx context.Context, pk string) (*instapi.Story, error) {
	story, rsp, err := rcv.apiCli.StoryApi.StoryInfoStoryInfoPost(ctx).
		Sessionid(rcv.authToken).
		StoryPk(pk).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("StoryInfoStoryInfoPost err: %w", err)
	} else if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("StoryInfoStoryInfoPost status: %s", rsp.Status)
	}

	return story, nil
}
