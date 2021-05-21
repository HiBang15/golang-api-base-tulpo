package _interface

import "github.com/HiBang15/golang-api-base-tulpo.git/models"

type PermissionInterface interface {
	CreatePermission (request models.CreatePermissionRequest)(response models.Permission,err error)
	//UpdatePermission(request models.UpdatePermissionRe)(response models.Permission,err error)
	//DeletePermission(id int32)(bool,error)
	//GetPermission(id int32)(response models.Permission,err error)
}