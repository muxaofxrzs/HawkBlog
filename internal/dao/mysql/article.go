package mysql

import (
	"context"
	"hawk/model"
)

func AddEssay(article model.Article) error {
	sql := "insert into article(ArticleId,UserId,Title,Content,Label) values (?,?,?,?,?)"
	_, err := GlobalConn.ExecCtx(context.Background(), sql, article.ArticleId, article.UserId, article.Title, article.Content, article.Label)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEssay(artileId int64) error {
	sql := "DELETE FROM article WHERE ArticleId = ?"
	_, err := GlobalConn.ExecCtx(context.Background(), sql, artileId)
	if err != nil {
		return err
	}
	return nil
}

func FindEssay(title string) (err error, article model.Article) {
	sql := "SELECT * FROM article WHERE Title =?"

	err = GlobalConn.QueryRowCtx(context.Background(), &article, sql, title)
	return err, article
}
func UpdataEssay(articleId int64, content string) error {
	sql := "UPDATE article SET Content =? WHERE ArticleId =?"
	_, err := GlobalConn.ExecCtx(context.Background(), sql, content, articleId)
	return err
}
