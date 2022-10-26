package graph

import (
	"context"
	"fmt"
	"github.com/Yegor-own/ghqllibrary/graph/model"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, title string, maxRentalPeriod string, coverLink string, authorID int) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: CreateBook - createBook"))
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, name string, photoLink string, dob string, dod string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: CreateAuthor - createAuthor"))
}

// CreateReader is the resolver for the createReader field.
func (r *mutationResolver) CreateReader(ctx context.Context, name string, email string) (*model.Reader, error) {
	panic(fmt.Errorf("not implemented: CreateReader - createReader"))
}

// RentBook is the resolver for the rentBook field.
func (r *mutationResolver) RentBook(ctx context.Context, bookID string, readerID string) (*model.Rent, error) {
	panic(fmt.Errorf("not implemented: RentBook - rentBook"))
}
