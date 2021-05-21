package services

import (
	"context"
	_interface "github.com/HiBang15/golang-api-base-tulpo.git/interface"
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
	"github.com/HiBang15/golang-api-base-tulpo.git/utils"
	"log"
)

type PermissionService struct {
	Connector *models.Connector
}



func NewPermissionService() _interface.PermissionInterface {
	connect := models.NewConnector(connDB)
	return &PermissionService{Connector: connect}
}

func (p PermissionService) CreatePermission(request models.CreatePermissionRequest) (response models.Permission, err error) {
	log.Println("receive a create permission")

	permission, err := p.Connector.CreatePermissions(context.Background(), request)
	if err != nil {
		return models.Permission{}, err
	}


	var listActivity []models.Activity
	if len(request.ActivityIds) > 0 {
		permissionActivity, err := p.Connector.GetPermissionActivityByPermissionId(context.Background(), permission.ID)
		if err != nil {
			return models.Permission{}, err
		}
		for _, item := range permissionActivity {
			activity, err := p.Connector.GetActivityByID(context.Background(), item.ActivityID)
			if err != nil {
				return models.Permission{}, err
			}

			i := utils.ConvertActivity(activity)
			listActivity = append(listActivity, i)
		}
	}

	response = utils.ConvertPermission(permission, listActivity)
	return response, nil
}

func (p PermissionService) UpdatePermission(request interface{}) (response interface{}, err error) {
	panic("implement me")
}

func (p PermissionService) DeletePermission(id int32) (bool, error) {
	panic("implement me")
}

func (p PermissionService) GetPermission(id int32) (response interface{}, err error) {
	panic("implement me")
}