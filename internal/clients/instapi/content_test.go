package instapi

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/haski007/insta-bot/internal/bot/model"
	"github.com/haski007/pretty"
	"github.com/stretchr/testify/assert"
)

func TestApi_GetPostVideos(t *testing.T) {
	uploadVideoDate, err := time.Parse(time.RFC3339, "2023-06-12T14:38:52-07:00")
	if err != nil {
		t.Fatalf("error parsing upload date: %s", err)
	}

	exprLikesAndCommentsCount := regexp.MustCompile(`\b(\d+[\w]*)\slikes,\s(\d+[\w]*)\scomments`)

	tt := []struct {
		name    string
		postURL string
		want    *GetPostContentResponse
		wantErr bool
	}{
		{
			name:    "valid url",
			postURL: "https://www.instagram.com/p/CtZ9wdkg22s/",
			want: &GetPostContentResponse{
				ArticleBody: "отмечай миллионера",
				Author: model.Author{
					Identifier: model.Identifier{
						Value: "lekrav",
					},
					Name: "🕊 Добрый татарин — Ленар",
					URL:  "https://www.instagram.com/lekrav",
				},
				Video: []*model.Video{
					{
						Name:        "🕊 Добрый татарин — Ленар on Instagram: \"отмечай миллионера\"",
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/o1/v/t16/f1/m82/1C45CFA6B9721F4C8D80CBBA8D1EB6BA_video_dashinit.mp4?efg=eyJ2ZW5jb2RlX3RhZyI6InZ0c192b2RfdXJsZ2VuLjcyMC5jbGlwcyJ9&_nc_ht=scontent-waw1-1.cdninstagram.com&_nc_cat=104&vs=791105329137348_2062355818&_nc_vs=HBksFQIYT2lnX3hwdl9yZWVsc19wZXJtYW5lbnRfcHJvZC8xQzQ1Q0ZBNkI5NzIxRjRDOEQ4MENCQkE4RDFFQjZCQV92aWRlb19kYXNoaW5pdC5tcDQVAALIAQAVABgkR0RSZkdCVjB2aHNjelRrQ0FEcTB2Z2NwbXkxSWJxX0VBQUFGFQICyAEAKAAYABsBiAd1c2Vfb2lsATEVAAAm3riAodDq5D8VAigCQzMsF0A%2BogxJul41GBJkYXNoX2Jhc2VsaW5lXzFfdjERAHUAAA%3D%3D&ccb=9-4&oh=00_AfDvtvxRn3brcyn3IMHUscvFnA2EFTjKl3ixg9QA4H_gvw&oe=6494E377&_nc_sid=c07a80&_nc_rid=9ccca2d6be",
						Description: "84K likes, 591 comments - 🕊 Добрый татарин — Ленар (@lekrav) on Instagram: \"отмечай миллионера\"",
						Caption:     "отмечай миллионера",
						UploadDate:  uploadVideoDate,
					},
				},
				Image: []*model.Image{},
			},
			wantErr: false,
		},
		{
			name:    "1 photo post",
			postURL: "https://www.instagram.com/p/Cst95wuKEed/?utm_source=ig_web_copy_link&igshid=MzRlODBiNWFlZA==",
			want: &GetPostContentResponse{
				ArticleBody: "Викурили Амстердам\n\nКоржа захейтили",
				Author: model.Author{
					Identifier: model.Identifier{
						Value: "demianhaski",
					},
					Name: "Демян Пікалюк",
					URL:  "https://www.instagram.com/demianhaski",
				},
				Video: []*model.Video{},
				Image: []*model.Image{
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/349440129_923402708739024_6754856176829253906_n.jpg?stp=dst-jpg_s640x640&_nc_cat=109&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=-IfKLNmqiTwAX9IwtaW&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfDXP3RttEpXLUq-invGdCTKU1ncaL7fke4MO-c1JluDSA&oe=64996D89",
						Caption:     "Викурили Амстердам\n\nКоржа захейтили",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "many photos post",
			postURL: "https://www.instagram.com/p/CttHHqoMTqj/?utm_source=ig_web_copy_link&igshid=MzRlODBiNWFlZA==",
			want: &GetPostContentResponse{
				ArticleBody: "Немає в цій жорстокій війні, що триває на нашій землі, абсолютно нічого, що робило б саме нашу Україну джерелом війни або територією тих, хто хоче щось загарбати в іншого народу. Ми захищаємо життя своїх людей, свою свободу, свою незалежність, свої цінності. Мир для українських дітей!\n____\n\nThere is absolutely nothing in this brutal war that is taking place on our land that would make our Ukraine a source of war or the territory of those who want to take something away from another people. We are defending the lives of our people, our freedom, our independence, our values. Peace for Ukrainian children!\n\nPhoto: Enes Yıldırım, Ole_g_, Kostiantyn Liberov and Vlada Liberova, Virginie Nguyen Hoang, Donetsk separate brigade of the Territorial Defense Forces.",
				Author: model.Author{
					Identifier: model.Identifier{
						Value: "zelenskiy_official",
					},
					Name: "Володимир Зеленський",
					URL:  "https://www.instagram.com/zelenskiy_official",
				},
				Video: []*model.Video{},
				Image: []*model.Image{
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/355277172_1450428849102922_40158232980824626_n.jpg?stp=dst-jpg_s640x640&_nc_cat=1&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=Arjv6WUo9j4AX95lzNC&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfDku5zL8rStDQICYexmTuEoB4L_fIcmOAjzV5O9Owr0zw&oe=64978F2B",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/354957734_653568029545877_6091119027660760841_n.jpg?stp=dst-jpg_s640x640&_nc_cat=1&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=d5sHoxd-xuoAX9Xz4ql&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfAN52gGBka3W5l1mSjOWJtoQUCohGIki8llBUNJuVaZPg&oe=6498397C",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/355155953_640805061010869_6836313689671919572_n.jpg?stp=dst-jpg_s640x640&_nc_cat=104&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=Yt6AFo5DbOkAX_nPJck&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfBMYIZpzptUzd4phNTxGCO7uESKk0eany0OIvXgS2-qag&oe=64991110",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/354578136_277801724724934_6825190096586588656_n.jpg?stp=dst-jpg_s640x640&_nc_cat=105&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=Tf7iPpbpwlEAX8lg45F&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfCykUm15IDmT5oW9IqcTXO3y1SnOResQBtsJTeXXoGSOA&oe=64982C8A",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/355103430_214136961044858_7148117496017219853_n.jpg?stp=dst-jpg_s640x640&_nc_cat=109&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=ojAeIFH1EBIAX_Ctp6w&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfBTve8aCPHAbQz2utwtnaU72JuQ6c0jVHddxmh9ikSHdQ&oe=649946D8",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/355057245_5745546388880535_4085508577173903224_n.jpg?stp=dst-jpg_s640x640&_nc_cat=106&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=TAVyEOSWfl8AX-iQuhM&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfCVRQHoRs0hRVcimJwUy7ACn0EG5xlmXFlt_xf6VIcE0g&oe=64991A03",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/354628569_826229718517892_150645444225370515_n.jpg?stp=dst-jpg_s640x640&_nc_cat=102&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=o8Gz0yTDkhsAX9GzATc&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfBcaeyrnbZFPkTVX64yQkW1LW-V6jIY8031biedYEQUwA&oe=6498AA9B",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/355068643_1397585241029654_5485590344518731889_n.jpg?stp=dst-jpg_s640x640&_nc_cat=104&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=LW3IyoP3ULgAX8gPfWK&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfAtia0mR8viviC1V64D-QODXGS_BcXKlOGFYL2_2DWqLA&oe=649816CF",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/354521893_590041963268598_1079117761502576216_n.jpg?stp=dst-jpg_s640x640&_nc_cat=107&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=FH_5od7DdagAX_HCuaN&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfAUNpoFHS0hnqi88W1lom3dq_yYSbAzOCtz52YvywxI4Q&oe=64980118",
						Caption:     "",
					},
					{
						DownloadUrl: "https://scontent-waw1-1.cdninstagram.com/v/t51.29350-15/355075779_654695166077015_2904956912314294556_n.jpg?stp=dst-jpg_s640x640&_nc_cat=103&ccb=1-7&_nc_sid=8ae9d6&_nc_ohc=XeGHaRqO1hsAX8qNUsY&_nc_ht=scontent-waw1-1.cdninstagram.com&oh=00_AfC-I75bF-xyI8jhj326n23ekF2LTThASnBgHWrK44BLhA&oe=6498226E",
						Caption:     "",
					},
				},
			},
			wantErr: false,
		},

		//{
		//	name:    "story with photo",
		//	postURL: "https://www.instagram.com/stories/yulia_pavytska/3130184198920366497/",
		//	want:    "",
		//	wantErr: false,
		//},
		//{
		//	name:    "story with shared post",
		//	postURL: "https://www.instagram.com/stories/tatiianna_13/3130000275199360909/",
		//	want:    "",
		//	wantErr: false,
		//},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rcv := &Api{}
			got, err := rcv.GetPostContent(tc.postURL)
			if (err != nil) != tc.wantErr {
				assert.NoError(t, err, "getReelsDownloadURL() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			for _, v := range got.Video {
				v.DownloadUrl = v.DownloadUrl[0 : len(v.DownloadUrl)-10]
				v.Description = exprLikesAndCommentsCount.ReplaceAllString(v.Description, "0 likes, 0 comments")
			}
			for _, v := range tc.want.Video {
				v.DownloadUrl = v.DownloadUrl[0 : len(v.DownloadUrl)-10]
				v.Description = exprLikesAndCommentsCount.ReplaceAllString(v.Description, "0 likes, 0 comments")
			}

			for _, i := range got.Image {
				i.DownloadUrl = i.DownloadUrl[0 : len(i.DownloadUrl)-10]
			}
			for _, i := range tc.want.Image {
				i.DownloadUrl = i.DownloadUrl[0 : len(i.DownloadUrl)-10]
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Expected and result do not match. Expected: %s, got: %s", pretty.String(tc.want), pretty.String(got))
			}
		})
	}
}
