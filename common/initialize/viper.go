package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go_im/common/global"
)

func LoadConfig() {

	viper.AddConfigPath("./")
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error resources file: %w \n", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct %w \n", err))
	}
}
