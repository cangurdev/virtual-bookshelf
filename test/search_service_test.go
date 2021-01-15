package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"virtual-bookshelf/service"
)

func TestSearchService(t *testing.T) {
	result := service.Search("sherlock")
	assert.NotNil(t, result)
}
func TestEmptySearchResult(t *testing.T) {
	result := service.Search("netflix")
	assert.Nil(t, result)
}
