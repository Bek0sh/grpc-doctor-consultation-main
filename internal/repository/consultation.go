package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Bek0sh/online-market/main-page/internal/models"
	"github.com/Bek0sh/online-market/main-page/internal/service"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) service.Repository {
	return &repo{db: db}
}

func (r *repo) CreatePatientRequest(input *models.ConsultationRequest) int {
	var id int
	query := "INSERT INTO requests (user_id, description, created_at) VALUES($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, input.Patient.Id, input.Description, input.CreatedAt).Scan(&id)
	if err != nil {
		log.Print(err)
		return 0
	}

	return id
}

func (r *repo) GetRecommendation(id int) ([]models.DoctorRecomemndation, error) {
	var recs []models.DoctorRecomemndation
	query := "SELECT req_id, recomm FROM recommendations rec JOIN requests r ON rec.req_id = r.id where r.user_id = $1"

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find recommendations, error: %s", err.Error())
	}

	for rows.Next() {
		recom := models.DoctorRecomemndation{}
		err := rows.Scan(&recom.Request.Id, &recom.Recommendation)
		if err != nil {
			return nil, fmt.Errorf("failed to scan recommendations, error: %s", err.Error())
		}
		consReq, err := r.GetPatientRequest(recom.Request.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to find cons req, error: %s", err.Error())
		}
		consReq.Id = recom.Request.Id
		recom.Request = *consReq
		recs = append(recs, recom)
	}
	return recs, nil
}

func (r *repo) GetPatientRequest(id int) (*models.ConsultationRequest, error) {
	var response models.ConsultationRequest

	query := "SELECT user_id, description, created_at FROM requests WHERE id=$1"
	err := r.db.QueryRow(query, &id).Scan(&response.Patient.Id, &response.Description, &response.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to find patient request with id=%d, error: %s", id, err.Error())
	}

	return &response, nil
}

func (r *repo) CreateRecommendation(input *models.DoctorRecomemndation) int {
	var id int
	query := "INSERT INTO recommendations (req_id, recomm) VALUES($1, $2) RETURNING id"
	err := r.db.QueryRow(query, input.Request.Id, input.Recommendation).Scan(&id)
	if err != nil {
		log.Print(err)
		return 0
	}
	return id
}
