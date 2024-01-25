package svc_db

import "github.com/jmoiron/sqlx"

// RegisterDB 注册数据库相关接口
func RegisterDB(driver, db, tblName string, swapDB func(db *sqlx.DB) error) {

}

// SyncTableColumnsAndIndex 服务启动自动同步表字段级索引信息
var SyncTableColumnsAndIndex = true
