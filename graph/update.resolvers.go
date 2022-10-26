package graph

import (
	"context"
	"fmt"
	"github.com/Yegor-own/ghqllibrary/graph/model"
)

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, title *string, maxRentalPeriod *string, coverLink *string, authorID *int) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: UpdateBook - updateBook"))
}

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, name *string, photoLink *string, dob *string, dod *string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: UpdateAuthor - updateAuthor"))
}

// UpdateReader is the resolver for the updateReader field.
func (r *mutationResolver) UpdateReader(ctx context.Context, id string, name *string, email *string) (*model.Reader, error) {
	panic(fmt.Errorf("not implemented: UpdateReader - updateReader"))
}
