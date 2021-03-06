package logic

import (
	"context"

	"subModule/userApi/internal/svc"
	"subModule/userApi/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.InfoReq) (*types.InfoReply, error) {
	// todo: add your logic here and delete this line

	return &types.InfoReply{}, nil
}
