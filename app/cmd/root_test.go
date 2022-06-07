package cmd

import (
	"path/filepath"
	"testing"

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
			"root",
			[]string{"--help"},
			false,
		},
		{
			"root",
			[]string{"-t"},
			false,
		},
		{
			"root",
			[]string{"--config", filepath.Join(rootPath, "app/config/base.yaml")},
			false,
		},
		{
			"root",
			[]string{"--mergeconfig", filepath.Join(rootPath, "app/config/dev.yaml")},
			false,
		},
		{
			"root",
			[]string{"--randomflag", "xyz"},
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
