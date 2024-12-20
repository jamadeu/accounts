package user

import (
	"fmt"
	"net/mail"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Document == "" && r.Email == "" {
		return fmt.Errorf("reqest body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Document == "" {
		return errParamIsRequired("document", "string")
	}
	if r.Email == "" || validEmailFormat(r.Email) {
		return errParamIsRequired("email", "string")
	}
	return nil
}

func validEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Email != "" && validEmailFormat(r.Email) {
		return errParamIsRequired("email", "string")
	}
	if r.Name != "" || r.Document != "" || r.Email != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
