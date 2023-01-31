package logic

import (
	"context"

	"mall/common/cryptx"
	"mall/service/user/model"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	// "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	l.Logger.Debug("debug")
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == model.ErrNotFound {
		return nil, status.Error(500, "该用户不存在")
	}
	if err == nil {
		pw := cryptx.PasswordEncrypt(l.svcCtx.Config.Slat, in.Password)
		if pw != res.Password {
			return nil, status.Error(500, "密码错误")
		}
		return &user.LoginResponse{
			Id:     res.Id,
			Name:   res.Name,
			Mobile: res.Mobile,
			Gender: res.Gender,
		}, nil
	}
	return nil, status.Error(500, err.Error())
}
