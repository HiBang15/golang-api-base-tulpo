package utils

import (
	"encoding/json"
	database "github.com/HiBang15/golang-api-base-tulpo.git/database/sqlc"
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
)

func MapDataToStruct(in interface{}, out interface{}) (err error) {
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, out)
	return err
}

func ConvertUserAccount(data database.UserAccount) models.UserAccount {
	return models.UserAccount{
		ID:               data.ID,
		FirstName:        data.FirstName.String,
		LastName:         data.LastName.String,
		FullName:         data.FullName.String,
		Email:            data.Email,
		Address:          data.Address.String,
		Password:         data.Password,
		PhoneNumber:      data.PhoneNumber.String,
		VerifyEmail:      data.VerifyEmail.Bool,
		RegistrationTime: data.RegistrationTime.Time.Unix(),
		CreatedAt:        data.CreatedAt.Time.Unix(),
		UpdatedAt:        data.UpdatedAt.Time.Unix(),
	}
}
