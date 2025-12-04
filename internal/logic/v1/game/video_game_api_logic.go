package game

import (
	"context"

	"video_game_api/internal/svc"
	"video_game_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Video_game_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideo_game_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Video_game_apiLogic {
	return &Video_game_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Video_game_apiLogic) Video_game_api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
