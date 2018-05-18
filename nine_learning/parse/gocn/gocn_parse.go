package gocn

import (
	"errors"

	"go-example-for-live/nine_learning/infra"

	"go-example-for-live/nine_learning/download"

	"go-example-for-live/nine_learning/domain"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorParse = errors.New("parse error")
)

func TitleParse(document *goquery.Document) domain.Contents {

	var (
		AllContents domain.Contents
		OneContent  domain.Content
	)

	document.Find("div.aw-common-list div.aw-item").Each(func(i int, selection *goquery.Selection) {
		var (
			user           string
			userHome       string
			url            string
			passageContent string
			title          string
			tag            string
			lastAnswer     string
			one            int
			two            int
			three          int
		)
		user = infra.StringSplit(selection.Find("a").Eq(0).AttrOr("href", "None"))
		userHome = selection.Find("a").Eq(0).AttrOr("href", "None")
		url = selection.Find("div h4 a").AttrOr("href", "None")
		doc, _ := download.Download(url)
		passageContent, _ = PassageParse(doc)
		title = infra.StringTrim(selection.Find("div h4").Text())
		tag = selection.Find("p a").Eq(0).Text()
		lastAnswer = selection.Find("p a").Eq(1).Text()
		one, two, three = infra.StringSplitByDot(selection.Find("p span").Eq(0).Text())
		//fmt.Println(user, userHome, url, title, tag, lastAnswer, one, two, three, passageContent)

		OneContent = domain.Content{
			URL:          url,
			Title:        title,
			Tag:          tag,
			Reporter:     user,
			ReporterHome: userHome,
			Followers:    one,
			Answers:      two,
			Answerer:     lastAnswer,
			See:          three,
			Passage:      passageContent,
		}
		//fmt.Println(OneContent)
		AllContents = append(AllContents, OneContent)

	})
	return AllContents

}

func PassageParse(document *goquery.Document) (string, error) {
	var content string
	content = infra.StringTrim(document.Find("div.content.markitup-box").Text())
	return infra.StringsReplace(content), nil
}
