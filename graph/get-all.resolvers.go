package graph

import (
	"context"
	"fmt"
	"github.com/Yegor-own/ghqllibrary/graph/model"
)

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented: Books - books"))
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented: Authors - authors"))
}

// Rents is the resolver for the rents field.
func (r *queryResolver) Rents(ctx context.Context) ([]*model.Rent, error) {
	panic(fmt.Errorf("not implemented: Rents - rents"))
}

// Readers is the resolver for the readers field.
func (r *queryResolver) Readers(ctx context.Context) ([]*model.Reader, error) {
	panic(fmt.Errorf("not implemented: Readers - readers"))
}
