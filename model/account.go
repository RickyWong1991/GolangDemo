package model

import (
	"errors"
	"fmt"
	"tutorial/golangdemo/model/database"

	uuid "github.com/gofrs/uuid"
)

var accountMaxID = 3
var accounts = []Account{
	{ID: 1, Name: "account_1"},
	{ID: 2, Name: "account_2"},
	{ID: 3, Name: "account_3"},
}

// Account example
type Account struct {
	//gorm.Model
	ID   int       `json:"id" example:"1" format:"int64"`
	Name string    `json:"name" example:"account name"`
	UUID uuid.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

// example
var (
	ErrNameInvalid = errors.New("name is empty")
)

// AddAccount example
type AddAccount struct {
	Name string `json:"name" example:"account name"`
}

// Validation example
func (a *AddAccount) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UpdateAccount example
type UpdateAccount struct {
	Name string `json:"name" example:"account name"`
}

// Validation example
func (u *UpdateAccount) Validation() error {
	switch {
	case len(u.Name) == 0:
		return ErrNameInvalid
	default:
		return nil

	}
}

// AccountsAll example
func AccountsAll(q string) ([]Account, error) {
	if q == "" {
		return accounts, nil
	}

	as := []Account{}

	for k, v := range accounts {
		if q == v.Name {
			as = append(as, accounts[k])
		}
	}
	return as, nil
}

// AccountOne example
func AccountOne(id int) (Account, error) {
	for _, v := range accounts {
		if id == v.ID {
			return v, nil
		}
	}
	return Account{}, ErrNoRow
}

// Insert example
func (a Account) Insert() (int, error) {
	accountMaxID++
	a.ID = accountMaxID
	a.Name = fmt.Sprintf("account_%d", accountMaxID)
	accounts = append(accounts, a)
	return accountMaxID, nil
}

// Delete example
func Delete(id int) error {
	for k, v := range accounts {
		if id == v.ID {
			accounts = append(accounts[:k], accounts[k+1:]...)
		}
	}
	return fmt.Errorf("account id=%d is not found", id)
}

func (a *Account) Update() error {
	for k, v := range accounts {
		if v.ID == a.ID {
			accounts[k].Name = a.Name
			return nil
		}
	}
	return fmt.Errorf("account id=%d is not found", a.ID)
}

func FindUserByUserName(username string) (Account, error) {
	var account Account
	err := database.Database.Where("name", username).Find(&account).Error

	if err != nil {
		return Account{}, err
	}

	return account, nil
}
