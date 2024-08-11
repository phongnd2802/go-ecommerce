package initializes

import (
	"fmt"

	"github.com/phongnd2802/go-ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	loadConfig()
	initLogger()
	global.Logger.Info("Config Log OK!", zap.String("ok", "success"))
	initMysql()

	router := initRouter()
	router.Run(fmt.Sprintf(":%v", global.Config.Server.Port))
}