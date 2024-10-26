package account

import (
	"fmt"
)

type CreateAccountRequest struct {
	Balance float64 `json:"accountBalance"`
	UserId  uint    `json:"userId"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateAccountRequest) Validate() error {
	if r.Balance < 0 && r.UserId <= 0 {
		return fmt.Errorf("reqest body is empty or malformed")
	}
	if r.Balance < 0 {
		return errParamIsRequired("balance", "float64")
	}
	if r.UserId <= 0 {
		return errParamIsRequired("UserId", "uint")
	}
	return nil
}
