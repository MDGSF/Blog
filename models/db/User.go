package db

import "time"

// TUser table UserInfo
type TUser struct {
	ID          uint64 `gorm:"primary_key"`
	CreatedAt   time.Time
	UserName    string `gorm:"type:varchar(255);unique_index"`
	PassWord    string `gorm:"type:varchar(255)"`
	NickName    string `gorm:"type:varchar(255)"`
	Email       string `gorm:"type:varchar(255);unique_index"`
	PhoneNumber string `gorm:"type:varchar(255);unique_index"`
}

// TableName TUser table name
func (TUser) TableName() string {
	return "UserInfo"
}

// Create create a new user.
func (user *TUser) Create() {
	gDB.Create(user)
}

// Query query one user info.
func (user *TUser) Query() {
	gDB.Where(user).First(user)
}

// Update update one user info.
func (user *TUser) Update() {
	gDB.Save(user)
}

// Delete delete one user.
func (user *TUser) Delete() {
	gDB.Delete(user)
}
