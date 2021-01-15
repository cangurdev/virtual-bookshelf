package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"virtual-bookshelf/service"
)

func TestSearchService(t *testing.T) {
	//Arrange
	query := "sherlock"
	//Act
	result := service.Search(query)
	//Assert
	assert.NotNil(t, result)
}
func TestEmptySearchResult(t *testing.T) {
	//Arrange
	query := "netflix"
	//Act
	result := service.Search(query)
	//Assert
	assert.Nil(t, result)
}
