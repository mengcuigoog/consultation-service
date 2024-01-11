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

func (pi *PatientInfoDao) GetOne() error {
	return nil
}
func (pi *PatientInfoDao) All() error {
	return nil
}

func (pi *PatientInfoDao) Del() error {
	return nil
}

func (pi *PatientInfoDao) Count() error {
	return nil
}
