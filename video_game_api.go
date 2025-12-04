package main

import (
	"flag"
	"fmt"
	"github.com/JathamJ/zero_base/httpo"
	"github.com/zeromicro/go-zero/rest/httpx"

	"video_game_api/internal/config"
	"video_game_api/internal/handler"
	"video_game_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/video_game_api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//注册response handler
	httpx.SetOkHandler(httpo.DefaultOkHandler)
	httpx.SetErrorHandlerCtx(httpo.DefaultErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
