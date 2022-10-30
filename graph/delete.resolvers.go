package graph

import (
	"context"
	"github.com/Yegor-own/ghqllibrary/database/models"
	"github.com/Yegor-own/ghqllibrary/graph/model"
	"log"
	"math"
	"strconv"
	"time"
)

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.Book, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	book := models.Book{}
	book.GetById(DBLib, intId)

	DBLib.Delete(&book)

	res := model.Book{
		ID: id,
		//Title:           book.Title,
		//MaxRentalPeriod: book.MaxRentalPeriod,
		//CoverLink:       book.CoverLink,
		//AuthorID:        book.AuthorID,
		//Author: &model.Author{
		//	ID:             strconv.Itoa(int(book.AuthorID)),
		//	Name:           book.Author.Name,
		//	PhotoLink:      book.Author.PhotoLink,
		//	Dob:            book.Author.DOB,
		//	Dod:            book.Author.DOD,
		//	Books:          nil,
		//	AvailableBooks: book.Author.AvailableBooks,
		//},
	}

	return &res, nil
}

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	author := models.Author{}
	author.GetById(DBLib, intId)
	for _, book := range author.Books {
		DBLib.Delete(&book)
	}
	DBLib.Delete(&author)

	res := model.Author{ID: id}
	return &res, nil
}

// DeleteReader is the resolver for the deleteReader field.
func (r *mutationResolver) DeleteReader(ctx context.Context, id string) (*model.Reader, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	reader := models.Reader{}
	reader.GetById(DBLib, intId)
	DBLib.Delete(&reader)

	res := model.Reader{ID: id}
	return &res, nil
}

// ReturnBook is the resolver for the returnBook field.
func (r *mutationResolver) ReturnBook(ctx context.Context, bookID string, readerID string) (*model.Rent, error) {
	intId, err := strconv.Atoi(bookID)
	if err != nil {
		log.Fatalln(err)
	}

	intReaderId, err := strconv.Atoi(readerID)
	if err != nil {
		log.Fatalln(err)
	}

	rent := models.Rent{}
	reader := models.Reader{}
	reader.GetById(DBLib, intReaderId)

	err = DBLib.Model(&models.Rent{}).Where("reader_id = ?", intReaderId).First(&rent, "book_id = ?", intId).Error
	if err != nil {
		log.Fatalln(err)
	}

	rent.GetById(DBLib, int(rent.ID))

	rentalTime, err := time.Parse(time.RFC822, rent.RentalTime)
	if err != nil {
		log.Fatalln(err)
	}

	rentalPeriod, err := time.ParseDuration(rent.RentalPeriod)
	if err != nil {
		log.Fatalln(err)
	}

	duration := rentalPeriod - time.Until(rentalTime)
	log.Println(rentalPeriod, time.Until(rentalTime), duration)
	tmp := duration.String()

	if tmp[0:1] == "-" {
		delay := math.Abs(math.Floor(duration.Hours() / 24))
		log.Println(delay)
		reader.LateFee += rent.AmountPenalty * int(delay)
		log.Println(rent.AmountPenalty * int(delay))
		log.Println(reader.ID, reader.Name, reader.LateFee)
	}

	DBLib.Save(&reader)

	res := model.Rent{ID: strconv.Itoa(int(rent.ID))}

	DBLib.Delete(&rent)

	return &res, nil
}
