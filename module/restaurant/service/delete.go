package restaurantService

import (
	"awesomeProject/common"
	"context"

	restaurantModel "awesomeProject/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantModel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}

type deleteRestaurantService struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantService(store DeleteRestaurantStore) *deleteRestaurantService {
	return &deleteRestaurantService{store: store}
}

func (biz *deleteRestaurantService) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(restaurantModel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantModel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantModel.EntityName, nil)
	}

	return nil
}
