package models

import (
	u "daas/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Engineer struct
type Engineer struct {
	gorm.Model

	FirstName    string
	LastName     string
	Skills       string
	Experience   string
	Availability string
}

// EngineerPortfolio struct
type EngineerPortfolio struct {
	gorm.Model

	Skills       string
	Experience   string
	Availability string
}

// ClientAccount struct
type ClientAccount struct {
	gorm.Model

	Name     string
	Username string
	Password string
}

// ClientDescription struct
type ClientDescription struct {
	gorm.Model

	Business       string
	TeamSize       string
	Location       string
	Specialization string
}

func (engineer *Engineer) Validate() (map[string]interface{}, bool) {

	if engineer.FirstName == "" {
		return u.Message(false, "Please provide your first name"), false
	}

	if engineer.LastName == "" {
		return u.Message(false, "Please provide your last name"), false
	}

	if engineer.Skills == "" {
		return u.Message(false, "Please provide your skills"), false
	}

	if engineer.Experience == "" {
		return u.Message(false, "Please provide your availability"), false
	}

	return u.Message(true, "Details saved successfully"), true
}

func (engineer *Engineer) Create() map[string]interface{} {

	if resp, ok := engineer.Validate(); !ok {
		return resp
	}

	GetDB().Create(engineer)

	resp := u.Message(true, "success")
	resp["engineer"] = engineer
	fmt.Println(resp)
	return resp
}

func GetEngineers() []*Engineer {

	engineers := make([]*Engineer, 0)
	err := GetDB().Table("engineers").Find(&engineers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return engineers
}

func GetEngineer(id string) *Engineer {

	engineer := &Engineer{}
	err := GetDB().Table("engineers").Where("email_address = ?", id).First(engineer).Error
	if err != nil {
		return nil
	}

	return engineer
}
