package service

import (
	"my_lib/internal/storage/db/sqlite"
	"my_lib/models/author"
	"my_lib/models/book"
)

type ServiceInterface interface {
	AddAuthor(a author.Author) (int, error)
	GetAuthorList() ([]author.Author, error)
	AddBook(b book.BookAdd) (book.BookAdd, error)
	GetPublishingHouseList() ([]book.PublishingHouse, error)
	GetBookCount() (int, error)
	GetBookList(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	IsAuthorExists(a author.Author) (bool, error)
}

type Service struct {
	ServiceInterface
}

func New(storage sqlite.StorageInterface) *Service {
	return &Service{
		ServiceInterface: NewService(storage),
	}
}
