package userModel

import (
	"awesomeProject/common"
	"errors"
)

const EntityName = "User"

type User struct {
	common.BaseModel `json:",inline"`
	Email            string `json:"email" gorm:"column:email"`
	Password         string `json:"-" gorm:"column:password"`
	Salt             string `json:"-" gorm:"column:salt"`
	LastName         string `json:"last_name" gorm:"column:last_name"`
	FirstName        string `json:"first_name" gorm:"column:first_name"`
	Phone            string `json:"phone" gorm:"column:phone"`
	Role             string `json:"role" gorm:"column:role"`
	Avatar           string `json:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

// Create User Model
type CreateUserDTO struct {
	common.BaseModel `json:",inline"`
	Email            string        `json:"email" gorm:"column:email"`
	Password         string        `json:"-" gorm:"column:password"`
	Salt             string        `json:"-" gorm:"column:salt"`
	LastName         string        `json:"last_name" gorm:"column:last_name"`
	FirstName        string        `json:"first_name" gorm:"column:first_name"`
	Role             common.Role   `json:"role" gorm:"column:role"`
	Avatar           *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (CreateUserDTO) TableName() string {
	return User{}.TableName()
}
func (u *CreateUserDTO) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLoginDTO struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid")

	ErrEmailExisted = common.NewCustomError(
		errors.New("username has already existed"),
		"username has already existed",
		"ErrEmailExisted")
)
