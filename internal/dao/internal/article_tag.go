// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleTagDao is the data access object for table article_tag.
type ArticleTagDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ArticleTagColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleTagColumns defines and stores column names for table article_tag.
type ArticleTagColumns struct {
	Id        string // 文章标签关联ID
	ArticleId string // 文章ID
	TagId     string // 标签ID
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

// articleTagColumns holds the columns for table article_tag.
var articleTagColumns = ArticleTagColumns{
	Id:        "id",
	ArticleId: "article_id",
	TagId:     "tag_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewArticleTagDao creates and returns a new DAO object for table data access.
func NewArticleTagDao() *ArticleTagDao {
	return &ArticleTagDao{
		group:   "default",
		table:   "article_tag",
		columns: articleTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleTagDao) Columns() ArticleTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
