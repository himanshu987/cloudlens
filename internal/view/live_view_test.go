package view

import (
	"testing"

	"github.com/sahilm/fuzzy"
	"github.com/stretchr/testify/assert"
)

func TestHighlightMatchesWrapsContiguousMatch(t *testing.T) {
	lines := []string{`"Name": "cloudlens-demo-web"`}
	matches := fuzzy.Matches{
		{Index: 0, MatchedIndexes: []int{24, 25, 26}},
	}

	got := highlightMatches(lines, matches)

	assert.Equal(t, []string{`"Name": "cloudlens-demo-` + `["search_0"]web[""]` + `"`}, got)
}

func TestHighlightMatchesWrapsFullSpanForNonContiguousFuzzyMatch(t *testing.T) {
	line := `cloudlens-demo-web`
	lines := []string{line}
	matches := fuzzy.Matches{
		{Index: 0, MatchedIndexes: []int{0, 4, 15}},
	}

	got := highlightMatches(lines, matches)

	want := line[:0] + `["search_0"]` + line[0:16] + `[""]` + line[16:]
	assert.Equal(t, []string{want}, got)
}

func TestHighlightMatchesHandlesMultipleMatchesAcrossLines(t *testing.T) {
	lines := []string{
		`"Name": "cloudlens-demo-web"`,
		`"State": "running"`,
		`"Tag": "web-tier"`,
	}
	matches := fuzzy.Matches{
		{Index: 0, MatchedIndexes: []int{24, 25, 26}},
		{Index: 2, MatchedIndexes: []int{8, 9, 10}},
	}

	got := highlightMatches(lines, matches)

	assert.Equal(t, `"Name": "cloudlens-demo-`+`["search_0"]web[""]`+`"`, got[0])
	assert.Equal(t, `"State": "running"`, got[1])
	assert.Equal(t, `"Tag": "`+`["search_1"]web[""]`+`-tier"`, got[2])
}

func TestHighlightMatchesReturnsLinesUnchangedWhenNoMatches(t *testing.T) {
	lines := []string{`"State": "running"`}

	got := highlightMatches(lines, nil)

	assert.Equal(t, lines, got)
}

func TestHighlightMatchesHandlesSingleCharacterMatch(t *testing.T) {
	lines := []string{`abc`}
	matches := fuzzy.Matches{
		{Index: 0, MatchedIndexes: []int{1}},
	}

	got := highlightMatches(lines, matches)

	assert.Equal(t, []string{`a` + `["search_0"]b[""]` + `c`}, got)
}
