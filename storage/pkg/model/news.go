package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type News struct {
	gorm.Model
	Header string
	Date   time.Time
}
