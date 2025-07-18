package svc

import (
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	c  config.Config
	DB *gorm.DB

	ProductCommentModel       model.ProductCommentModel
	ProductCommentReplayModel model.ProductCommentReplayModel
	ProductOperateLogModel    model.ProductOperateLogModel
	ProductVertifyRecordModel model.ProductVertifyRecordModel
	ProductCollectModel       model.ProductCollectModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		panic(err)
	}

	logx.Debug("mysql已连接")
	query.SetDefault(DB)

	ProductCommentModel := model.NewProductCommentModel(c.Mongo.Datasource, c.Mongo.Db, "product_comment")
	ProductCommentReplayModel := model.NewProductCommentReplayModel(c.Mongo.Datasource, c.Mongo.Db, "product_comment_replay")
	ProductOperateLogModel := model.NewProductOperateLogModel(c.Mongo.Datasource, c.Mongo.Db, "product_operate_log")
	ProductVertifyRecordModel := model.NewProductVertifyRecordModel(c.Mongo.Datasource, c.Mongo.Db, "product_vertify_record")
	ProductCollectModel := model.NewProductCollectModel(c.Mongo.Datasource, c.Mongo.Db, "product_collect")
	return &ServiceContext{
		c:  c,
		DB: DB,

		ProductCommentModel:       ProductCommentModel,
		ProductCommentReplayModel: ProductCommentReplayModel,
		ProductOperateLogModel:    ProductOperateLogModel,
		ProductVertifyRecordModel: ProductVertifyRecordModel,
		ProductCollectModel:       ProductCollectModel,
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
