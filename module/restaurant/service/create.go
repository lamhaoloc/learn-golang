package restaurantService

import (
	restaurantModel "awesomeProject/module/restaurant/model"
	"context"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantModel.RestaurantCreate) error
}

type createRestaurantService struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantService(store CreateRestaurantStore) *createRestaurantService {
	return &createRestaurantService{store: store}
}
func (service *createRestaurantService) CreateRestaurant(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := service.store.CreateRestaurant(ctx, data); err != nil {
		return err
	}

	return nil
}
