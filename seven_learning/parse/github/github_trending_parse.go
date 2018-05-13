package github

import (
	"fmt"

	"go-example-for-live/seven_learning/infra"

	"github.com/PuerkitoBio/goquery"
)

func ParseForGithub(document *goquery.Document) {

	document.Find("div.explore-content ol.repo-list li").Each(func(i int, selection *goquery.Selection) {
		RespName, _ := infra.HandleCommon(selection.Find("div h3 a").Text())
		URL, _ := infra.HandlerURL(selection.Find("div").Eq(0).Find("h3 a").AttrOr("href", "None"))
		Description, _ := infra.HandleCommon(selection.Find("div").Eq(2).Find("p").Text())
		Stars, _ := infra.HandleCommon(selection.Find("div").Eq(3).Find("a").Eq(0).Text())
		Fork, _ := infra.HandleCommon(selection.Find("div").Eq(3).Find("a").Eq(1).Text())
		TodayStars, _ := infra.HandleCommon(selection.Find("div").Eq(3).Find("span").Eq(1).Text())

		fmt.Println(RespName, URL, Description, Stars, Fork, TodayStars)
	})

}

func ParseForDevelopers(document *goquery.Document) {

	document.Find("div.explore-content ol li").Each(func(i int, selection *goquery.Selection) {
		DevName, _ := infra.HandleCommon(selection.Find("li div div").Eq(1).Find("h2 a").Text())
		Description, _ := infra.HandleCommon(selection.Find("li div div").Eq(1).Find("a span").Text())
		URL, _ := infra.HandleCommon(selection.Find("li div div").Eq(1).Find("h2 a").AttrOr("href", "None"))

		fmt.Println(DevName, Description, URL)
	})
}
