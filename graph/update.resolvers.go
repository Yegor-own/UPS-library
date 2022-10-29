package graph

import (
	"context"
	"github.com/Yegor-own/ghqllibrary/database/models"
	"github.com/Yegor-own/ghqllibrary/graph/model"
	"log"
	"strconv"
)

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, title *string, maxRentalPeriod *string, coverLink *string, authorID *int) (*model.Book, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}
	book := models.Book{
		ID:              uint(intId),
		Title:           *title,
		MaxRentalPeriod: *maxRentalPeriod,
		CoverLink:       *coverLink,
		AuthorID:        *authorID,
	}

	DBLib.Save(&book)
	book.GetById(DBLib, int(book.ID))

	res := model.Book{
		ID:              id,
		Title:           *title,
		MaxRentalPeriod: *maxRentalPeriod,
		CoverLink:       *coverLink,
		AuthorID:        *authorID,
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

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, name *string, photoLink *string, dob *string, dod *string) (*model.Author, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	author := models.Author{
		ID:        uint(intId),
		Name:      *name,
		PhotoLink: *photoLink,
		DOB:       *dob,
		DOD:       *dod,
	}
	DBLib.Save(&author)
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

// UpdateReader is the resolver for the updateReader field.
func (r *mutationResolver) UpdateReader(ctx context.Context, id string, name *string, email *string) (*model.Reader, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	reader := models.Reader{
		ID:    uint(intId),
		Name:  *name,
		Email: *email,
	}
	DBLib.Save(&reader)
	reader.GetById(DBLib, int(reader.ID))

	leases := []*model.Rent{}
	for _, rent := range reader.Rents {
		rent.Book.GetById(DBLib, rent.BookID)
		rent.Reader.GetById(DBLib, rent.ReaderID)
		leases = append(leases, &model.Rent{
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
