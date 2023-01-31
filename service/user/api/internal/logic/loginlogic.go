package logic

import (
	"context"
	"time"

	"mall/common/jwtx"
	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, res.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccessToken:  token,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire + time.Now().Unix(),
	}, nil
}
