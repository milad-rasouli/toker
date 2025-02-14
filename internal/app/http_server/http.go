package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	urlHttp "github.com/milad-rasouli/toker/internal/app/http_server/url"
	"github.com/milad-rasouli/toker/internal/infra/config"
	"go.uber.org/zap"
)

type HttpServer struct {
	cfg    *config.Config
	uHttp  *urlHttp.UrlHttp
	logger *zap.Logger
}

func NewHttpServer(cfg *config.Config, logger *zap.Logger, u *urlHttp.UrlHttp) *HttpServer {
	return &HttpServer{
		cfg:    cfg,
		logger: logger,
		uHttp:  u,
	}
}

func (h *HttpServer) Start() error {
	//if h.cfg.Development == true {
	//	gin.SetMode(gin.DebugMode)
	//} else {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	urlGroup := router.Group("/url")
	h.uHttp.Register(urlGroup)

	serverAddr := h.cfg.App.Port
	h.logger.Info("Starting server", zap.String("address", serverAddr))

	err := router.Run(serverAddr)
	if err != nil {
		return fmt.Errorf("failed to run http_server")
	}
	return nil
}
