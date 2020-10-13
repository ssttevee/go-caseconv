package caseconv_test

import (
	"testing"

	"github.com/ssttevee/go-caseconv"
	"github.com/stretchr/testify/assert"
)

func TestSplitWords(t *testing.T) {
	assert.Equal(t, []string{"Some", "Words", "In", "Pascal", "Case"}, caseconv.SplitWords("SomeWordsInPascalCase"))
	assert.Equal(t, []string{"some", "Words", "In", "Camel", "Case"}, caseconv.SplitWords("someWordsInCamelCase"))
	assert.Equal(t, []string{"some", "words", "in", "snake", "case"}, caseconv.SplitWords("some_words_in_snake_case"))
	assert.Equal(t, []string{"some", "words", "in", "kebab", "case"}, caseconv.SplitWords("some-words-in-kebab-case"))
	assert.Equal(t, []string{"All", "CAPS", "Makes", "An", "Acronym"}, caseconv.SplitWords("AllCAPSMakesAnAcronym"))
	assert.Equal(t, []string{"some", "1", "Words", "2", "And", "3", "Numbers", "4"}, caseconv.SplitWords("some1Words2_And3Numbers4"))
}
