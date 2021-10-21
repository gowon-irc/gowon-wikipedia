package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	wikipediaAPISearchURL = "https://en.wikipedia.org/w/api.php?" +
		"format=json&action=query&list=search&srlimit=1&" +
		"srprop=timestamp&srwhat=text&srsearch=%s"
	wikipediaAPIExtractURL = "https://en.wikipedia.org/w/api.php?" +
		"format=json&action=query&prop=extracts&exintro&explaintext&" +
		"exchars=400&redirects&titles=%s"
)

type wikiSearchJSON struct {
	Query SearchQuery `json:"query"`
}

func (wsj wikiSearchJSON) String() string {
	return wsj.Query.String()
}

// SearchQuery represents the search results from the wikipedia api
type SearchQuery struct {
	Search searchResults `json:"search"`
}

func (sq SearchQuery) String() string {
	return sq.Search.String()
}

type searchResults []Search

func (sr searchResults) String() string {
	if len(sr) == 0 {
		return ""
	}
	return sr[0].Title
}

// Search represents a single search result from the wikipedia api
type Search struct {
	Title string `json:"title"`
}

func (s Search) String() string {
	return s.Title
}

type wikiExtractJSON struct {
	Query ExtractQuery `json:"query"`
}

func (wej wikiExtractJSON) String() string {
	return wej.Query.String()
}

// ExtractQuery represents the query object from the wikipedia api
type ExtractQuery struct {
	Pages Pages `json:"pages"`
}

func (q ExtractQuery) String() string {
	return q.Pages.String()
}

// Pages represents the page results from the wikipedia api
type Pages map[string]Page

func (ps Pages) String() string {
	for _, v := range ps {
		return v.String()
	}
	return "No search results found."
}

// Page represents a single page result from wikipedia api
type Page struct {
	Extract string `json:"extract"`
}

func (p Page) String() string {
	return p.Extract
}

func wiki(args string) (msg string, err error) {
	urlArgs := url.QueryEscape(args)
	searchURL := fmt.Sprintf(wikipediaAPISearchURL, urlArgs)

	res, err := http.Get(searchURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	sj := &wikiSearchJSON{}
	err = json.Unmarshal(body, &sj)

	urlTitle := url.QueryEscape(sj.String())
	extractURL := fmt.Sprintf(wikipediaAPIExtractURL, urlTitle)

	res, err = http.Get(extractURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	ej := &wikiExtractJSON{}
	err = json.Unmarshal(body, &ej)

	if err != nil {
		return "", err
	}

	return ej.String(), nil
}

// func wiki(args string) (msg string, err error) {
//     args := command.RawArgs
//     urlArgs := url.QueryEscape(args)

//     searchURL := fmt.Sprintf(wikipediaAPISearchURL, urlArgs)
//     sj := &wikiSearchJSON{}
//     err = web.GetJSON(searchURL, sj)

//     if err != nil {
//         return "", err
//     }

//     urlTitle := url.QueryEscape(sj.String())

//     extractURL := fmt.Sprintf(wikipediaAPIExtractURL, urlTitle)
//     ej := &wikiExtractJSON{}
//     err = web.GetJSON(extractURL, ej)

//     if err != nil {
//         return "", err
//     }

//     return ej.String(), nil
// }
