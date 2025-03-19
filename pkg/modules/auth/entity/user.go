package entity

import "github.com/indrabay/helloibe-api/utils"

const UserTableName = "users"

type User struct {
	ID        int64  `json:"id" gorm:"column:id"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	Name      string `json:"name" gorm:"column:name"`
	Role      int    `json:"role_id" gorm:"column:role_id"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

type CompleteUser struct {
	ID        int64  `json:"id" gorm:"column:id"`
	Username  string `json:"username" gorm:"column:username"`
	Name      string `json:"name" gorm:"column:name"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
	Role      Role   `json:"role"`
}

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (param *LoginParam) Validate() error {
	if param.Username == "" || param.Password == "" {
		return utils.ErrUsernamePasswordRequired
	}

	return nil
}

type RegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
}

func (param *RegisterParam) Validate() error {
	if param.Username == "" || param.Password == "" || param.Name == "" || param.Role == 0 {
		return utils.ErrCreateUserRequiredParam
	}

	return nil
}

func (User) TableName() string {
	return UserTableName
}
