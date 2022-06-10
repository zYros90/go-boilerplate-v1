package config

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/zYros90/pkg/testutils"
)

func TestReadConfig(t *testing.T) {
	rootPath, err := testutils.GetRootPath()
	if err != nil {
		t.Errorf("error getting root path: %v", err)
	}
	correctConfig := &Config{
		Develop:  true,
		LogLevel: "debug",
		Server: Server{
			Host:         "127.0.0.0",
			Port:         9090,
			JWTSecret:    "thisismysecretkey",
			AllowOrigins: []string{"http://localhost:3000"},
		},
		PG: PG{
			Host:   "localhost",
			Port:   5432,
			DBName: "postgres",
			SSL:    "disable",
			Auth: Auth{
				User: "postgres",
				Pass: "postgres",
			},
		},
	}

	correctMergedConfig := &Config{
		Develop:  true,
		LogLevel: "debug",
		Server: Server{
			Host:         "127.0.0.0",
			Port:         9091,
			JWTSecret:    "thisismysecretkey",
			AllowOrigins: []string{"http://localhost:3000"},
		},
		PG: PG{
			Host:   "localhost",
			Port:   5432,
			DBName: "postgres",
			SSL:    "disable",
			Auth: Auth{
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
		want    *Config
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig(tt.args.basePath, tt.args.overwritePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
