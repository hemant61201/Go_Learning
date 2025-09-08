package test

import (
	"Simple_Programs/Day10-11/MockTest/model"
	"Simple_Programs/Day10-11/MockTest/repositary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByID_Success(t *testing.T) {

	mockRepo := new(repositary.MockBookRepository)

	mockBook := &model.Book{
		Id:     1,
		Title:  "Hard Things About Hard Things",
		Author: "Hemant",
	}

	mockRepo.On("FindById", int64(1)).Return(mockBook, nil)

	book, err := mockRepo.FindById(1)

	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, "Hemant", book.Author)
	assert.Equal(t, "Hard Things About Hard Things", book.Title)

	mockRepo.AssertExpectations(t)
}

func TestFindByID_Failed(t *testing.T) {

	mockRepo := new(repositary.MockBookRepository)

	mockRepo.On("FindById", int64(2)).Return(nil, assert.AnError)

	book, err := mockRepo.FindById(2)

	assert.Error(t, err)
	assert.Nil(t, book)

	mockRepo.AssertExpectations(t)
}
