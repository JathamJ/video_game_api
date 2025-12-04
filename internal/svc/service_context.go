package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"video_game_api/internal/config"
	"video_game_api/internal/middleware"
	"video_game_api/internal/models"
)

type ServiceContext struct {
	Config config.Config

	// middleware
	AuthMiddleware rest.Middleware

	// model
	GSceneModel models.GSceneModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		// middleware
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,

		// model
		GSceneModel: models.NewGSceneModel(dbConn),
	}
}
