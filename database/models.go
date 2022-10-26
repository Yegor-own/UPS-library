package database

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	PhotoLink      string  `json:"photoLink"`
	Dob            string  `json:"dob"`
	Dod            string  `json:"dod"`
	Books          []*Book `json:"books"`
	AvailableBooks int     `json:"availableBooks"`
}

type Book struct {
	gorm.Model
	ID              int     `json:"id"`
	Title           string  `json:"title"`
	MaxRentalPeriod string  `json:"maxRentalPeriod"`
	CoverLink       string  `json:"coverLink"`
	AuthorID        int     `json:"authorId"`
	Author          *Author `json:"author"` // gorm:"foreignKey:AuthorID"
}

type Reader struct {
	gorm.Model
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Leases  []*Rent `json:"leases"`
	LateFee int     `json:"lateFee"`
}

type Rent struct {
	gorm.Model
	ID            int     `json:"id"`
	BookID        int     `json:"bookId"`
	Book          *Book   `json:"book"` // gorm:"foreignKey:BookID"
	ReaderID      int     `json:"readerId"`
	Reader        *Reader `json:"reader"` // gorm:"foreignKey:ReaderID"
	RentalTime    string  `json:"rentalTime"`
	RentalPeriod  string  `json:"rentalPeriod"`
	AmountPenalty int     `json:"amountPenalty"`
}
