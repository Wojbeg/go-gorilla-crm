package dto

import (
	"time"

	"github.com/Wojbeg/go-gorilla-crm/models"
)

type PersonInfoRequest struct {
	FirstName string `json:"first_name" schema:"name,required"`
	Surname   string `json:"surname" schema:"surname,required"`
	Company   string `json:"company" schema:"company"`
	Domicile  string `json:"domicile" schema:"domicile"`
	Notes     string `json:"notes" schema:"notes"`
	Telephone string `json:"telephone" schema:"phone"`
	Email     string `json:"email" schema:"email"`
}

func (info *PersonInfoRequest) InfoRequestToModel() *models.PersonInfo {
	return &models.PersonInfo{
		FirstName: info.FirstName,
		Surname:   info.Surname,
		Company:   info.Company,
		Domicile:  info.Domicile,
		Notes:     info.Notes,
		Telephone: info.Telephone,
		Email:     info.Email,
	}
}

func (info *PersonInfoRequest) CopyInfoRequestToModel(model *models.PersonInfo) {
	model.FirstName = info.FirstName
	model.Surname = info.Surname
	model.Company = info.Company
	model.Domicile = info.Domicile
	model.Notes = info.Notes
	model.Telephone = info.Telephone
	model.Email = info.Email
}

type PersonInfoResponse struct {
	Id        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	Surname   string    `json:"surname"`
	Company   string    `json:"company"`
	Domicile  string    `json:"domicile"`
	Notes     string    `json:"notes"`
	Telephone string    `json:"telephone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func InfoModelToResponse(info *models.PersonInfo) *PersonInfoResponse {
	return &PersonInfoResponse{
		Id:        info.ID,
		FirstName: info.FirstName,
		Surname:   info.Surname,
		Company:   info.Company,
		Domicile:  info.Domicile,
		Notes:     info.Notes,
		Telephone: info.Telephone,
		Email:     info.Email,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}
}

func InfoModelsToResponses(info *[]models.PersonInfo) *[]PersonInfoResponse {
	slice := make([]PersonInfoResponse, len(*info))

	for i, inf := range *info {
		slice[i] = *InfoModelToResponse(&inf)
	}

	return &slice
}
