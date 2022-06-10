package cmd_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/zYros90/go-boilerplate-v1/app/cmd"
	"github.com/zYros90/pkg/testutils"
)

func TestRunExecute(t *testing.T) {
	rootPath, err := testutils.GetRootPath()
	if err != nil {
		t.Errorf("error getting root path: %v", err)
	}

	correctBaseConfigPath := "app/config/test.yaml"
	correctMergeConfigPath := "app/config/test_merge.yaml"

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			"run with --config",
			[]string{
				"run",
				"--config", filepath.Join(rootPath, correctBaseConfigPath),
			},
			false,
		},
		{
			"run with --mergeconfig",
			[]string{
				"run",
				"--config", filepath.Join(rootPath, correctBaseConfigPath),
				"--mergeconfig", filepath.Join(rootPath, correctMergeConfigPath),
			},
			false,
		},
		{
			"run with wrong flags",
			[]string{
				"run",
				"--abc", "xyz",
			},
			true,
		},
		{
			"run with wrong config path",
			[]string{
				"run",
				"--config", filepath.Join(rootPath, "app/config/xyz.yaml"),
			},
			true,
		},
		{
			"run wrong merge config path",
			[]string{
				"run",
				"--config", filepath.Join(rootPath, correctBaseConfigPath),
				"--mergeconfig", filepath.Join(rootPath, "app/wrong/path/x.yaml"),
			},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		cmd.RootCmd.SetArgs(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			err := cmd.RootCmd.Execute()
			if (err != nil) != tt.wantErr {
				if !strings.Contains(err.Error(), cmd.ErrMsg) {
					t.Log(err)
					t.Errorf("rootCmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
