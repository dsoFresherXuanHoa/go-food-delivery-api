package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type AreaStorage interface {
	CreateArea(ctx context.Context, area *models.AreaCreatable) (*uint, error)
	ReadAreaById(ctx context.Context, areaId uint) (*models.Area, error)
}

type areaBusiness struct {
	storage AreaStorage
}

func NewAreaBusiness(storage AreaStorage) *areaBusiness {
	return &areaBusiness{storage: storage}
}

func (business *areaBusiness) CreateArea(ctx context.Context, area *models.AreaCreatable) (*uint, error) {
	if areaId, err := business.storage.CreateArea(ctx, area); err != nil {
		fmt.Println("Error while save Area in service: " + err.Error())
		return nil, err
	} else {
		return areaId, nil
	}
}

func (business *areaBusiness) ReadAreaById(ctx context.Context, areaId uint) (*models.Area, error) {
	if area, err := business.storage.ReadAreaById(ctx, areaId); err != nil {
		fmt.Println("Error while read Area by id in service: " + err.Error())
		return nil, err
	} else {
		return area, nil
	}
}
