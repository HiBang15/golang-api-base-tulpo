package models

import (
	"context"
	database "github.com/HiBang15/golang-api-base-tulpo.git/database/sqlc"
)

type CreatePermissionActivityRe struct {
	ActivityIds []int32 `json:"activity"`
}

type UpdatePermissionActivityRe struct {
	PermissionID int32   `json:"permission_id"`
	ActivityIds  []int32 `json:"activity_ids"`
}

//func (cnt *Connector) CreatePermissionActivity(ctx context.Context, request CreatePermissionActivityRe) (response database.PermissionActivity, err error) {
//	err = cnt.execTx(ctx, func(queries *database.Queries) error {
//		for _, item := range request.ActivityIds {
//			response, err = queries.CreatePermissionActivity(ctx, database.CreatePermissionActivityParams{
//				PermissionID: response.PermissionID,
//				ActivityID:   item,
//			})
//			if err != nil {
//				return err
//			}
//		}
//		return nil
//	})
//	if err != nil {
//		return database.PermissionActivity{}, err
//	}
//	return response, nil
//}
//
//func (cnt *Connector) UpdatePermissionActivity(ctx context.Context, request UpdatePermissionActivityRe) (response database.PermissionActivity, err error) {
//	err = cnt.execTx(ctx, func(queries *database.Queries) error {
//		for _, item := range request.ActivityIds {
//			response, err = queries.UpdatePermissionActivity(ctx, database.UpdatePermissionActivityParams{
//				PermissionID: request.PermissionID,
//				ActivityID:   item,
//			})
//			if err != nil {
//				return err
//			}
//		}
//		return nil
//	})
//	if err != nil {
//		return database.PermissionActivity{}, err
//	}
//	return response, nil
//}

func (cnt *Connector) DeletePermissionActivity(ctx context.Context, id int32) (bool, error) {
	err := cnt.execTx(ctx, func(queries *database.Queries) error {
		var err error
		err = queries.DeleteActivity(ctx, id)
		return err
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cnt *Connector) GetPermissionActivityByPermissionId(ctx context.Context, id int32) (response []database.PermissionActivity, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		response, err = queries.GetPermissionActivityByPermissionId(ctx, id)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
