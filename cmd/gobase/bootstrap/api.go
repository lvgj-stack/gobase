package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	"github.com/Mr-LvGJ/gobase/pkg/gobase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func Run() error {
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	gobase.LoadRouter(g)

	insecureServer := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}
	go func() {
		log.Info("Start to listening the incoming request on http address: %s", "addr", viper.GetString("addr"))
		go func() {
			if err := insecureServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Error("Listen: %s\n", "err", err)
			}
		}()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func() {
		if err := pingServer(ctx); err != nil {
			log.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	if err := insecureServer.Shutdown(ctx); err != nil {
		log.Error("Insecure server fored to shutdown")
		return err
	}
	return nil

}

func pingServer(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", viper.GetString("addr"))
	bind := strings.Split(viper.GetString("addr"), ":")[0]
	if bind == "" || bind == "0.0.0.0" {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(viper.GetString("addr"), ":")[1])
	}
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			return nil
		}
		log.Info("Wait for router, retry in 1 second.")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			return fmt.Errorf("can not ping server within the specified time interval")
		default:
		}
	}
}
