package game

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JathamJ/zero_base/errx"
	"github.com/JathamJ/zero_base/httpo"
	"github.com/JathamJ/zero_base/utilx"
	"github.com/patrickmn/go-cache"
	"video_game_api/internal/models"

	"video_game_api/internal/svc"
	"video_game_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	CacheKeySceneVideoResult = "s:scene:video:%d"
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
	resp := new(types.GameVideoResp)
	// read cache
	cacheKey := fmt.Sprintf(CacheKeySceneVideoResult, req.SceneId)
	value, ok := l.svcCtx.Cache.Get(cacheKey)
	if ok {
		valueStr := utilx.MustString(value)
		err := json.Unmarshal([]byte(valueStr), &resp)
		if err == nil {
			return resp, nil
		} else {
			l.Errorf("VideoLogic failed, json.Unmarshal cache, value: %s, err: %v", valueStr, err)
		}
	}

	// find video
	scene, err := l.svcCtx.GSceneModel.FindOne(l.ctx, req.SceneId)
	if errors.Is(err, models.ErrNotFound) {
		l.Infof("VideoLogic failed, GSceneModel.FindOne, record empty")
		return nil, httpo.NewCodeMsg(errx.RecordNotFound, "not found")
	}
	if err != nil {
		l.Errorf("VideoLogic failed, GSceneModel.FindOne, err: %v", err)
		return nil, httpo.NewCodeMsg(errx.SystemError, "Please retry later.")
	}

	// select children
	children, err := l.svcCtx.GSceneModel.FindListByGameIdAndParentId(l.ctx, req.GameId, req.SceneId)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		l.Errorf("VideoLogic failed, FindListByGameIdAndParentId, err: %v", err)
		return nil, httpo.NewCodeMsg(errx.SystemError, "Please retry later.")
	}

	resp = &types.GameVideoResp{
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

	// set cache
	respByte, err := json.Marshal(resp)
	if err != nil {
		l.Errorf("VideoLogic failed, json.Marshal, err: %v", err)
	} else {
		l.svcCtx.Cache.Set(cacheKey, string(respByte), cache.DefaultExpiration)
	}

	return resp, nil
}
