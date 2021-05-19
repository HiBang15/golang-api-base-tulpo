package _interface

import "github.com/HiBang15/golang-api-base-tulpo.git/models"

type ActivityInterface interface {
	CreateActivity(resquest models.CreateActivity)
}
