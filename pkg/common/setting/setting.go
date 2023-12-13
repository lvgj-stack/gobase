package setting

import (
	"fmt"
	"os"
	"sync/atomic"
	"unsafe"

	"github.com/Mr-LvGJ/jota/log"
	"github.com/Mr-LvGJ/jota/models"

	"github.com/spf13/viper"
)

var gConfig = unsafe.Pointer(&Config{})

type Config struct {
	RunMode   string                 `yaml:"RunMode"  default:"debug"`
	Addr      string                 `yaml:"Addr"     default:"addr"`
	Database  *models.DatabaseConfig `yaml:"Database"`
	Jwt       JwtConfig              `yaml:"Jwt"`
	Log       *log.Config            `yaml:"Log"`
	AccessLog *log.Config            `yaml:"AccessLog"`
}

type JwtConfig struct {
	Key         string `yaml:"Key"`
	IdentityKey string `yaml:"IdentityKey"`
}

func C() *Config {
	return (*Config)(atomic.LoadPointer(&gConfig))
}

func InitConfig(configFile string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // 读取匹配的环境变量
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.Unmarshal(C()); err != nil {
		fmt.Fprintln(os.Stderr, "unmarshal error")
	}

}
