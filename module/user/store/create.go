package userStore

import (
	"awesomeProject/common"
	"context"

	userModel "awesomeProject/module/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *userModel.CreateUserDTO) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
