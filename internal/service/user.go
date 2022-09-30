package service

import (
	"context"
	"github.com/davidchou93/wallets/internal/model"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	repository model.Repository
	userRepo   model.UserRepo
	logger     *logrus.Logger
}

func NewUserServiceImpl(repo model.Repository, logger *logrus.Logger) *UserServiceImpl {
	svc := &UserServiceImpl{}
	svc.repository = repo
	svc.userRepo = repo.UserRepo()
	svc.logger = logger
	return svc
}

func (svc *UserServiceImpl) GetList(ctx context.Context, option model.GetUserOption) ([]model.User, error) {
	return svc.userRepo.GetList(ctx, option)
}
