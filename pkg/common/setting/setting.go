package setting

import (
	"fmt"
	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	"os"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/spf13/viper"
)

var gConfig = unsafe.Pointer(&Config{})

type Config struct {
	RunMode  string         `yaml:"RunMode"`
	Addr     string         `yaml:"Addr"`
	Database DatabaseConfig `yaml:"Database"`
	Jwt      JwtConfig      `yaml:"Jwt"`
	Log      log.Options    `yaml:"Log"`
}

type JwtConfig struct {
	Key         string `yaml:"Key"`
	IdentityKey string `yaml:"IdentityKey"`
}

type DatabaseConfig struct {
	Host            string
	Username        string
	Password        string
	DatabaseName    string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	LoggerLevel     int
}

func C() *Config {
	return (*Config)(atomic.LoadPointer(&gConfig))
}

func InitConfig(configFile string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()           // 读取匹配的环境变量
	viper.SetEnvPrefix("GOSERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.Unmarshal(C()); err != nil {
		fmt.Fprintln(os.Stderr, "unmarshal error")
	}

}
