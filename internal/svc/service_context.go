package svc

import (
	"github.com/patrickmn/go-cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"time"
	"video_game_api/internal/config"
	"video_game_api/internal/middleware"
	"video_game_api/internal/models"
)

type ServiceContext struct {
	Config config.Config

	// resource
	Cache *cache.Cache

	// middleware
	AuthMiddleware rest.Middleware

	// model
	GSceneModel models.GSceneModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		// resource
		Cache: cache.New(5*time.Minute, 10*time.Minute),

		// middleware
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,

		// model
		GSceneModel: models.NewGSceneModel(dbConn),
	}
}
