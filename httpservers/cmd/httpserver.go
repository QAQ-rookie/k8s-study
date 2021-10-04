package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/klog"

	"httpserver.demo/cmd/options"
	"httpserver.demo/pkg/api"
)

var (
	httpConfig = &options.HTTPConfig{}
	httpCmd    = &cobra.Command{
		Use: "httpserver",
		RunE: func(cmd *cobra.Command, args []string) error {
			klog.Infoln("Start http server.")
			srv := &http.Server{
				Addr:    httpConfig.Addr,
				Handler: api.RegisterRouter(),
			}
			go func(srv *http.Server) {
				// service connections
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					klog.Errorf("listen: %s\n", err)
				}
			}(srv)
			WaitSignal()
			GraceShutdown(srv)
			return nil
		},
	}
)

func init() {
	httpCmd.PersistentFlags().StringVar(&httpConfig.Addr, "http-address", "0.0.0.0:8080", "--http-address=:8080")
}

func Execute() {
	if err := httpCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func WaitSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	signalName := <-sigs
	klog.Errorf("receive signal is %v\n", signalName)
}

func GraceShutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		klog.Fatalf("HTTP service shutdown failed, err is %v\n", err)
	}

	klog.Infoln("HTTP service is grace shutdown")
}
