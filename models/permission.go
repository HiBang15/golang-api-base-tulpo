package models

import (
	"context"

	"errors"
	"fmt"
	database "github.com/HiBang15/golang-api-base-tulpo.git/database/sqlc"
)

type Permission struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Activities []Activity `json:"activities"`
}
type CreatePermissionRequest struct {
	Name        string  `json:"name"`
	ActivityIds []int32 `json:"activity"`
}

type UpdatePermissionRequest struct {
	ID   int32  `json:"permission_id"`
	Name string `json:"name"`
}

func (cnt *Connector) CreatePermissions(ctx context.Context, request CreatePermissionRequest) (response database.Permission, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		response, err = queries.CreatePermission(ctx, request.Name)
		if err != nil {
			return err
		}

		if len(request.ActivityIds) > 0 {
			for _, item := range request.ActivityIds {
				var activityExists bool
				activityExists, err = queries.CheckActivityExists(ctx, item)
				if err != nil {
					return err
				}
				if activityExists == false {
					return errors.New("ActivityID " + fmt.Sprintf("%d", item) + "does not exists!")
				}
				_, err = queries.CreatePermissionActivity(ctx, database.CreatePermissionActivityParams{
					PermissionID: response.ID,
					ActivityID:   item,
				})
				if err != nil {
					return err
				}
			}
		}

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return database.Permission{}, err
	}
	return response, nil
}

func (cnt *Connector) UpdatePermission(ctx context.Context, request UpdatePermissionRequest) (response database.Permission, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		response, err = queries.UpdatePermission(ctx, database.UpdatePermissionParams{
			ID:   request.ID,
			Name: request.Name,
		})
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return database.Permission{}, err
	}
	return response, nil
}
