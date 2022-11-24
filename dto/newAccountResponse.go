package dto

// the api should return the new account id, when the new account is opened with the status
// code as 201

type NewAccountResponse struct {
	AccountId string `json:"account_id"`
}
