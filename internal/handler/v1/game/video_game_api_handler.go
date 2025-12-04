package game

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"video_game_api/internal/logic/v1/game"
	"video_game_api/internal/svc"
	"video_game_api/internal/types"
)

func Video_game_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := game.NewVideo_game_apiLogic(r.Context(), svcCtx)
		resp, err := l.Video_game_api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
