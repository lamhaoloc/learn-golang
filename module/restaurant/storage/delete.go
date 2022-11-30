package restaurantStorage

import (
	"awesomeProject/common"
	restaurantModel "awesomeProject/module/restaurant/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantModel.Restaurant{}.TableName()).
		Where("id = ?").
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
