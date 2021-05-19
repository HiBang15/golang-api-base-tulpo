package models

type Activity struct {
	Id       int64  `json:"id"`
	Url      string `json:"url"`
	Method   string `json:"method"`
	UrlRegex string `json:"url_regex"`
}

type CreateActivity struct {
	Url      string `json:"url"`
	Method   string `json:"method"`
	UrlRegex string `json:"url_regex"`
}
