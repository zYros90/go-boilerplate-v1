package cmd_test

import (
	"path/filepath"
	"testing"

	"github.com/zYros90/go-boilerplate-v1/app/cmd"
	"github.com/zYros90/pkg/testutils"
)

func TestRootExecute(t *testing.T) {
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
			"root with --help flag",
			[]string{"--help"},
			false,
		},
		{
			"root with -t flag",
			[]string{"-t"},
			false,
		},
		{
			"root with --config flag",
			[]string{"--config", filepath.Join(rootPath, "app/config/base.yaml")},
			false,
		},
		{
			"root with --mergeconfig flag",
			[]string{"--mergeconfig", filepath.Join(rootPath, "app/config/dev.yaml")},
			false,
		},
		{
			"root with wrong flag",
			[]string{"--randomflag", "xyz"},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		cmd.RootCmd.SetArgs(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			err := cmd.RootCmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("rootCmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
