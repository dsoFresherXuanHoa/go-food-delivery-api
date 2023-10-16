package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateTable(ctx context.Context, table *models.TableCreatable) (*uint, error) {
	if err := s.db.Table(models.TableCreatable{}.GetTableName()).Create(&table).Error; err != nil {
		fmt.Println("Error while create table in repository: " + err.Error())
		return nil, err
	} else {
		return &table.ID, nil
	}
}

func (s *sqlStorage) ReadTableByEmployeeId(ctx context.Context, employeeId uint) (models.Tables, error) {
	var tables models.Tables
	if err := s.db.Where("employee_id = ?", employeeId).Find(&tables).Error; err != nil {
		fmt.Println("Error while find table by id in repository: " + err.Error())
		return nil, err
	} else if len(tables) == 0 {
		fmt.Println("No record found while find table by employeeId in repository: " + err.Error())
		return nil, err
	}
	return tables, nil
}

func (s *sqlStorage) ReadTableByEmployeeIdAndStatus(ctx context.Context, employeeId uint, status bool) (models.Tables, error) {
	var tables models.Tables
	if err := s.db.Where("employee_id = ?", employeeId).Where("available = ?", status).Find(&tables).Error; err != nil {
		fmt.Println("Error while find table by id in repository: " + err.Error())
		return nil, err
	} else if len(tables) == 0 {
		fmt.Println("No record found while find table by employeeId and status in repository: " + err.Error())
		return nil, err
	}
	return tables, nil
}
