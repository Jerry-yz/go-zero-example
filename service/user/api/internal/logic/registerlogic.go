package logic

import (
	"context"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Gender:   req.Gender,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Error(500, "注册失败")
	}
	return &types.RegisterResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
