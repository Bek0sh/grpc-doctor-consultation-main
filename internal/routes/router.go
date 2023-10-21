package routes

import (
	"github.com/Bek0sh/online-market/main-page/internal/handler"
	"github.com/gin-gonic/gin"
)

func Consultation(r *gin.Engine, h handler.Handler) {
	r.POST("/v1/request", h.CreatePatientRequest())
	r.POST("/v1/recom", h.CreateRecommendation())
	r.GET("/v1/:user_id/recom", h.GetRecommendations())
}
