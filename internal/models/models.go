package models

import "time"

type UserInfo struct {
	Id          int    `json:"-"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	UserType    string `json:"user_type"`
}

type ConsultationRequest struct {
	Id          int       `json:"id, omitempty"`
	Patient     UserInfo  `json:"patient"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type DoctorRecomemndation struct {
	Id             int                 `json:"-"`
	Request        ConsultationRequest `json:"consultation_request"`
	Recommendation string              `json:"recommendation"`
}
