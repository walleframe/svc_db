package svc_db

import (
	"github.com/jmoiron/sqlx"
	"github.com/walleframe/walle/app/bootstrap"
)

// 数据库服务
var svc *DatabaseService = NewService()

func init() {
	// 注册数据库服务
	bootstrap.RegisterServiceByPriority(30, svc)
}

// RegisterDB 注册数据库相关接口
var RegisterDB func(driver, db, tblName string, swapDB func(db *sqlx.DB) error) = svc.RegisterDB

// SyncTableColumnsAndIndex 服务启动自动同步表字段级索引信息
var SyncTableColumnsAndIndex = true
