package model

import (
	"github.com/jinzhu/gorm"
)


// Log 日志
type Log struct {
	gorm.Model
	console       string
	date        string

}





