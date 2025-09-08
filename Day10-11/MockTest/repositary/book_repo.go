package repositary

import "Simple_Programs/Day10-11/MockTest/model"

type BookRepository interface {
	FindById(id int64) (*model.Book, error)
}
