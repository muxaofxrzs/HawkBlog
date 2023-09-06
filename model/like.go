package model

type CommentLike struct {
	UserId    int64 `json:"user_id"`
	ArticleId int64 `json:"article_id"`
	CommentId int64 `json:"comment_id"`
	Status    int64 `json:"status"`
}
