package models

import (
	"context"
	"database/sql"
	database "github.com/HiBang15/golang-api-base-tulpo.git/database/sqlc"
	"log"
	"time"
)

type UserAccount struct {
	ID               int32  `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	Address          string `json:"address"`
	TimeZone         string `json:"time_zone"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phone_number"`
	VerifyEmail      bool   `json:"verify_email"`
	RegistrationTime int64  `json:"registration_time"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
}

type CreateUserAccountRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Address         string `json:"address"`
	TimeZone        string `json:"time_zone"`
	Password        string `json:"password"`
	PasswordCost    string `json:"password_cost"`
	PhoneNumber     string `json:"phone_number"`
	AcceptMarketing bool   `json:"accept_marketing"`
	CodeVerifyEmail int32 `json:"code_verify_email"`
}

func (cnt *Connector) CreateUserAccount(ctx context.Context, request CreateUserAccountRequest) (response database.UserAccount, err error)  {
	//var user database.UserAccount
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		response, err = queries.CreateUserAccount(ctx, database.CreateUserAccountParams{
			FirstName:              sql.NullString{
				String: request.FirstName,
				Valid:  true,
			},
			LastName:               sql.NullString{
				String: request.LastName,
				Valid:  true,
			},
			FullName:               sql.NullString{
				String: request.FullName,
				Valid:  true,
			},
			Address:                sql.NullString{
				String: request.Address,
				Valid:  true,
			},
			Email:                  request.Email,
			Password:               request.Password,
			PhoneNumber:            sql.NullString{
				String: request.PhoneNumber,
				Valid:  true,
			},
			//todo fix data for test sign-in
			VerifyEmail: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
			AcceptsMarketing:       false,
			CodeVerifyEmail:        sql.NullInt32{
				Int32: request.CodeVerifyEmail,
				Valid: true,
			},
			RegistrationTime:       sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		})
		return err
	})
	if err != nil {
		log.Println("Create user fail with err: ", err.Error())
		return database.UserAccount{}, err
	}
	
	//response = utils.ConvertUserAccount(user)
	return response, nil
}

func (cnt *Connector)GetUserAccountByEmail(ctx context.Context, email string) (response database.UserAccount, err error)  {
	//var user database.UserAccount
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		response, err = queries.GetUserAccountByUsernameOrEmail(ctx, email)
		return err
	})

	if err != nil {
		log.Println("Get user by email fail with err: ", err.Error())
		return database.UserAccount{}, err
	}

	//response = utils.ConvertUserAccount(user)
	return response, nil
}

func (cnt *Connector) CheckEmailExists(ctx context.Context, email string) (success bool, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		success, err = queries.CheckEmailExists(ctx, email)
		return err
	})
	if err != nil {
		log.Println("[UserService][Model]Check email exists fail with err: ", err.Error())
		return true, err
	}
	return success, nil
}

func (cnt *Connector) CheckPhoneNumberExists(ctx context.Context, phone string) (success bool, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		success, err = queries.CheckPhoneNoExists(ctx, sql.NullString{phone, true})
		return err
	})
	if err != nil {
		log.Println("[UserService][Model]Check phone exists fail with err: ", err.Error())
		return true, err
	}
	return success, nil
}

