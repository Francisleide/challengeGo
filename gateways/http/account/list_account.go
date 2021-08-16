package account

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/francisleide/ChallangeGo/domain/entities"
)

// ShowAccount godoc
// @Summary Get accounts
// @Description List all accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} []entities.Account
// @Failure 400 "Failed to decode"
// @Failure 404 "Accounts not found"
// @Failure 500 "Unexpected internal server error"
// @Router /accounts [GET]
func (h Handler) List_all_accounts(w http.ResponseWriter, r *http.Request) {
	var accounts []entities.Account
	accounts = h.account.List_all_accounts()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		log.Fatal(err)
	}

}