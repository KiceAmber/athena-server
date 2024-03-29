// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleDao is the data access object for table article.
type ArticleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ArticleColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleColumns defines and stores column names for table article.
type ArticleColumns struct {
	Id          string // 文章ID
	Title       string // 文章标题
	Content     string // 文章内容
	Cover       string // 文章图片封面
	AuthorId    string // 文章作者ID
	CategoryId  string // 文章分类ID
	Description string // 文章概述
	IsVisible   string // 文章是否可见 1-可见 0-不可见
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// articleColumns holds the columns for table article.
var articleColumns = ArticleColumns{
	Id:          "id",
	Title:       "title",
	Content:     "content",
	Cover:       "cover",
	AuthorId:    "author_id",
	CategoryId:  "category_id",
	Description: "description",
	IsVisible:   "is_visible",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewArticleDao creates and returns a new DAO object for table data access.
func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		group:   "default",
		table:   "article",
		columns: articleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleDao) Columns() ArticleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
