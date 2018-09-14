package models

import (
	"time"

	"github.com/MDGSF/Blog/u"
)

// TUser table UserInfo
type TUser struct {
	ID          uint64 `gorm:"primary_key"`
	CreatedAt   time.Time
	UserName    string `gorm:"type:varchar(255);unique_index"`
	PassWord    string `gorm:"type:varchar(255)"`
	NickName    string `gorm:"type:varchar(255)"`
	Email       string `gorm:"type:varchar(255)"`
	PhoneNumber string `gorm:"type:varchar(255)"`
}

// TableName TUser table name
func (TUser) TableName() string {
	return "UserInfo"
}

// Create create a new user.
func (user *TUser) Create() error {
	if err := gDB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Query query one user info.
func (user *TUser) Query() error {
	if err := gDB.Where(user).First(user).Error; err != nil {
		return err
	}
	return nil
}

// Update update one user info.
func (user *TUser) Update() {
	gDB.Save(user)
}

// Delete delete one user.
func (user *TUser) Delete() {
	gDB.Delete(user)
}

// GetUserSalt return a user salt token
func GetUserSalt() string {
	return u.GetRandomString(8)
}
