package usecase

import (
	"errors"
)

func (c AccountUc) GetBalance(accountID string) (float64, error) {
	account, error := c.r.FindByID(accountID)
	if error != nil {
		//TODO: add a sentinel
		return 0, errors.New("account not found")
	}
	return account.Balance, nil
}
