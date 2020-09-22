package models

import (
	"encoding/json"
	"gorone_api/db"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Request タスク
type Request struct {
	gorm.Model
	Body datatypes.JSON
}

func NewRequest(params interface{}) (*Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req := Request{Body: json}
	db := db.DbManager()
	db.Create(&req)

	return &req, nil
}

func (r Request) GetValues() []string {
	var dat map[string][]string
	err := json.Unmarshal(r.Body, &dat)
	if err != nil {
		return []string{}
	}
	return dat["values"]
}
