package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/HiBang15/golang-api-base-tulpo.git/constant"
	_interface "github.com/HiBang15/golang-api-base-tulpo.git/interface"
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
	"github.com/HiBang15/golang-api-base-tulpo.git/utils"
	"log"
	"github.com/joho/godotenv"
	"os"
	_ "github.com/lib/pq"
)

var connDB *sql.DB

func init() {
	// load config from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Define connect to DB
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("DB_USER"),
		os.Getenv("DB_USER_PASSWORD"), os.Getenv("DB_USER_HOST"),
		os.Getenv("DB_USER_PORT"), os.Getenv("DB_USER_NAME"), os.Getenv("DB_SSL_MODE"))
	connDB, err = sql.Open(os.Getenv("DB_DRIVER"), dbSource)
	if err != nil {
		log.Fatalf("has error occur when init connect to DB: %v", err)
	}
}

type UserService struct {
	Connector *models.Connector
}

func NewUserService() _interface.UserInterface {
	connect := models.NewConnector(connDB)
	return &UserService{Connector: connect}
}

func (u *UserService) CreateUserAccount(request models.CreateUserAccountRequest) (response models.UserAccount, err error) {
	log.Println("receive create user account with info: ", request)

	//validate email
	if isEmail := utils.IsEmailValid(request.Email); !isEmail {
		log.Println("[UserService][Service] Invalid Email!")
		return models.UserAccount{}, errors.New(constant.INVALID_EMAIL)
	}

	if emailExists, _ := u.Connector.CheckEmailExists(context.Background(), request.Email); emailExists {
		log.Println("[UserService][Service] Email Already Exists!")
		return models.UserAccount{}, errors.New(constant.EMAIL_ALREADY_EXISTS)
	}

	//validate password
	if isPassword := utils.IsPassword(request.Password); !isPassword {
		log.Println("[UserService][Service] Invalid Password!")
		return models.UserAccount{}, errors.New(constant.INVALID_PASSWORD)
	}

	//hash password
	passwordHash, err := utils.HashPassword(request.Password)
	if err != nil {
		log.Println("[UserService][Service] Hash password fail with error: ", err.Error())
		return models.UserAccount{}, errors.New("Create user account fail with err: "+err.Error())
	}

	//valid phoneNo
	if request.PhoneNumber != "" {
		if phoneExists, _ := u.Connector.CheckPhoneNumberExists(context.Background(), request.PhoneNumber); phoneExists {
			log.Println("[UserService][Service] Phone number Already Exists!")
			return models.UserAccount{}, errors.New(constant.PHONE_NUMBER_EXISTS)
		}
	}

	request.Password = passwordHash
	request.PasswordCost = os.Getenv("PASSWORD_COST")

	//random code verify email
	codeVerifyEmail := utils.RandomInt32(100000, 999999)
	request.CodeVerifyEmail = codeVerifyEmail

	var req models.CreateUserAccountRequest
	err = utils.MapDataToStruct(request, &req)

	userCreated, err := u.Connector.CreateUserAccount(context.Background(), req)
	if err != nil {
		log.Println("[UserService][Service] Create user account fail with error: ", err.Error())
		return models.UserAccount{}, errors.New("Create user account fail with err: "+err.Error())
	}

	//todo send mail
	response = utils.ConvertUserAccount(userCreated)

	log.Println("create account successful!")
	return response, nil
}

func (u *UserService) GetUserAccountByEmail(email string) (response models.UserAccount, err error) {
	log.Println("receive get user account by username or email with info ", email)

	isEmail := utils.IsEmailValid(email)
	if !isEmail{
		log.Println("[UserService][Service] Invalid Email !")
		return models.UserAccount{}, errors.New(constant.INVALID_EMAIL)
	}
	userAccount, err := u.Connector.GetUserAccountByEmail(context.Background(), email)
	if err != nil {
		log.Println("[UserService][Service]Get userAccount fail with err: ", err.Error())
		return models.UserAccount{}, errors.New("Get user account fail with err: "+err.Error())
	}

	response = utils.ConvertUserAccount(userAccount)

	return response, nil
}
