package entities

import (
	"errors"
	"time"

	"github.com/satori/uuid.go"
)

type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               float64
	CreatedAt            string
}


func ValidateAmount(amount float64) bool {
	return amount > 0

}

func NewTransfer(accountOriginID, accountDestinationID string, amount float64) (Transfer, error) {
	if !ValidateAmount(amount) {
		//TODO: add a sentinel
		return Transfer{}, errors.New("invalid amount")
	}
	return Transfer{
		ID:                   uuid.NewV4().String(),
		AccountOriginID:      accountOriginID,
		AccountDestinationID: accountDestinationID,
		Amount:               amount,
		CreatedAt:            time.Now().UTC().String(),
	}, nil

}
