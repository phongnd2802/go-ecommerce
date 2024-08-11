package initializes

import (
	"github.com/phongnd2802/go-ecommerce/global"
	"github.com/phongnd2802/go-ecommerce/pkg/logger"
)

func initLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}