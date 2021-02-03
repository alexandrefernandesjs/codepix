package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Bank struct{
	ID string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} 

func newBank(code.string, name.string)(*Bank, error){
	bank := Bank{
		Code: code, 
		Name: name,
	}

	return  &bank null
}

