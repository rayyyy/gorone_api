package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// CalcResult タスク
type CalcResult struct {
	gorm.Model
	KeyName string
	Result  datatypes.JSON
}
