package svc_db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/walleframe/svc_db/mysqlcfg"
	"github.com/walleframe/walle/app"
	"github.com/walleframe/walle/zaplog"
	"go.uber.org/zap"
)

// DatabaseService mysql链接管理
type DatabaseService struct {
	app.NoopService

	// 数据库配置
	mysqlDB map[string]*mysqlDBLink
}

func NewService() *DatabaseService {
	return &DatabaseService{
		mysqlDB: map[string]*mysqlDBLink{},
	}
}

func (svc *DatabaseService) Name() string {
	return "database-svc"
}

func (svc *DatabaseService) Init(s app.Stoper) (err error) {
	for _, link := range svc.mysqlDB {
		err = link.initDB(s)
		if err != nil {
			return err
		}
	}
	return
}

func (svc *DatabaseService) Stop() {
}

func (svc *DatabaseService) RegisterDB(driver, db, tblName string, swapDB func(db *sqlx.DB) error) {
	switch driver {
	case "mysql":
		link, ok := svc.mysqlDB[db]
		if ok {
			link.add(tblName, swapDB)
			return
		}
		svc.mysqlDB[db] = newMysql(db, tblName, swapDB)
	default:
		panic(fmt.Sprintf("database driver %s not support, use for %s.%s ", driver, db, tblName))
	}
}

type mysqlDBLink struct {
	useTables   []string                  // 使用的表
	updateLinks []func(db *sqlx.DB) error // 对应表的更新链接
	cfg         *mysqlcfg.Config          // 配置信息
	log         *zaplog.Logger
}

func newMysql(db, tblName string, swapDB func(db *sqlx.DB) error) *mysqlDBLink {
	// 新建mysql配置
	cfg := mysqlcfg.NewConfig(fmt.Sprintf("mysql.%s", db))
	// 更新默认值
	cfg.DB = db

	return &mysqlDBLink{
		useTables:   []string{tblName},
		updateLinks: []func(db *sqlx.DB) error{swapDB},
		cfg:         cfg,
	}
}

func (x *mysqlDBLink) add(tblName string, swapDB func(db *sqlx.DB) error) {
	for _, v := range x.useTables {
		if tblName == v {
			panic(fmt.Sprintf("register mysql %s.%s repeated", x.cfg.DB, tblName))
		}
	}
	x.useTables = append(x.useTables, tblName)
	x.updateLinks = append(x.updateLinks, swapDB)
}

func (x *mysqlDBLink) initDB(s app.Stoper) (err error) {
	x.log = zaplog.GetFrameLogger().Named(fmt.Sprintf("mysql.%s", x.cfg.DB))
	log := x.log.New("initDB")
	// new db link
	db, err := mysqlcfg.NewMysqlClient(x.cfg)
	if err != nil {
		log.Error("init db failed", zap.Any("config", x.cfg), zap.Error(err))
		return err
	}

	// check db link
	err = db.Ping()
	if err != nil {
		log.Error("link db failed", zap.Any("config", x.cfg), zap.Error(err))
		return err
	}

	// update all table
	for _, swap := range x.updateLinks {
		err = swap(db)
		if err != nil {
			log.Error("swap db failed", zap.Any("config", x.cfg), zap.Error(err))
			return err
		}
	}

	// 监听mysql配置变化
	x.cfg.AddNotifyFunc(x.updateConfig)

	return
}
func (x *mysqlDBLink) updateConfig(cfg *mysqlcfg.Config) {
	log := x.log.New("updateConfig")
	// new db link
	db, err := mysqlcfg.NewMysqlClient(x.cfg)
	if err != nil {
		log.Error("new mysql client failed", zap.Any("cfg", x.cfg), zap.Error(err))
		return
	}

	// update all table
	for _, swap := range x.updateLinks {
		err = swap(db)
		if err != nil {
			log.Error("swap db failed", zap.Any("config", x.cfg), zap.Error(err))
		}
	}
}
