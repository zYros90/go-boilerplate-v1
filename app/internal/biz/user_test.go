package biz_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/biz"
	"github.com/zYros90/go-boilerplate-v1/app/mocks"
	"github.com/zYros90/pkg/hash"
	"github.com/zYros90/pkg/logger"
)

func TestUserBiz_Create(t *testing.T) {
	t.Parallel()
	conf := &config.Config{
		Develop:  true,
		LogLevel: "debug",
	}

	log, err := logger.NewLogger("warn", true, false, false)
	require.NoError(t, err)

	t.Run("Successfully get User", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		userRepo := mocks.NewMockUserRepo(ctrl)

		password := "password123"
		hashedPswd, err := hash.HashPassword(password)
		require.NoError(t, err)

		input := &biz.User{
			Username:  "Max",
			Password:  password,
			Email:     "max@example.com",
			FirstName: "Max",
			LastName:  "Mustermann",
		}

		expected := &biz.User{
			Username:  "Max",
			Password:  hashedPswd,
			Email:     "max@example.com",
			FirstName: "Max",
			LastName:  "Mustermann",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		userRepo.EXPECT().Create(gomock.Any(), input).Return(expected, nil)

		bizUser := biz.NewUserUsecase(userRepo, log.Logger, conf)

		actual, err := bizUser.Create(context.Background(), input)
		require.NoError(t, err)

		assert.Equal(t, expected, actual)
	})
}
