package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data"
	"github.com/zYros90/go-boilerplate-v1/app/internal/server"
	"github.com/zYros90/go-boilerplate-v1/app/internal/service"
	"github.com/zYros90/pkg/logger"
)

// ErrMsg is used by unit test from runCmd.
const ErrMsg = "error starting app"

// runCmd represents the run command.
var runCmd = &cobra.Command{ // nolint:exhaustruct,gochecknoglobals
	Use:   "run",
	Short: "run the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// init config
		conf, err := config.ReadConfig(cfgFile, cfgMergeFile)
		if err != nil {
			return err
		}

		// init logger
		logger, err := logger.NewLogger(conf.LogLevel, conf.Develop, false, false)
		if err != nil {
			return err
		}

		// start app
		err = startApp(conf, logger)
		if err != nil {
			return errors.Wrap(err, ErrMsg)
		}

		return nil
	},
}

func init() { // nolint:gochecknoinits
	RootCmd.AddCommand(runCmd)
}

func startApp(conf *config.Config, logger *logger.Log) error {
	// init data layer
	db, err := data.New(logger, conf)
	if err != nil {
		return err
	}

	// init repos
	usrRepo := data.NewUser(db, logger)
	todoRepo := data.Newtodo(db, logger)

	// init usecases
	bizUsr := biz.NewUserUsecase(usrRepo, logger.Logger, conf)
	bizLogin := biz.NewLoginUsecase(logger.Logger, conf, bizUsr)
	bizTodo := biz.NewTodoUsecase(todoRepo, logger.Logger, conf)

	// init service
	svc := service.New(conf, logger.Logger, bizUsr, bizLogin, bizTodo)

	// init server
	srv, err := server.NewEcho(conf, logger.Logger, svc)
	if err != nil {
		return err
	}

	// start server
	logger.Sugar().Info("starting server")
	return srv.Start(fmt.Sprintf(":%d", conf.Server.Port))
}
