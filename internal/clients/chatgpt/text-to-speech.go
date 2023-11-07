package chatgpt

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	VoiceBlackMan = "onyx"
	VoiceDefault  = VoiceBlackMan

	VoicePizda    = "nova"
	VoiceAlloy    = "alloy"
	VoiceEcho     = "echo"
	VoiceFable    = "fable"
	VoiceShrimmer = "shimmer"
)

var mapVoice = map[string]struct{}{
	VoiceBlackMan: {},
	VoicePizda:    {},
	VoiceAlloy:    {},
	VoiceEcho:     {},
	VoiceFable:    {},
	VoiceShrimmer: {},
}

func (srv *Service) TextToSpeech(_ context.Context, voice, message string) ([]byte, error) {
	url := ApiBaseURL + TTSPath

	modelVoice := VoiceDefault
	if _, ok := mapVoice[voice]; ok {
		modelVoice = voice
	}

	payload := strings.NewReader(fmt.Sprintf("{"+
		"\"model\": \"%s\","+
		"\"input\": \"%s\","+
		"\"voice\": \"%s\""+
		"}", GPTTTSModel, message, modelVoice))

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, fmt.Errorf("new request err: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", srv.apiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request err: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read body err: %w", err)
	}

	return body, nil
}
