package cmd

import (
	"path/filepath"
	"testing"

	"github.com/zYros90/pkg/testutils"
)

func TestRunExecute(t *testing.T) {
	rootPath, err := testutils.GetRootPath()
	if err != nil {
		t.Errorf("error getting root path: %v", err)
	}
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			"run",
			[]string{"run", "--config", filepath.Join(rootPath, "app/config/base.yaml"), "--testcmd"},
			false,
		},
		{
			"run",
			[]string{"run", "--mergeconfig", filepath.Join(rootPath, "app/config/base.yaml"), "--testcmd"},
			false,
		},
		{
			"run",
			[]string{"run", "--config", filepath.Join(rootPath, "app/config/xyz.yaml"), "--testcmd"},
			true,
		},
		{
			"run",
			[]string{"run", "--mergeconfig", filepath.Join(rootPath, "app/config/xyz.yaml"), "--testcmd"},
			true,
		},
	}
	for _, tt := range tests {
		rootCmd.SetArgs(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			err := rootCmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("rootCmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
