package dao

import (
	"consultation-service/models"
	"fmt"

	"gorm.io/gorm"
)

type PrescriptionDao struct {
	Db *gorm.DB
}

func NewPrescriptionDao() *PrescriptionDao {
	p := &PrescriptionDao{}
	if Db == nil {
		panic("db is nil")
	}
	p.Db = Db.Model(&models.Prescription{})
	return p
}

func (p *PrescriptionDao) Create(prescription *models.Prescription) error {
	createRet := Db.Model(&models.Prescription{}).Create(prescription)
	fmt.Printf(">>>> create err:%v\n", createRet.Error)
	return createRet.Error
}

func (p *PrescriptionDao) Get(prescription *models.Prescription) error {
	// createRet := Db.Model(&models.Prescription{}).Find(prescription)
	createRet := Db.Model(&models.Prescription{}).Where("uid = ?", prescription.Uid).Find(prescription)
	fmt.Printf(">>  get prescription err:%v\n", createRet.Error)
	return createRet.Error
}

func (p *PrescriptionDao) Del(prescription *models.Prescription) error {
	createRet := Db.Model(&models.Prescription{}).Where("uid = ?", prescription.Uid).Delete(prescription)
	fmt.Printf(">>  delete prescription err:%v\n", createRet.Error)
	return createRet.Error
}
