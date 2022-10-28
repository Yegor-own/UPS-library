package graph

import (
	"context"
	"github.com/Yegor-own/ghqllibrary/database/models"
	"github.com/Yegor-own/ghqllibrary/graph/model"
	"log"
	"strconv"
)

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	books := []models.Book{}
	DBLib.Find(&books)

	res := []*model.Book{}
	for _, book := range books {
		book.Author.GetById(DBLib, book.AuthorID)
		res = append(res, &model.Book{
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
				Books:          nil,
				AvailableBooks: book.Author.AvailableBooks,
			},
		})
	}
	return res, nil
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	authors := []models.Author{}
	DBLib.Find(&authors)

	res := []*model.Author{}
	for _, author := range authors {
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

		res = append(res, &model.Author{
			ID:             strconv.Itoa(int(author.ID)),
			Name:           author.Name,
			PhotoLink:      author.PhotoLink,
			Dob:            author.DOB,
			Dod:            author.DOD,
			Books:          books,
			AvailableBooks: author.AvailableBooks,
		})
	}

	return res, nil
}

// Rents is the resolver for the rents field.
func (r *queryResolver) Rents(ctx context.Context, readerID string) ([]*model.Rent, error) {
	readerIntID, err := strconv.Atoi(readerID)
	if err != nil {
		log.Println(err)
	}

	rents := []models.Rent{}
	DBLib.Where("id = ?", readerIntID).Find(&rents)

	res := []*model.Rent{}
	for _, rent := range rents {
		rent.Book.GetById(DBLib, rent.BookID)
		rent.Reader.GetById(DBLib, rent.ReaderID)

		res = append(res, &model.Rent{
			ID:     strconv.Itoa(int(rent.ID)),
			BookID: rent.BookID,
			Book: &model.Book{
				ID:              strconv.Itoa(int(rent.Book.ID)),
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
		})
	}

	return res, nil
}

// Readers is the resolver for the readers field.
func (r *queryResolver) Readers(ctx context.Context) ([]*model.Reader, error) {
	return []*model.Reader{}, nil
}
