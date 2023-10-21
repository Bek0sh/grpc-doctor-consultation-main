package service

import (
	"fmt"

	"github.com/Bek0sh/online-market/main-page/internal/handler"
	"github.com/Bek0sh/online-market/main-page/internal/models"
)

type Repository interface {
	CreatePatientRequest(*models.ConsultationRequest) int
	GetRecommendation(req *models.UserInfo) ([]models.DoctorRecomemndation, error)
	CreateRecommendation(*models.DoctorRecomemndation) int
}

type service struct {
	client grpcClient
	repo   Repository
}

func NewService(repo Repository) handler.Service {
	client := newGrpcClient()
	return &service{repo: repo, client: *client}
}

func (s *service) CreatePatientRequest(req *models.ConsultationRequest) (int, error) {

	err := s.client.CheckToken()
	if err != nil {
		return 0, fmt.Errorf("you need to authorize at first, error: %s", err.Error())
	}

	user, err := s.client.GetCurrentUser()

	if err != nil {
		return 0, fmt.Errorf("failed to find current user, error: %s", err.Error())
	}

	req.Patient = *user

	id := s.repo.CreatePatientRequest(req)

	return id, nil
}

func (s *service) GetRecommendations(id int) ([]models.DoctorRecomemndation, error) {
	err := s.client.CheckToken()
	if err != nil {
		return nil, fmt.Errorf("you need to authorize at first, error: %s", err.Error())
	}
	user, err := s.client.GetCurrentUser()
	if err != nil {
		return nil, fmt.Errorf("failed to find current user, error: %s", err.Error())
	}
	return s.repo.GetRecommendation(user)
}

func (s *service) CreateRecommendation(req *models.DoctorRecomemndation) (int, error) {
	err := s.client.CheckToken()
	if err != nil {
		return 0, fmt.Errorf("you need to authorize at first, error: %s", err.Error())
	}
	err = s.client.CheckUserType()
	if err != nil {
		return 0, fmt.Errorf("do not have permission, error: %s", err.Error())
	}
	return s.repo.CreateRecommendation(req), nil
}
