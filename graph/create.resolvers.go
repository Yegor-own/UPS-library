package graph

import (
	"context"
	"github.com/Yegor-own/ghqllibrary/database/models"
	"github.com/Yegor-own/ghqllibrary/graph/model"
	"log"
	"strconv"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, title string, maxRentalPeriod string, coverLink string, authorID int) (*model.Book, error) {
	book := models.Book{
		Title:           title,
		MaxRentalPeriod: maxRentalPeriod,
		CoverLink:       coverLink,
		AuthorID:        authorID,
	}
	DBLib.Create(&book)
	book.GetById(DBLib, int(book.ID))

	res := model.Book{
		ID:              strconv.Itoa(int(book.ID)),
		Title:           book.Title,
		MaxRentalPeriod: book.MaxRentalPeriod,
		CoverLink:       book.CoverLink,
		AuthorID:        book.AuthorID,
		Author: &model.Author{
			ID:             strconv.Itoa(int(book.AuthorID)),
			Name:           book.Author.Name,
			PhotoLink:      book.Author.PhotoLink,
			Dob:            book.Author.DOB,
			Dod:            book.Author.DOD,
			Books:          nil,
			AvailableBooks: book.Author.AvailableBooks,
		},
	}
	return &res, nil
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, name string, photoLink string, dob string, dod string) (*model.Author, error) {
	author := models.Author{
		Name:      name,
		PhotoLink: photoLink,
		DOB:       dob,
		DOD:       dod,
	}
	DBLib.Create(&author)
	author.GetById(DBLib, int(author.ID))

	books := []*model.Book{}
	for _, book := range author.Books {
		books = append(books, &model.Book{
			ID:              strconv.Itoa(int(book.ID)),
			Title:           book.Title,
			MaxRentalPeriod: book.MaxRentalPeriod,
			CoverLink:       book.CoverLink,
			AuthorID:        book.AuthorID,
			Author:          nil,
		})
	}

	res := model.Author{
		ID:             strconv.Itoa(int(author.ID)),
		Name:           author.Name,
		PhotoLink:      author.PhotoLink,
		Dob:            author.DOB,
		Dod:            author.DOD,
		Books:          books,
		AvailableBooks: author.AvailableBooks,
	}
	return &res, nil
}

// CreateReader is the resolver for the createReader field.
func (r *mutationResolver) CreateReader(ctx context.Context, name string, email string) (*model.Reader, error) {
	reader := models.Reader{
		Name:  name,
		Email: email,
	}
	DBLib.Create(&reader)
	reader.GetById(DBLib, int(reader.ID))

	res := model.Reader{
		ID:      strconv.Itoa(int(reader.ID)),
		Name:    reader.Name,
		Email:   reader.Email,
		Leases:  nil,
		LateFee: reader.LateFee,
	}

	return &res, nil
}

// RentBook is the resolver for the rentBook field.
func (r *mutationResolver) RentBook(ctx context.Context, bookID string, readerID string) (*model.Rent, error) {
	bookIntID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println(err)
	}
	readerIntID, err := strconv.Atoi(readerID)
	if err != nil {
		log.Println(err)
	}

	rent := models.Rent{
		BookID:   bookIntID,
		ReaderID: readerIntID,
	}
	DBLib.Create(&rent)
	rent.Book.GetById(DBLib, rent.BookID)
	rent.Reader.GetById(DBLib, rent.ReaderID)
	rent.Book.Author.GetById(DBLib, rent.Book.AuthorID)

	res := model.Rent{
		ID:     strconv.Itoa(int(rent.ID)),
		BookID: rent.BookID,
		Book: &model.Book{
			ID:              strconv.Itoa(int(rent.Book.ID)),
			Title:           rent.Book.Title,
			MaxRentalPeriod: rent.Book.MaxRentalPeriod,
			CoverLink:       rent.Book.CoverLink,
			AuthorID:        rent.Book.AuthorID,
			Author: &model.Author{
				ID:             strconv.Itoa(int(rent.Book.AuthorID)),
				Name:           rent.Book.Author.Name,
				PhotoLink:      rent.Book.Author.PhotoLink,
				Dob:            rent.Book.Author.DOB,
				Dod:            rent.Book.Author.DOD,
				Books:          nil,
				AvailableBooks: rent.Book.Author.AvailableBooks,
			},
		},
		ReaderID: rent.ReaderID,
		Reader: &model.Reader{
			ID:      strconv.Itoa(rent.ReaderID),
			Name:    rent.Reader.Name,
			Email:   rent.Reader.Email,
			Leases:  nil,
			LateFee: rent.Reader.LateFee,
		},
		RentalTime:    rent.RentalTime,
		RentalPeriod:  rent.RentalPeriod,
		AmountPenalty: rent.AmountPenalty,
	}
	return &res, nil
}
