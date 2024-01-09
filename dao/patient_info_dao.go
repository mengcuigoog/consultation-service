package dao

import (
	"gorm.io/gorm"
)

type PatientInfoDao struct {
	Db *gorm.DB
}

func (pi *PatientInfoDao) Create() error {
	return nil
}
