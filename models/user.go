package models

import "time"

type User struct {
  ID        int       `json:"id"`
  Fullname  string    `json:"fullname" gorm:"type: varchar(255)"`
  Username	string	  `json:"username" gorm:"type: varchar(255)"`
  Email     string    `json:"email" gorm:"type: varchar(255)"`
  Password  string    `json:"-" gorm:"type: varchar(255)"`
  ListAs  	string    `json:"listAs" gorm:"type: varchar(255)"`
  Gender  	string    `json:"gender" gorm:"type: varchar(255)"`
  Phone  	  string    `json:"phone" gorm:"type: varchar(255)"`
  Address  	string    `json:"address" gorm:"type: varchar(255)"`
  CreatedAt time.Time `json:"-"`
  UpdatedAt time.Time `json:"-"`
}

type UserResponse struct {
  ID        int       `json:"id"`
  Fullname  string    `json:"fullname"`
  Username  string    `json:"username"`
  Email     string    `json:"email"`
  ListAs  	string    `json:"listAs"`
}

func (UserResponse) TableName() string {
	return "users"
}