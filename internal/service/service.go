package service

import "github.com/Bek0sh/online-market/main-page/internal/models"

type Repository interface {
	CreatePatientRequest(*models.ConsultationRequest) int
	GetRecommendation(int) ([]models.DoctorRecomemndation, error)
	CreateRecommendation(*models.DoctorRecomemndation) int
}
