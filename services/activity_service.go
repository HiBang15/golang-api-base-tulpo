package services

import (
	"context"
	"log"

	_interface "github.com/HiBang15/golang-api-base-tulpo.git/interface"
	"github.com/HiBang15/golang-api-base-tulpo.git/models"
	"github.com/HiBang15/golang-api-base-tulpo.git/utils"
	_ "github.com/lib/pq"
)

type ActivityService struct {
	Connector *models.Connector
}

func NewActivityService() _interface.ActivityInterface {
	connect := models.NewConnector(connDB)
	return &ActivityService{Connector: connect}
}

func (a *ActivityService) CreateActivity(request models.CreateActivity) (response models.Activity, err error) {
	log.Println("receive a create activity")

	activity, err := a.Connector.CreateActivity(context.Background(), request)
	if err != nil {
		//todo log
		return models.Activity{}, err
	}

	response = utils.ConvertActivity(activity)
	log.Println("")
	return response, nil
}

func (a *ActivityService) UpdateActivity(request models.UpdateActivityRequest) (response models.Activity, err error) {
	log.Println("receive a update activity")

	activity, err := a.Connector.UpdateActivity(context.Background(), request)
	if err != nil {
		return models.Activity{}, err
	}

	response = utils.ConvertActivity(activity)
	log.Println("update activity success")
	return response, nil
}

func (a *ActivityService) DeleteActivity(id int32) (bool, error) {
	log.Println("receive delete activity")
	activity, err := a.Connector.DeleteActivity(context.Background(), id)
	if err != nil {
		return false, err
	}
	log.Println("delete activity success")
	return activity, nil

}

func (a *ActivityService) GetActivityByID(id int32) (response models.Activity, err error) {
	log.Println("receive get activity id")
	activity, err := a.Connector.GetActivityByID(context.Background(), id)
	if err != nil {
		return models.Activity{}, err
	}
	log.Println("get activity success")
	response = utils.ConvertActivity(activity)
	return response, nil
}

func (a *ActivityService) GetListActivity() (response []models.Activity, err error) {
	log.Println("receive get list activity")
	activity, err := a.Connector.GetListActivity(context.Background())
	if err != nil {
		return []models.Activity{}, err
	}
	log.Println("get list activity success")
	for _ ,val := response {
		var item utils.ConvertActivity(activity)
		reponse = append(response,item)
	}
	return response, nil

}
