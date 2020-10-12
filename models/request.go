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

// Request Request
type Request struct {
	RedisTagName string
	JobType      string
	DesiredCount int
	gorm.Model
	Body   datatypes.JSON
	DoneAt *time.Time
}

// NewRequest 新規リクエスト発行
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

// DecodedBody bodyをJSONデコードしたものを返す
func (r Request) DecodedBody() []string {
	var data map[string][]string
	err := json.Unmarshal(r.Body, &data)
	if err != nil {
		return []string{}
	}
	return data["values"]
}

// randTagName redis用のタグに使う文字列を生成する
func randTagName() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, 5)
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
