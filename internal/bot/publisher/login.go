package publisher

import (
	"context"
	"fmt"
)

func (rcv *Publisher) Login(ctx context.Context) error {
	body, rsp, err := rcv.apiCli.AuthApi.AuthLoginAuthLoginPost(ctx).
		Username(rcv.username).
		Password(rcv.password).
		Execute()
	if err != nil {
		return fmt.Errorf("auth post req err: %w", err)
	}

	if rsp.StatusCode != 200 {
		return fmt.Errorf("auth post resp status code: %s", rsp.Status)
	}

	if token, ok := body.(string); ok {
		rcv.authToken = token
	} else {
		return fmt.Errorf("no token in response: %+v", body)
	}

	return nil
}
