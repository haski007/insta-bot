package publisher

import (
	"context"
	"fmt"
	"strconv"
)

func (rcv *Publisher) GetMediaPkFromUrl(ctx context.Context, url string) (int, error) {
	return rcv.getMediaPkFromURL(ctx, url)
}

func (rcv *Publisher) getMediaPkFromURL(ctx context.Context, url string) (int, error) {
	body, rsp, err := rcv.apiCli.MediaApi.MediaPkFromUrlMediaPkFromUrlGet(ctx).Url(url).Execute()

	if err != nil {
		return 0, fmt.Errorf("get req err: %w", err)
	} else if rsp.StatusCode != 200 {
		return 0, fmt.Errorf("get req status: %s", rsp.Status)
	}

	if str, ok := body.(string); ok {
		pk, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot parse body [%s]] to int err: %w", str, err)
		}
		return int(pk), nil
	}

	return 0, fmt.Errorf("can't cast body to string, body: %+v", body)
}

func (rcv *Publisher) GetStoryPkFromUrl(ctx context.Context, url string) (string, error) {
	return rcv.getStoryPkFromURL(ctx, url)
}

func (rcv *Publisher) getStoryPkFromURL(ctx context.Context, url string) (string, error) {
	body, rsp, err := rcv.apiCli.StoryApi.StoryPkFromUrlStoryPkFromUrlGet(ctx).Url(url).Execute()

	if err != nil {
		return "", fmt.Errorf("get req err: %w", err)
	} else if rsp.StatusCode != 200 {
		return "", fmt.Errorf("get req status: %s", rsp.Status)
	}

	if pk, ok := body.(string); ok {
		return pk, nil
	}

	return "", fmt.Errorf("can't cast body to string, body: %+v", body)
}
