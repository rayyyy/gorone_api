package models

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"gorone_api/db"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Request タスク
type Request struct {
	RedisTagName string
	JobType      string
	DesiredCount int
	gorm.Model
	Body   datatypes.JSON
	DoneAt *time.Time
}

func NewRequest(jobType string, params interface{}, desiredCount int) (*Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req := Request{Body: json, JobType: jobType, RedisTagName: randTagName(), DesiredCount: desiredCount, DoneAt: nil}
	db := db.DbManager()
	if err := db.Create(&req).Error; err != nil {
		return nil, err
	}

	return &req, nil
}

func (r Request) DecodedBody() []string {
	var data map[string][]string
	err := json.Unmarshal(r.Body, &data)
	if err != nil {
		return []string{}
	}
	return data["values"]
}

func randTagName() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		fmt.Println("Error: " + err.Error())
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	// TODO: すでに存在しないかを確認したい
	return result
}
