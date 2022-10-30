package graph

import (
	"context"
	"github.com/Yegor-own/ghqllibrary/database/models"
	"github.com/Yegor-own/ghqllibrary/graph/model"
	"log"
	"strconv"
)

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.Book, error) {

	book := models.Book{}
	bookIdInt, err := strconv.Atoi(bookID)
	if err != nil {
		log.Fatalln(err)
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
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, authorID string) (*model.Author, error) {
	author := models.Author{}
	authorIntID, err := strconv.Atoi(authorID)
	if err != nil {
		log.Fatalln(err)
	}
	author.GetById(DBLib, authorIntID)

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

// Rent is the resolver for the rent field.
func (r *queryResolver) Rent(ctx context.Context, readerID string) (*model.Rent, error) {
	rent := models.Rent{}
	readerIntID, err := strconv.Atoi(readerID)
	if err != nil {
		log.Fatalln(err)
	}

	rent.GetById(DBLib, readerIntID)

	rent.Book.GetById(DBLib, rent.BookID)
	rent.Reader.GetById(DBLib, rent.ReaderID)
	rent.Book.Author.GetById(DBLib, rent.Book.AuthorID)

	leases := []*model.Rent{}
	for _, lease := range rent.Reader.Rents {
		leases = append(leases, &model.Rent{
			ID:            strconv.Itoa(int(lease.ID)),
			BookID:        lease.BookID,
			Book:          nil,
			ReaderID:      lease.ReaderID,
			Reader:        nil,
			RentalTime:    lease.RentalTime,
			RentalPeriod:  lease.RentalPeriod,
			AmountPenalty: lease.AmountPenalty,
		})
	}

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
				ID:             strconv.Itoa(int(rent.Book.Author.ID)),
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
			ID:      strconv.Itoa(int(rent.Reader.ID)),
			Name:    rent.Reader.Name,
			Email:   rent.Reader.Email,
			Leases:  leases,
			LateFee: rent.Reader.LateFee,
		},
		RentalTime:    rent.RentalTime,
		RentalPeriod:  rent.RentalPeriod,
		AmountPenalty: rent.AmountPenalty,
	}
	return &res, nil
}

// Reader is the resolver for the reader field.
func (r *queryResolver) Reader(ctx context.Context, readerID string) (*model.Reader, error) {
	reader := models.Reader{}
	readerIntID, err := strconv.Atoi(readerID)
	if err != nil {
		log.Fatalln(err)
	}

	reader.GetById(DBLib, readerIntID)
	leases := []*model.Rent{}
	for _, rent := range reader.Rents {
		rent.Book.GetById(DBLib, rent.BookID)
		leases = append(leases, &model.Rent{
			ID:     strconv.Itoa(int(rent.ID)),
			BookID: rent.BookID,
			Book: &model.Book{
				ID:              strconv.Itoa(rent.BookID),
				Title:           rent.Book.Title,
				MaxRentalPeriod: rent.Book.MaxRentalPeriod,
				CoverLink:       rent.Book.CoverLink,
				AuthorID:        rent.Book.AuthorID,
				Author: &model.Author{
					ID:             strconv.Itoa(rent.Book.AuthorID),
					Name:           rent.Book.Author.Name,
					PhotoLink:      rent.Book.Author.PhotoLink,
					Dob:            rent.Book.Author.DOB,
					Dod:            rent.Book.Author.DOD,
					Books:          nil,
					AvailableBooks: rent.Book.Author.AvailableBooks,
				},
			},
			ReaderID:      rent.ReaderID,
			Reader:        nil,
			RentalTime:    rent.RentalTime,
			RentalPeriod:  rent.RentalPeriod,
			AmountPenalty: rent.AmountPenalty,
		})
	}

	res := model.Reader{
		ID:      strconv.Itoa(int(reader.ID)),
		Name:    reader.Name,
		Email:   reader.Email,
		Leases:  leases,
		LateFee: reader.LateFee,
	}
	return &res, nil
}
