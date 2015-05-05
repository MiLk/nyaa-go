package nyaa

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

type Entry struct {
	Name    string
	Link    string
	Torrent string
}

type API struct {
	endpoint string
	request  *gorequest.SuperAgent
}

func (api *API) Search(title string) (entries []Entry, errs []error) {
	_url := api.endpoint + "?page=search&cats=1_37&filter=0&term=" + url.QueryEscape(title) + "&sort=2"
	_, body, errs := api.request.Get(_url).End()
	if len(errs) != 0 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		errs = append(errs, err)
		return
	}

	doc.Find(".tlist tbody tr.tlistrow").Each(func(i int, s *goquery.Selection) {
		var entry Entry
		entry.Link, _ = s.Find("td.tlistname a").Attr("href")
		entry.Name = s.Find("td.tlistname a").Text()
		entry.Torrent, _ = s.Find("td.tlistdownload a").Attr("href")
		entries = append(entries, entry)
	})

	return
}

func NewAPI() *API {
	api := new(API)
	api.endpoint = "http://www.nyaa.se/"
	api.request = gorequest.New()
	return api
}
