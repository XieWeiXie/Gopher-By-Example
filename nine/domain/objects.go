package domain

import "github.com/PuerkitoBio/goquery"

type Request struct {
	Url       string
	ParseFunc func(*goquery.Document) Contents
}

type Content struct {
	URL          string
	Title        string
	Tag          string
	Reporter     string
	ReporterHome string
	Followers    int
	Answers      int
	Answerer     string
	See          int
	Passage      string
}

type Contents []Content
