package publisher

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (rcv *Publisher) Login(ctx context.Context) error {
	logrus.Println("verify code:", rcv.verificationCode)
	body, rsp, err := rcv.apiCli.AuthApi.AuthLoginAuthLoginPost(ctx).
		Username(rcv.username).
		Password(rcv.password).
		VerificationCode(rcv.verificationCode).
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
