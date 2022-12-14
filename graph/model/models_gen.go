// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	PhotoLink      string  `json:"photoLink"`
	Dob            string  `json:"dob"`
	Dod            string  `json:"dod"`
	Books          []*Book `json:"books"`
	AvailableBooks int     `json:"availableBooks"`
}

type Book struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	MaxRentalPeriod string  `json:"maxRentalPeriod"`
	CoverLink       string  `json:"coverLink"`
	AuthorID        int     `json:"authorId"`
	Author          *Author `json:"author"`
}

type Reader struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Leases  []*Rent `json:"leases"`
	LateFee int     `json:"lateFee"`
}

type Rent struct {
	ID            string  `json:"id"`
	BookID        int     `json:"bookId"`
	Book          *Book   `json:"book"`
	ReaderID      int     `json:"readerId"`
	Reader        *Reader `json:"reader"`
	RentalTime    string  `json:"rentalTime"`
	RentalPeriod  string  `json:"rentalPeriod"`
	AmountPenalty int     `json:"amountPenalty"`
}
