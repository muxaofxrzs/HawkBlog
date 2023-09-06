package mysql

import (
	"context"
	"fmt"
	"hawk/model"
	"math"
	"time"
)

func AddEssay(article model.Article) error {
	sql := "insert into article(ArticleId, UserId, Title, Content, Label, StartTime, UpdateTime, Status, UserName,Heat) values (?,?,?,?,?,?,?,?,?,?)"
	_, err := GlobalConn.ExecCtx(context.Background(), sql, article.ArticleId, article.UserId, article.Title, article.Content, article.Label, article.StartTime, article.UpdataTime, article.Status, article.UserName, article.Heat)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEssay(artileId int64) error {
	delete := 2
	sql := "UPDATE article SET Status =? WHERE ArticleId =?"
	_, err := GlobalConn.ExecCtx(context.Background(), sql, delete, artileId)
	return err
}

// 传进来文章题目后查到文章的内容。并且进行文章的浏览量更新。
func FindEssay(articleId int64) (err error, article model.Article) {

	sql := "SELECT * FROM article WHERE ArticleId = ? AND Status = 1"

	err = GlobalConn.QueryRowCtx(context.Background(), &article, sql, articleId)
	if err != nil {
		return err, article
	}
	if article.ArticleId != 0 {
		article.PageViews++
		err = UpdatePageViews(articleId, article.PageViews)
		if err != nil {
			return err, article
		}
	}
	return err, article
}

// 用户进行更新文章
func UpdataEssay(article model.Article) error {
	sql := "UPDATE article SET Content =? WHERE ArticleId =?"
	_, err := GlobalConn.ExecCtx(context.Background(), sql, article.Content, article.ArticleId)
	return err
}

// 进行浏览的时候，浏览量进行加一
func UpdatePageViews(articleId int64, pageViews int) error {
	sql := "UPDATE article SET PageViews = ? WHERE ArticleId = ?"
	_, err := GlobalConn.ExecCtx(context.Background(), sql, pageViews, articleId)
	return err
}

// 用户进行点赞
func LikeArticleInMySQL(articleID int64) (int, error) {
	var article model.Article
	// 查询文章信息
	sql := "SELECT * FROM article WHERE ArticleId = ?"
	err := GlobalConn.QueryRowCtx(context.Background(), &article, sql, articleID)
	if err != nil {
		return 0, err
	}
	article.LikeCount++
	// 更新点赞量
	_, err = GlobalConn.Exec("UPDATE article SET LikeCount = ? WHERE ArticleId = ?", article.LikeCount, articleID)
	return article.LikeCount, err
}

//获取浏览量最高的

// 起个定时器，每天更新热度
func InitHeat() {
	t := time.NewTicker(1 * time.Hour)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			UpdateHeat()
		}
	}
}

// UpdateHeat 利用热度算法，计算热度，更新数据库的热度。
func UpdateHeat() {
	//获取当前的时间
	//定义衰减系数
	decayFactor := 0.5
	fmt.Println(decayFactor)
	currentTime := time.Now()
	var articles []model.Article
	var article model.Article
	//查询所有的文章
	sql := "SELECT * FROM article"
	err := GlobalConn.QueryRowsCtx(context.Background(), &articles, sql)
	if err != nil {
		fmt.Println("获取热度时，查询文章出错。")
		fmt.Println(err)
		return
	}

	for _, article = range articles {
		//计算发布时间与当前时间的时间差，以小时为单位
		hoursSincePublish := currentTime.Sub(article.StartTime).Hours()
		//计算时间复杂度
		article.Heat = float64(article.PageViews) + float64(article.LikeCount)
		article.Heat /= math.Pow(hoursSincePublish+2, decayFactor)
		sql1 := "UPDATE article SET Heat =? WHERE ArticleId = ?"
		_, err := GlobalConn.ExecCtx(context.Background(), sql1, article.Heat, article.ArticleId)
		if err != nil {
			fmt.Println("热度更新失败")
			return
		}
	}
	return
}

// 获取热度最高的前六篇文章
func GetTopSixArticles() ([]model.Article, error) {
	var articles []model.Article
	sql := "SELECT * FROM article WHERE Status =1 ORDER BY Heat DESC LIMIT 6"
	err := GlobalConn.QueryRowsCtx(context.Background(), &articles, sql)
	if err != nil {
		fmt.Println("获取失败")
		fmt.Println(err)
		return nil, err
	}
	return articles, nil
}

// 当传进来的第一个id为0的时候，获取最新的id
func GetTopOneTitle(Title string) (articleId int64) {
	title := "%" + Title + "%"
	var article model.Article
	sql := "SELECT * FROM article WHERE Status =1 AND Title LIKE ? ORDER BY ArticleId DESC LIMIT 1"
	err := GlobalConn.QueryRowCtx(context.Background(), &article, sql, title)
	if err != nil {
		fmt.Println("查询第一个id失败")
		return
	}
	return article.ArticleId
}

// 用户在输入题目的时候，根据模糊查询返回题目，并且会有游标分页
func GetNextTitleId(nextId int64, pageSize int, Title string) ([]model.Article, error) {
	var articles []model.Article
	title := "%" + Title + "%"
	sql := "SELECT * FROM article WHERE Status = 1 AND ArticleId < ? AND Title LIKE ?  ORDER BY ArticleId DESC LIMIT ?"
	err := GlobalConn.QueryRowsCtx(context.Background(), &articles, sql, nextId, title, pageSize)
	if err != nil {
		return articles, err
	}
	return articles, nil
}
