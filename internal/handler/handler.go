package handler

import (
	"net/http"
	"strconv"

	"github.com/Bek0sh/online-market/main-page/internal/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreatePatientRequest(*models.ConsultationRequest) (int, error)
	GetRecommendations(int) ([]models.DoctorRecomemndation, error)
	CreateRecommendation(*models.DoctorRecomemndation) (int, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreatePatientRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.ConsultationRequest

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
					"status":  "fail",
				},
			)
			return
		}

		id, err := h.service.CreatePatientRequest(&input)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
					"status":  "fail",
				},
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data":   id,
				"status": "success",
			},
		)
	}
}

func (h *Handler) GetRecommendations() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId := ctx.Param("user_id")

		id, err := strconv.Atoi(userId)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
					"status":  "fail",
				},
			)
			return
		}

		recs, err := h.service.GetRecommendations(id)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
					"status":  "fail",
				},
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data":   recs,
				"status": "success",
			},
		)
	}
}

func (h *Handler) CreateRecommendation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.DoctorRecomemndation

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
					"status":  "fail",
				},
			)
			return
		}

		id, err := h.service.CreateRecommendation(&input)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
					"status":  "fail",
				},
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data":   id,
				"status": "success",
			},
		)
	}
}
