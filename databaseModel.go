package main

import (
	"time"
)

type PublicModel struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	deleted   int       `json:"deleted,omitempty"`
}

type Passages struct {
	Id          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Category    string    `json:"category"`
	Tag         string    `json:"tag"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	ImgLink     string    `json:"img_link"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

/*注册用户*/
type Users struct {
	Id       int    `json:"id,omitempty"`
	Sex      int    `json:"sex,omitempty"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	PublicModel
}

/*登录记录*/
type LoginToken struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	ExpireAt  int64     `json:"expire_at"`
	Valid     int       `json:"valid,omitempty"`
}
