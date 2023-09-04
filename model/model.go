package model

type User struct {
	Id       string `json:"id" db:"Id"`
	UserName string `json:"userName" db:"UserName"`
	Name     string `json:"name" db:"Name"`
}

type Article struct {
	ArticleId  int64  `json:"articleId" db:"ArticleId"`
	UserId     int64  `json:"userId" db:"UserId"`
	Title      string `json:"title" db:"Title"`
	Content    string `json:"content" db:"Content"`
	Label      string `json:"label" db:"Label"`
	StartTime  string `json:"startTime" db:"StartTime"`
	UpdataTime string `json:"updataTime" db:"UpdateTime"`
	Delete     int    `json:"delete" db:"Delete"`
	LikeCount  int    `json:"likeCount" db:"LikeCount"`
	PageViews  int    `json:"pageViews" db:"PageViews"`
}
