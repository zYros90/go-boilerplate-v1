package config_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/pkg/testutils"
)

func TestReadConfig(t *testing.T) {
	rootPath, err := testutils.GetRootPath()
	if err != nil {
		t.Errorf("error getting root path: %v", err)
	}

	correctConfig := &config.Config{
		Develop:  true,
		LogLevel: "debug",
		Server: config.Server{
			Host:         "127.0.0.0",
			Port:         9090,
			JWTSecret:    "thisismysecretkey",
			AllowOrigins: []string{"http://localhost:3000"},
		},
		PG: config.PG{
			Host:   "localhost",
			Port:   5432,
			DBName: "postgres",
			SSL:    "disable",
			Auth: config.Auth{
				User: "postgres",
				Pass: "postgres",
			},
		},
	}

	correctMergedConfig := &config.Config{
		Develop:  true,
		LogLevel: "debug",
		Server: config.Server{
			Host:         "127.0.0.0",
			Port:         9091,
			JWTSecret:    "thisismysecretkey",
			AllowOrigins: []string{"http://localhost:3000"},
		},
		PG: config.PG{
			Host:   "localhost",
			Port:   5432,
			DBName: "postgres",
			SSL:    "disable",
			Auth: config.Auth{
				User: "postgres",
				Pass: "postgres",
			},
		},
	}

	type args struct {
		basePath      string
		overwritePath string
	}

	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			"read base config",
			args{
				basePath:      filepath.Join(rootPath, "app/config/test.yaml"),
				overwritePath: "",
			},
			correctConfig,
			false,
		},
		{
			"read config and merge",
			args{
				basePath:      filepath.Join(rootPath, "app/config/test.yaml"),
				overwritePath: filepath.Join(rootPath, "app/config/test_merge.yaml"),
			},
			correctMergedConfig,
			false,
		},
		{
			"read config with wrong config path",
			args{
				basePath:      filepath.Join(rootPath, "app/config/random_config.yaml"),
				overwritePath: "",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.ReadConfig(tt.args.basePath, tt.args.overwritePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
