package jianshu

import (
	"github.com/PuerkitoBio/goquery"
)

func InfoParse(response *goquery.Document) []string {
	// #note-28947518 > div > a

	var title []string

	response.Find("div#list-container ul.note-list li").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("div.content a.title").Text()
		//fmt.Println(text)
		title = append(title, text)
	})
	return title
}
