package models

import (
	"errors"

	"github.com/Wojbeg/go-gorilla-crm/database"
	"gorm.io/gorm"
)

type PersonInfo struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100); not null" json:"first_name"`
	Surname   string `gorm:"type:varchar(100); not null" json:"surname"`
	Company   string `gorm:"type:varchar(100)" json:"company"`
	Domicile  string `gorm:"type:varchar(100)" json:"domicile"`
	Notes     string `gorm:"type:varchar(255)" json:"notes"`
	Telephone string `gorm:"type:varchar(100)" json:"telephone"`
	Email     string `gorm:"type:varchar(100)" json:"email"`
}

func (PersonInfo) TableName() string {
	return "info"
}

func NewPersonInfo(firstName, surname, company, domicile, notes, telephone, email string) *PersonInfo {
	return &PersonInfo{
		FirstName: firstName,
		Surname:   surname,
		Company:   company,
		Domicile:  domicile,
		Notes:     notes,
		Telephone: telephone,
		Email:     email,
	}
}

func CreatePersonInfo(info *PersonInfo, rep database.Repository) (*PersonInfo, error) {
	err := rep.Create(&info).Error

	if err != nil {
		return nil, err
	}

	return info, nil
}

func GetAllInfos(rep database.Repository) []PersonInfo {
	var infos []PersonInfo
	rep.Find(&infos)

	return infos
}

func GetInfoById(Id int64, rep database.Repository) (*PersonInfo, error) {
	var info PersonInfo
	error := rep.Where("ID=?", Id).First(&info).Error

	if errors.Is(error, gorm.ErrRecordNotFound) {
		return nil, error
	}

	return &info, nil
}

func DeleteInfo(Id int64, rep database.Repository) (*PersonInfo, error) {
	info := &PersonInfo{}
	err := rep.Where("ID=?", Id).First(&info).Delete(Id, info).Error

	if err != nil {
		return nil, err
	}

	return info, nil
}

func Save(info *PersonInfo, rep database.Repository) (*PersonInfo, error) {
	err := rep.Save(&info).Error
	return info, err
}
