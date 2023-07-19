package instapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/haski007/insta-bot/internal/bot/model"
	"github.com/sirupsen/logrus"
)

type GetPostContentResponse []*GetPostContentResponseObj

type GetPostContentResponseObj struct {
	ArticleBody string `json:"articleBody"`

	Author model.Author   `json:"author"`
	Video  []*model.Video `json:"video"`
	Image  []*model.Image `json:"image"`
}

func (rcv *Api) GetPostContent(url string) (*GetPostContentResponseObj, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting reels download URL: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		return nil, err
	}

	jsonPart := doc.Text()
	fixedCurve := getSubstrBefore(jsonPart, "}:")
	if fixedCurve != "" {
		jsonPart = fixedCurve + "}"
	}
	fixedSquare := getSubstrBefore(jsonPart, "]:")
	if fixedSquare != "" {
		jsonPart = fixedSquare + "]"
	}

	if err := os.WriteFile("test.json", []byte(jsonPart), 0777); err != nil {
		logrus.Fatalf("error writing test.json: %s", err)
	}

	var response GetPostContentResponse

	if err := json.Unmarshal([]byte(jsonPart), &response); err != nil {
		return nil, fmt.Errorf("error decoding reels download URL: %w", err)
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("empty response")
	}

	return response[0], nil
}
