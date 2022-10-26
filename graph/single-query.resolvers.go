package graph

import (
	"context"
	"fmt"
	"github.com/Yegor-own/ghqllibrary/graph/model"
)

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.Book, error) {
	return &model.Book{
		"12",
		"Text",
		"ll",
		"ll",
		2,
		nil,
	}, nil
	//book := models.Book{}
	//bookIdInt, err := strconv.Atoi(bookID)
	//if err != nil {
	//	panic(err)
	//}
	//
	//DBLib.First(&book, bookIdInt)
	//
	//var res model.Book
	//
	//data, _ := json.Marshal(book)
	//err = json.Unmarshal(data, &res)
	//if err != nil {
	//	panic(err)
	//}
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
