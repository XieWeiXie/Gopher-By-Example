package v2ex

import (
	"github.com/PuerkitoBio/goquery"
)

func JobParse(response *goquery.Document) []string {

	var title []string

	response.Find("div#Main div.box div.cell.item").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("table tbody tr td").Eq(2).Find("span a").Text()
		//fmt.Println(text)
		title = append(title, text)
	})
	return title
}
