package mysqlcfg

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// redis config
//
//go:generate gogen cfggen -n Config -o gen_configs.go
func generateMysqlConfig() interface{} {
	// root:123456@tcp(127.0.0.1:3306)/test2024?charset=utf8mb4,utf8&collation=utf8mb4_unicode_ci
	return map[string]interface{}{
		"DB":        "",
		"User":      "root",
		"Pass":      "123456",
		"Addr":      "127.0.0.1:3306",
		"Charset":   "utf8mb4,utf8",
		"Collation": "utf8mb4_unicode_ci",
	}
}

func NewMysqlClient(cfg *Config) (*sqlx.DB, error) {
	return sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&collation=%s",
		cfg.User, cfg.Pass,
		cfg.Addr, cfg.DB,
		cfg.Charset, cfg.Collation,
	))
}
