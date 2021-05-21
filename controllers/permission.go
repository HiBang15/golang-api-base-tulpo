package controllers

import (
	"github.com/HiBang15/golang-api-base-tulpo.git/constant"
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
	"github.com/HiBang15/golang-api-base-tulpo.git/services"
	"github.com/HiBang15/golang-api-base-tulpo.git/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

func CreatePermission(c *gin.Context) {
	var permission models.CreatePermissionRequest
	if err := c.ShouldBindJSON(&permission); err != nil {
		utils.SetResponse(c, http.StatusUnprocessableEntity, err, constant.INVALID_REQUEST_BODY, nil)
		return
	}
	var permissionService = services.NewPermissionService()
	response, err := permissionService.CreatePermission(permission)
	if err != nil {
		utils.SetResponse(c, http.StatusInternalServerError, err, constant.CANNOT_CREATE_PERMISSION, nil)
		return
	}
	utils.SetResponse(c, http.StatusOK, nil, constant.CREATE_PERMISSION_SUCCESS, response)
	return
}

//func UpdatePermission(c *gin.Context) {
//	var permission models.UpdatePermissionRe
//	if err := c.ShouldBindJSON(&permission); err != nil {
//		utils.SetResponse(c, http.StatusUnprocessableEntity, err, constant.INVALID_REQUEST_BODY, nil)
//		return
//	}
//	var permissionService = services.NewPermissionService()
//	response, err := permissionService.UpdatePermission(permission)
//	if err != nil {
//		utils.SetResponse(c, http.StatusOK, nil, constant.UPDATE_PERMISSION_SUCCESS, response)
//		return
//	}
//}
//
//func DeletePermission(c *gin.Context) {
//	id := c.Param("id")
//
//	if id == "" {
//		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
//		return
//	}
//
//	permissionID, err := strconv.Atoi(id)
//	if err != nil || permissionID <= 0 {
//		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
//		return
//	}
//
//	permissionService := services.NewPermissionService()
//	response, err := permissionService.DeletePermission(int32(permissionID))
//	if err != nil {
//		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.CANNOT_DELETE_PERMISSION, response)
//		return
//	}
//	utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.DELETE_PERMISSION_SUCCESSFUL, response)
//	return
//}
//
//func GetPermission(c *gin.Context) {
//
//	id := c.Param("id")
//	if id == "" {
//		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
//		return
//	}
//
//	permissionID, err := strconv.Atoi(id)
//	if err != nil || permissionID <= 0 {
//		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.INVALID_REQUEST_PARAM, nil)
//		return
//	}
//	permissionService := services.NewPermissionService()
//	response, err := permissionService.GetPermission(int32(permissionID))
//	if err != nil {
//		utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.CANNOT_GETPERMISSION, response)
//		return
//	}
//	utils.SetResponse(c, http.StatusUnprocessableEntity, nil, constant.GETPERMISSION_SUCCESSFUL, response)
//	return
//}
