package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	wsj = wikiSearchJSON{Query: sq}
	sr  = searchResults{s}
	srn = searchResults{}
	sq  = SearchQuery{sr}
	s   = Search{Title: "result"}
	wej = wikiExtractJSON{Query: eq}
	eq  = ExtractQuery{Pages: ps}
	ps  = Pages{"1": p}
	psn = Pages{}
	p   = Page{Extract: "extract"}
)

func TestSearch(t *testing.T) {
	cases := []struct {
		name string
		obj  fmt.Stringer
	}{
		{
			name: "Search JSON",
			obj:  wsj,
		},
		{
			name: "Search results",
			obj:  sr,
		},
		{
			name: "Search queries",
			obj:  sq,
		},
		{
			name: "Search result",
			obj:  s,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.obj.String(), "result")
		})
	}
}

func TestExtract(t *testing.T) {
	cases := []struct {
		name string
		obj  fmt.Stringer
	}{
		{
			name: "Extract JSON",
			obj:  wej,
		},
		{
			name: "Extract query results",
			obj:  eq,
		},
		{
			name: "Pages",
			obj:  ps,
		},
		{
			name: "Page",
			obj:  p,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.obj.String(), "extract")
		})
	}
}

func TestSearchNoResult(t *testing.T) {
	assert.Equal(t, srn.String(), "")
}

func TestPagesNoResult(t *testing.T) {
	assert.Equal(t, psn.String(), "No search results found.")
}
