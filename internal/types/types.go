// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

type CaptchaReq struct {
	Email string `json:"email"`
}

type RegisterReq struct {
	Code        string `json:"code"`
	UserName    string `json:"userName"`
	PassWord    string `json:"passWord"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Age         int64  `json:"age"`
	Interest    string `json:"interest"`
	PhoneNumber int64  `json:"phoneNumber"`
}

type UpdateUserInformationReq struct {
	UserName    string `json:"userName"`
	PassWord    string `json:"passWord"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Age         int64  `json:"age"`
	Interest    string `json:"interest"`
	PhoneNumber int64  `json:"phoneNumber"`
}

type HttpCode struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type AddArticleReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Label   string `json:"label"`
}

type AddDraftReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Label   string `json:"label"`
}

type UpdateArticleReq struct {
	ArticleId int64  `json:"articleId"`
	Content   string `json:"content"`
}

type ExamineReq struct {
	ArticleId int64 `json:"articleId"`
}

type ExamineArticlesReq struct {
	ArticleId int64 `json:"articleId"`
	PageSize  int64 `json:"pageSize"`
}

type DeleteReq struct {
	ArticleId int64 `json:"articleId"`
}

type LikeReq struct {
	ArticleId int64 `json:"articleId"`
}

type ListTitleReq struct {
	Title     string `json:"title"`
	ArticleId int64  `json:"articleId"`
	PageSize  int    `json:"pageSize"`
}

type HttpCodeResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateCommentReq struct {
	ArticleId int64  `json:"articleId"`
	Comment   string `json:"comment"`
}

type DeleteCommentReq struct {
	ArticleId int64 `json:"articleId"`
	CommentId int64 `json:"commentId"`
}

type UpdateCommentReq struct {
	ArticleId int64  `json:"articleId"`
	CommentId int64  `json:"commentId"`
	Comment   string `json:"comment"`
}

type GetAllCommentReq struct {
	ArticleId int64 `json:"articleId"`
}

type PostCommentLikeReq struct {
	ArticleId int64 `json:"articleId"`
	CommentId int64 `json:"commentId"`
}

type CommentToCommentReq struct {
	CommmentId int64  `json:"commmentId"`
	Comment    string `json:"comment"`
}
