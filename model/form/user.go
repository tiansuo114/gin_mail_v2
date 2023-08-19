package form

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             uint   `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeleteAt       int64  `json:"delete_at,omitempty" bson:"delete_at,omitempty"`
	UserName       string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Email          string `json:"email,omitempty" bson:"email,omitempty"`
	PasswordDigest string `json:"password_digest,omitempty" bson:"password_digest,omitempty"`
	NickName       string `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
	Status         string `json:"status,omitempty" bson:"status,omitempty"`
	Avatar         string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Money          string `json:"money,omitempty" bson:"money,omitempty"`
}

const (
	PassWordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarUrl 头像地址
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}
