package model

type User struct {
	Id       string `json:"id" db:"Id"`
	UserName string `json:"userName" db:"UserName"`
	Name     string `json:"name" db:"Name"`
	//PassWord string `json:"passWord" db:"PassWord"`
	//Email       string `json:"email" db:"Email"`
	//Gender      string `json:"gender" db:"Gender"`
	//Age         int    `json:"age" db:"Age"`
	//Interest    string `json:"interest" db:"Interest"`
	//PhoneNumber int64  `json:"phoneNumber" db:"PhoneNumber"`
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

type UserInfo struct {
	Id          int64  `json:"id" db:"Id"`
	UserName    string `json:"userName" db:"UserName"`
	Name        string `json:"name" db:"Name"`
	PassWord    string `json:"passWord" db:"PassWord"`
	Email       string `json:"email" db:"Email"`
	Gender      string `json:"gender" db:"Gender"`
	Age         int    `json:"age" db:"Age"`
	Interest    string `json:"interest" db:"Interest"`
	PhoneNumber int64  `json:"phoneNumber" db:"PhoneNumber"`
}
