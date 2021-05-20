package models

import (
	"context"
	"database/sql"

	database "github.com/HiBang15/golang-api-base-tulpo.git/database/sqlc"
)

type Activity struct {
	Id       int32  `json:"id"`
	Url      string `json:"url"`
	Method   string `json:"method"`
	UrlRegex string `json:"url_regex"`
}

type CreateActivity struct {
	Url      string `json:"url"`
	Method   string `json:"method"`
	UrlRegex string `json:"url_regex"`
}

type UpdateActivityRequest struct {
	Id       int32  `json:"id"`
	Url      string `json:"url"`
	Method   string `json:"method"`
	UrlRegex string `json:"url_regex"`
}

func (cnt *Connector) CreateActivity(ctx context.Context, request CreateActivity) (response database.Activity, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		response, err = queries.CreateActivity(ctx, database.CreateActivityParams{
			Url: request.Url,
			Method: sql.NullString{
				String: request.Method,
				Valid:  true,
			},
			UrlRegex: request.UrlRegex,
		})
		return err
	})
	if err != nil {
		//@todo log err
		return database.Activity{}, err
	}
	return response, nil
}
func (cnt *Connector) UpdateActivity(ctx context.Context, request UpdateActivityRequest) (response database.Activity, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		response, err = queries.UpdateActivity(ctx, database.UpdateActivityParams{
			ID:  request.Id,
			Url: request.Url,
			Method: sql.NullString{
				String: request.Method,
				Valid:  true,
			},
			UrlRegex: request.UrlRegex,
		})
		return err
	})
	if err != nil {
		return database.Activity{}, err
	}
	return response, nil
}

func (cnt *Connector) DeleteActivity(ctx context.Context, id int32) (bool, error) {
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

func (cnt *Connector) GetActivityByID(ctx context.Context, id int32) (response database.Activity, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		var err error
		response, err = queries.GetActivityByID(ctx)
		return err
	})
	if err != nil {
		return database.Activity{}, err
	}
	return response, nil
}

func (cnt *Connector) GetListActivity(ctx context.Context) (response []database.Activity, err error) {
	err = cnt.execTx(ctx, func(queries *database.Queries) error {
		var err error
		response, err = queries.GetListActivity(ctx)
		return err
	})
	if err != nil {
		return []database.Activity{}, err
	}
	return response, nil
}
