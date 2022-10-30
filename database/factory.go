package database

import (
	"encoding/json"
	"github.com/Yegor-own/ghqllibrary/database/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type BooksAuthorsJSON struct {
	Author string
	Title  string
}
type NamesJSON struct {
	Names []string `json:"names"`
}

func GenerateBooksAndAuthors() []models.Author {
	b, err := ioutil.ReadFile("./database/books_authors.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data []BooksAuthorsJSON
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatalln(err)
	}

	var authors []models.Author
	var books []models.Book

	author := make(map[string][]string)
	for _, datum := range data {
		author[datum.Author] = append(author[datum.Author], datum.Title)
	}

	for name, bookslist := range author {
		for _, title := range bookslist {
			books = append(books, models.Book{Title: title, MaxRentalPeriod: time.Until(time.Date(2022, 10, 28, 0, 0, 0, 0, time.UTC)).String()})
		}
		authors = append(authors, models.Author{Name: name, Books: books, AvailableBooks: len(books), DOB: time.Now().Format(time.RFC822), DOD: time.Now().Format(time.RFC822)})
		books = []models.Book{}
	}
	return authors
}

func GenerateReadersAndLeases(amount int, db *gorm.DB) []models.Reader {
	b, err := ioutil.ReadFile("./database/names.json")
	if err != nil {
		log.Fatalln(err)
	}

	names := NamesJSON{}
	err = json.Unmarshal(b, &names)
	if err != nil {
		log.Fatalln(err)
	}

	readers := []models.Reader{}

	for i := 0; i < amount; i++ {
		leases := []models.Rent{}

		var book models.Book
		db.First(&book)
		for i := 0; i < 3; i++ {
			var tmp models.Book
			tmp.GetById(db, i+int(book.ID))
			leases = append(leases, models.Rent{
				Book:          tmp,
				RentalTime:    time.Date(2022, 10, 31, 0, 0, 0, 0, time.UTC).Format(time.RFC822),
				RentalPeriod:  time.Until(time.Date(2022, 10, 28, 0, 0, 0, 0, time.UTC)).String(),
				AmountPenalty: 12,
			})
		}

		reader := models.Reader{
			Name:  names.Names[rand.Intn(len(names.Names))],
			Email: strings.ToLower(names.Names[rand.Intn(len(names.Names))]) + "@email.com",
			Rents: leases,
		}
		readers = append(readers, reader)
	}

	return readers
}
