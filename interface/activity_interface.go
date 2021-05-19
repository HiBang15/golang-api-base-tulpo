package _interface

import (
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
)

type ActivityInterface interface {
	CreateActivity(request models.CreateActivity) (response models.Activity, err error)
	UpdateActivity(request models.UpdateActivityRequest) (response models.Activity, err error)
	DeleteActivity(id int32) (bool, error)
	GetActivityByID(id int32)(response models.Activity,err error)
}
