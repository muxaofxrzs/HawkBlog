package model

import "time"

const (
	ArticleCommentTime  = "hawk:comment:time:articleId-"
	ArticleCommentScore = "hawk:comment:score:articleId-"
	CommentCommentTime  = "hawk:comment:time:commentId-"
	CommentCommentScore = "hawk:comment:score:articleId-"
)

type ArticleComment struct {
	UserId     int64     `db:"user_id"`                        //用户ID
	CommentId  int64     `db:"comment_id"`                     //评论ID
	Comment    string    `db:"comment"`                        //评论内容
	CreateTime time.Time `db:"create_time"`                    //创建时间
	UpdateTime time.Time `db:"update_time"`                    //更新时间
	Like       int64     `db:"like"`                           //点赞数
	Status     int64     `json:"status,omitempty" db:"status"` //是否被删除
}
