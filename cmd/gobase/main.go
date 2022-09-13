package main

import (
	"flag"
	"github.com/Mr-LvGJ/gobase/cmd/gobase/bootstrap"
	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	"github.com/Mr-LvGJ/gobase/pkg/common/setting"
	"github.com/spf13/viper"
	"net/http"
)

var (
	configPath  = flag.String("config", "/etc/goserver/gobase.yaml", "The gobase config file default user home config")
	pprofAddr   = flag.String("pprof-addr", "", "The pprof addr")
	showVersion = flag.Bool("version", false, "")
)

func main() {
	flag.Parse()

	if pprofAddr != nil && len(*pprofAddr) > 0 {
		go func() {
			if err := http.ListenAndServe(*pprofAddr, nil); err != nil {
				log.Info("fail to start pprof", "err:", err)
			}
		}()
	}
	setting.InitConfig(*configPath)
	bootstrap.Run()
	log.Info("get addr", "addr", viper.Get("addr"))
}
