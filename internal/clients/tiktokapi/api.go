package tiktokapi

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/haski007/insta-bot/pkg/text"

	"github.com/PuerkitoBio/goquery"
	"github.com/haski007/insta-bot/internal/bot"
	"github.com/haski007/insta-bot/internal/bot/model"
)

const (
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"
)

type TikTokClient struct{}

func New() *TikTokClient {
	return &TikTokClient{}
}

func (rcv *TikTokClient) GetVideoDataFromUrl(url string) (bot.TikTokVideo, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new req err: %w", err)
	}

	req.Header.Add("user-agent", userAgent)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client do err: %w", err)
	}
	defer rsp.Body.Close()

	rspData, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body err: %s", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(rspData)))
	if err != nil {
		return nil, fmt.Errorf("new htmldoc from reader err: %w", err)
	}

	title := doc.Find("title").Text()
	// todo: check if not empty
	arr := regexp.MustCompile(`"downloadAddr"\s*:\s*"(https.*?)",`).FindStringSubmatch(string(rspData))
	if len(arr) < 2 {
		return nil, fmt.Errorf("download url not found in gey response")
	}

	//fmt.Println(pretty.String(arr))
	//for i, v := range arr {
	//	fmt.Printf("||||||||||||| #%d: %s\n", i, text.DecodeUrl(v))
	//}
	downloadUrl := text.DecodeUrl(arr[1])

	authorArr := regexp.MustCompile(`"author"\s*:\s*"(.*?)",`).FindStringSubmatch(string(rspData))

	var author string
	for i, a := range authorArr {
		if a != author && i != 0 {
			author = a
		}
	}

	return &model.Video{
		Title:        title,
		DownloadUrl:  downloadUrl,
		Author:       author,
		OriginalLink: url,
	}, nil
}
