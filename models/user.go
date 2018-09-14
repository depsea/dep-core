package models

import "time"

// UserID -
type UserID = uint

// UserStatus -
type UserStatus = uint

const (
	// UserStatusAble -
	UserStatusAble UserStatus = iota
)

// User -
type User struct {
	ID         UserID     `json:"id" sql:"id"`
	UserName   string     `json:"username" sql:"username"`
	NickName   string     `json:"nickname" sql:"nickname"`
	Password   string     `json:"password" sql:"password"`
	Status     UserStatus `json:"status" sql:"status"`
	CreateTime time.Timer `json:"create_time" sql:"create_time"`
	UpdateTime time.Timer `json:"update_time" sql:"update_time"`
}

// InsertUser -
func (u *User) InsertUser() (*User, error) {
	db, err := GetDB()

	defer db.Close()

	if err != nil {
		return nil, err
	}

	db.QueryRow("INSERT INTO users (username, nickname, password) VALUES ($1, $2, $3) RETUNING id", u.UserName, u.NickName, u.Password).Scan(u.ID)

	return u, nil
}
