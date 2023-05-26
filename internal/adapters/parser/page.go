package parser

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

type PageFetcher struct {
	url  string
	pSub string
}

func NewPageFetcher(url string, pSub string) *PageFetcher {
	return &PageFetcher{url: url, pSub: pSub}
}

// FetchTablePage implements application.PageFetcher
func (pf *PageFetcher) FetchTablePage(ctx context.Context, group string, date string) (*html.Node, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Make a post request to the page
	response, err := http.PostForm(pf.url, url.Values{
		"p_sub":        {pf.pSub},
		"p_group_name": {group},
		"p_date":       {date}},
	)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Convert the response body to utf-8 charset
	reader, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	// Parse the response body into a html node
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
