package controllers

import (
	"github.com/HiBang15/golang-api-base-tulpo.git/constant"
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
	"github.com/HiBang15/golang-api-base-tulpo.git/services"
	"github.com/HiBang15/golang-api-base-tulpo.git/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//var userService = services.NewUserService()

func CreateActivity(c *gin.Context)  {
	var activity models.CreateActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		utils.SetResponse(c, http.StatusUnprocessableEntity, err, constant.INVALID_REQUEST_BODY, nil)
		return
	}

	//create user via user service
	//userClient := services.NewUserService()
	var activityService = services.NewActivityService()
	response, err := activityService.CreateActivity(activity)
	if err != nil {
		utils.SetResponse(c, http.StatusInternalServerError, err, constant.CANNOT_CREATE_ACTIVITY, nil)
		return
	}

	utils.SetResponse(c, http.StatusOK, nil, constant.CREATE_ACTIVITY_SUCCESS, response)
	return
}

func UpdateActivity(c *gin.Context)  {
	var activity models.UpdateActivityRequest
	if err := c.ShouldBindJSON(&activity); err != nil{
		utils.SetResponse(c,http.StatusUnprocessableEntity,err,constant.INVALID_REQUEST_BODY,nil)
		return
	}
	var activityService = services.NewActivityService()
	response , err := activityService.UpdateActivity(activity)
	if err!= nil{
		utils.SetResponse(c,http.StatusOK,nil,constant.UPDATE_ACTIVITY_SUCCESS, response)
		return
	}
}

func DeleteActivity(c *gin.Context)  {
	id := c.Param("id")

	if id == "" {
		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
		return
	}

	activityID, err := strconv.Atoi(id)
	if err != nil || activityID <= 0 {
		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
		return
	}

	activityService := services.NewActivityService()
	response, err := activityService.DeleteActivity(int32(activityID))
	if err != nil {
		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.CANNOT_DELETE_ACTIVITY, response)
		return
	}
	utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.DELETE_ACTIVITY_SUCCESSFUL, response)
	return
}

func GetActivityByID(c *gin.Context)  {
	id := c.Param("id")
	if id == "" {
		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
		return
	}

	activityID, err := strconv.Atoi(id)
	if err != nil || activityID <= 0 {
		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
		return
	}

	activityService := services.NewActivityService()
	response, err := activityService.GetActivityByID(int32(activityID))
	if err != nil {
		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.CANNOT_GETACTIVITYBYID_ACTIVITY, response)
		return
	}
	utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.GETACTIVITYBYID_ACTIVITY_SUCCESSFUL, response)
	return

}