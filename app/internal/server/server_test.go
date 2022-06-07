package server

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"go.uber.org/zap"
)

func Test_newEcho(t *testing.T) {
	type args struct {
		config *config.Config
		logger *zap.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *echo.Echo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newEcho(tt.args.config, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("newEcho() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEcho() = %v, want %v", got, tt.want)
			}
		})
	}
}
