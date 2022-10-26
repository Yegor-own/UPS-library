package graph

import (
	"context"
	"fmt"
	"github.com/Yegor-own/ghqllibrary/graph/model"
)

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: DeleteBook - deleteBook"))
}

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: DeleteAuthor - deleteAuthor"))
}

// DeleteReader is the resolver for the deleteReader field.
func (r *mutationResolver) DeleteReader(ctx context.Context, id string) (*model.Reader, error) {
	panic(fmt.Errorf("not implemented: DeleteReader - deleteReader"))
}

// ReturnBook is the resolver for the returnBook field.
func (r *mutationResolver) ReturnBook(ctx context.Context, bookID string, readerID string) (*model.Rent, error) {
	panic(fmt.Errorf("not implemented: ReturnBook - returnBook"))
}
