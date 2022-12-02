package userStore

import (
	"awesomeProject/common"
	userModel "awesomeProject/module/user/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*userModel.User, error) {
	db := s.db.Table(userModel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user userModel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
