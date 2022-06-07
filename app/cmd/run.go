package cmd

import (
	"fmt"

	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"github.com/zYros90/go-boilerplate-v1/app/internal/data"
	"github.com/zYros90/go-boilerplate-v1/app/internal/server"
	"github.com/zYros90/go-boilerplate-v1/app/internal/service"
	"github.com/zYros90/pkg/logger"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// init config
		conf, err := config.ReadConfig(cfgFile, cfgMergeFile)
		if err != nil {
			return err
		}

		// init logger
		logger, err := logger.NewLogger("debug", true, false, false)
		if err != nil {
			return err
		}

		if testMode {
			return nil
		}

		// init server
		err = initServer(conf, logger)
		if err != nil {
			logger.Sugar().Error(err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func initServer(conf *config.Config, logger *logger.Log) error {
	// init db
	db, err := data.New(logger, conf)
	if err != nil {
		logger.Sugar().Fatal(err)
	}
	// init repos
	usrRepo := data.NewUser(db, logger)

	// init usecases
	bizUsr := biz.NewUserUsecase(usrRepo, logger.Logger, conf)
	bizLogin := biz.NewLoginUsecase(logger.Logger, conf, bizUsr)

	// init server
	srv, err := server.New(conf, logger.Logger)
	if err != nil {
		logger.Sugar().Fatal(err)
	}

	// init service
	service.Init(conf, logger.Logger, srv, bizUsr, bizLogin)

	// start server
	logger.Sugar().Info("starting server")
	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", conf.Server.Port)))
	return nil
}
