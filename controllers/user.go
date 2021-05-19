package controllers

import (
	"github.com/HiBang15/golang-api-base-tulpo.git/constant"
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
	"github.com/HiBang15/golang-api-base-tulpo.git/services"
	"github.com/HiBang15/golang-api-base-tulpo.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//var userService = services.NewUserService()

func CreateUserAccount(c *gin.Context)  {
	var userInfoCreate models.CreateUserAccountRequest
	if err := c.ShouldBindJSON(&userInfoCreate); err != nil {
		utils.SetResponse(c, http.StatusUnprocessableEntity, err, constant.INVALID_REQUEST_BODY, nil)
		return
	}

	errValid := (validator.New()).Struct(userInfoCreate)
	if errValid != nil {
		utils.SetResponse(c, http.StatusUnprocessableEntity, errValid, constant.INVALID_REQUEST_BODY, nil)
		return
	}

	//create user via user service
	//userClient := services.NewUserService()
	var userService = services.NewUserService()
	user, err := userService.CreateUserAccount(userInfoCreate)
	if err != nil {
		utils.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}

	utils.SetResponse(c, http.StatusOK, nil, constant.CREATE_USERACCOUNT_SUCCESSFUL, user)
	return
}

//func GetUserAccountByUsernameOrEmail(c *gin.Context) {
//	email := c.Params.ByName("email")
//
//	//get user account
//	//userClient := services.NewUserService()
//	var userService = services.NewUserService()
//	res, err := userService.GetUserAccountByEmail(email)
//	if err != nil {
//		log.Printf("get user fails with error: %v", err)
//		utils.SetResponse(c, http.StatusInternalServerError, err, constant.CANNOT_GET_USERACCOUNT, nil)
//		return
//	}
//	utils.SetResponse(c, http.StatusOK, err, constant.GET_USERACCOUNT_SUCCESSFUL, res)
//	return
//}
