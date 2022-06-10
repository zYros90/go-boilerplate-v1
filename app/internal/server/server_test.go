package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/pkg/logger"
	"go.uber.org/zap"
)

func Test_newEcho(t *testing.T) {
	type args struct {
		config *config.Config
		logger *zap.Logger
	}
	log, err := logger.NewLogger("debug", true, true, true)
	if err != nil {
		t.Error(err)
	}

	correctConfig := &config.Config{
		Develop:  true,
		LogLevel: "debug",
		Server: config.Server{
			Host:         "127.0.0.0",
			Port:         9090,
			AllowOrigins: []string{"http://localhost:3000"},
		},
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"new echo",
			args{
				config: correctConfig,
				logger: log.Logger,
			},
			false,
		},
		{
			"new echo",
			args{
				config: correctConfig,
				logger: log.Logger,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newEcho(tt.args.config, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("newEcho() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			notAcceptedEndTime := time.Now().Add(1 * time.Second)
			go func() {
				err = got.Start(fmt.Sprintf(":%d", correctConfig.Server.Port))
				if err != nil {
					if time.Now().Before(notAcceptedEndTime) {
						t.Errorf("newEcho() error starting echo = %v", err)
					}

				}
			}()
			time.Sleep(2 * time.Second)
			err = got.Shutdown(context.Background())
			if err != nil {
				t.Errorf("newEcho() shutdown echo %v", err)
			}
		})
	}
}
