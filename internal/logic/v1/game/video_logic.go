package game

import (
	"context"
	"errors"
	"github.com/JathamJ/zero_base/errx"
	"github.com/JathamJ/zero_base/httpo"
	"video_game_api/internal/models"

	"video_game_api/internal/svc"
	"video_game_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoLogic {
	return &VideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoLogic) Video(req *types.GameVideoReq) (*types.GameVideoResp, error) {

	// find video
	scene, err := l.svcCtx.GSceneModel.FindOne(l.ctx, req.SceneId)
	if errors.Is(err, models.ErrNotFound) {
		l.Infof("VideoLogic failed, GSceneModel.FindOne, record empty")
		return nil, httpo.NewCodeMsg(errx.RecordNotFound, "not found")
	}
	if err != nil {
		l.Infof("VideoLogic failed, GSceneModel.FindOne, err: %v", err)
		return nil, httpo.NewCodeMsg(errx.SystemError, "Please retry later.")
	}

	// select children
	children, err := l.svcCtx.GSceneModel.FindListByGameIdAndParentId(l.ctx, req.GameId, req.SceneId)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		l.Infof("VideoLogic failed, FindListByGameIdAndParentId, err: %v", err)
		return nil, httpo.NewCodeMsg(errx.SystemError, "Please retry later.")
	}

	resp := &types.GameVideoResp{
		Id:      scene.Id,
		Video:   scene.VideoUrl,
		Title:   scene.Title,
		Brief:   scene.Brief,
		Options: make([]types.GameVideoOption, 0),
	}

	for _, v := range children {
		resp.Options = append(resp.Options, types.GameVideoOption{
			Id:    v.Id,
			Value: v.Label,
			Audio: v.LabelAudio,
			Video: v.VideoUrl,
			Title: v.Title,
		})
	}

	return resp, nil
}
