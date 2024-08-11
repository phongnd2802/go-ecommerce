package initializes

import (
	"fmt"
	"log"

	"github.com/phongnd2802/go-ecommerce/global"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}

}