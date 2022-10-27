package graph

import (
	"context"
	"fmt"
	"github.com/Yegor-own/ghqllibrary/database/models"
	"github.com/Yegor-own/ghqllibrary/graph/model"
	"strconv"
)

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.Book, error) {
	//return &model.Book{
	//	"12",
	//	"Text",
	//	"ll",
	//	"ll",
	//	2,
	//	nil,
	//}, nil

	book := models.Book{}
	bookIdInt, err := strconv.Atoi(bookID)
	if err != nil {
		panic(err)
	}

	book.GetById(DBLib, bookIdInt)

	var books []*model.Book
	book.Author.GetById(DBLib, book.AuthorID)

	for _, b := range book.Author.Books {
		tmp := model.Book{
			ID:              strconv.Itoa(int(b.ID)),
			Title:           b.Title,
			MaxRentalPeriod: b.MaxRentalPeriod,
			CoverLink:       b.CoverLink,
			AuthorID:        b.AuthorID,
			Author:          nil,
		}
		books = append(books, &tmp)
	}
	//log.Println(&books[0])

	var (
		res = model.Book{
			ID:              strconv.Itoa(int(book.ID)),
			Title:           book.Title,
			MaxRentalPeriod: book.MaxRentalPeriod,
			CoverLink:       book.CoverLink,
			AuthorID:        book.AuthorID,
			Author: &model.Author{
				ID:             strconv.Itoa(int(book.Author.ID)),
				Name:           book.Author.Name,
				PhotoLink:      book.Author.PhotoLink,
				Dob:            book.Author.DOB,
				Dod:            book.Author.DOD,
				Books:          books,
				AvailableBooks: book.Author.AvailableBooks,
			},
		}
	)
	return &res, nil
	//
	//log.Println(book, res)
	//
	//return &res, nil
	//db.First(&user, 10)
	//panic(fmt.Errorf("not implemented: Book - book"))
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, authorID string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Rent is the resolver for the rent field.
func (r *queryResolver) Rent(ctx context.Context, readerID string) (*model.Rent, error) {
	panic(fmt.Errorf("not implemented: Rent - rent"))
}

// Reader is the resolver for the reader field.
func (r *queryResolver) Reader(ctx context.Context, readerID string) (*model.Reader, error) {
	panic(fmt.Errorf("not implemented: Reader - reader"))
}
