package url_server

import (
	"github.com/gin-gonic/gin"
	"github.com/milad-rasouli/toker/internal/service"
	"go.uber.org/zap"
	"net/http"
)

type UrlHttp struct {
	logger     *zap.Logger
	urlService service.UrlService
}

func NewUrlHttp(logger *zap.Logger, urlService service.UrlService) *UrlHttp {
	return &UrlHttp{
		logger:     logger,
		urlService: urlService,
	}
}

func (u *UrlHttp) CreateOrGetUrl(ctx *gin.Context) {
	var (
		urlReq = ctx.Param("id")
		err    error
	)
	if urlReq == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing URL ID"})
		return
	}

	_, err = u.urlService.CreateOrGetUrl(ctx, urlReq)
	if err != nil {
		u.logger.Error("Failed to create or get URL", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "URL saved successfully", "data": "todo"})
}

func (u *UrlHttp) Register(group *gin.RouterGroup) {
	group.GET("/:id", u.CreateOrGetUrl)
}
