package _interface

import "github.com/HiBang15/golang-api-base-tulpo.git/models"

type UserInterface interface {
	CreateUserAccount (request models.CreateUserAccountRequest) (response models.UserAccount, err error)
	GetUserAccountByEmail(email string) (response models.UserAccount, err error)
}