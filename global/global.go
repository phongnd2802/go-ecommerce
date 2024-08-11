package global

import (
	"database/sql"

	"github.com/phongnd2802/go-ecommerce/pkg/logger"
	"github.com/phongnd2802/go-ecommerce/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Db *sql.DB
)