package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
)

func getVideoURL(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	os.WriteFile("test.json", []byte(doc.Text()), 0777)

	var videoURL string
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if property, _ := s.Attr("property"); property == "og:video" {
			videoURL, _ = s.Attr("content")
		}
	})

	return videoURL, nil
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	url := "https://www.instagram.com/p/CtZ9wdkg22s" // Example reel URL
	videoURL, err := getVideoURL(url)
	if err != nil {
		logrus.WithError(err).Fatalln("error getting video URL")
	}

	videoURL = `https:\/\/scontent-waw1-1.cdninstagram.com\/o1\/v\/t16\/f1\/m82\/1C45CFA6B9721F4C8D80CBBA8D1EB6BA_video_dashinit.mp4?efg=eyJ2ZW5jb2RlX3RhZyI6InZ0c192b2RfdXJsZ2VuLjcyMC5jbGlwcyJ9&_nc_ht=scontent-waw1-1.cdninstagram.com&_nc_cat=104&vs=791105329137348_2062355818&_nc_vs=HBksFQIYT2lnX3hwdl9yZWVsc19wZXJtYW5lbnRfcHJvZC8xQzQ1Q0ZBNkI5NzIxRjRDOEQ4MENCQkE4RDFFQjZCQV92aWRlb19kYXNoaW5pdC5tcDQVAALIAQAVABgkR0RSZkdCVjB2aHNjelRrQ0FEcTB2Z2NwbXkxSWJxX0VBQUFGFQICyAEAKAAYABsBiAd1c2Vfb2lsATEVAAAm3riAodDq5D8VAigCQzMsF0A\u00252BogxJul41GBJkYXNoX2Jhc2VsaW5lXzFfdjERAHUAAA\u00253D\u00253D&ccb=9-4&oh=00_AfDvtvxRn3brcyn3IMHUscvFnA2EFTjKl3ixg9QA4H_gvw&oe=6494E377&_nc_sid=c07a80&_nc_rid=9ccca2d6be`

	step1 := strings.ReplaceAll(videoURL, `\/`, `/`)

	// Now handle the unicode sequences
	step2, err := strconv.Unquote(`"` + step1 + `"`)
	if err != nil {
		logrus.WithError(err).Fatalln("error unquoting video URL")
	}

	logrus.Println(step2)
	err = downloadFile("reel.mp4", step2)
	if err != nil {
		logrus.WithError(err).Fatalln("error downloading file")
	}

	fmt.Println("Done!")
}
