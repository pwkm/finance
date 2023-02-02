package app

import (
	"github.com/pwkm/finance/pkg/domain"
)

type bookingSvc struct {
	DB domain.BookingDB
}

func NewBookingSvc(db domain.BookingDB) domain.BookingSvc {
	return bookingSvc{
		DB: db,
	}
}
