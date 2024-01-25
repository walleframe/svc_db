package svc_db

import (
	"github.com/walleframe/walle/app"
)

// DatabaseService mysql链接管理
type DatabaseService struct {
	app.NoopService

	cfgName string
}

func NewService() *DatabaseService {
	return &DatabaseService{}
}

func (svc *DatabaseService) Name() string {
	return "database-svc"
}

func (svc *DatabaseService) Init(s app.Stoper) (err error) {
	return
}

func (svc *DatabaseService) Stop() {
}
