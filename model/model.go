package model

import "time"

type User struct {
	Id       string `json:"id" db:"Id"`
	UserName string `json:"userName" db:"UserName"`
	Name     string `json:"name" db:"Name"`
}

type Article struct {
	ArticleId  int64     `json:"articleId" db:"ArticleId"`
	UserId     int64     `json:"userId" db:"UserId"`
	UserName   string    `json:"user_name" db:"UserName"`
	Title      string    `json:"title" db:"Title"`
	Content    string    `json:"content" db:"Content"`
	Label      string    `json:"label" db:"Label"`
	StartTime  time.Time `json:"startTime" db:"StartTime"`
	UpdataTime time.Time `json:"updataTime" db:"UpdateTime"`
	Status     int       `json:"status" db:"Status"`
	LikeCount  int       `json:"likeCount" db:"LikeCount"`
	PageViews  int       `json:"pageViews" db:"PageViews"`
	Heat       float64   `json:"heat" db:"Heat"`
}
