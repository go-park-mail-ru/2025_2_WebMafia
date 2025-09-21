package model

import "fmt"

type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (i *UserInput) ValidateUserInput() error {
	if len(i.Name) < 5 {
		return fmt.Errorf("name is too short (minimum 5 chars)")
	}
	if len(i.Password) < 8 {
		return fmt.Errorf("password is too short (minimum 8 chars)")
	}
	return nil
}

