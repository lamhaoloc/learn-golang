package restaurantStorage

import (
	restaurantModel "awesomeProject/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	err := s.db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}
