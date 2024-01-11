package dao

import (
	"consultation-service/models"
	"fmt"

	"gorm.io/gorm"
)

/*
gorm id 有值时会自动添加过滤条件where id = ?
*/

type PrescriptionDao struct {
	Db *gorm.DB
}

func NewPrescriptionDao() *PrescriptionDao {
	p := &PrescriptionDao{}
	if Db == nil {
		panic("db is nil")
	}
	// p.Db = Db.Model(&models.Prescription{})
	p.Db = Db //Db.Model(&models.Prescription{})
	return p
}

func (p *PrescriptionDao) Create(prescription *models.Prescription) error {
	// createRet := Db.Model(&models.Prescription{}).Create(prescription)
	createRet := p.Db.Create(prescription)
	fmt.Printf(">>>> create err:%v\n", createRet.Error)
	return createRet.Error
}

// 仅获取一条数据
func (p *PrescriptionDao) GetOne(prescription *models.Prescription) error {
	tmpDb := p.Db.Model(&models.Prescription{})

	if prescription.PatientInfoId > 0 {
		tmpDb.Where("patient_info_id = ?", prescription.PatientInfoId)
	}

	err := tmpDb.Find(&models.ApiPrescription{}).First(prescription).Error
	if err != nil {
		return err
	}
	return err
}

func (p *PrescriptionDao) All() ([]models.Prescription, error) {
	tmpDb := p.Db.Model(&models.Prescription{})
	prescriptions := make([]models.Prescription, 0)
	err := tmpDb.Find(&models.ApiPrescription{}).Find(&prescriptions, "id > 0").Error
	return prescriptions, err
}

func (p *PrescriptionDao) Del(prescription *models.Prescription) error {
	tmpDb := p.Db.Model(&models.Prescription{})
	if prescription.Id > 0 {
		tmpDb = tmpDb.Where("id = ?", prescription.Id)
	}
	if prescription.PatientInfoId > 0 {
		tmpDb = tmpDb.Where("patient_info_id = ?", prescription.PatientInfoId)
	}
	err := tmpDb.Delete(&models.Prescription{}).Error
	return err
}

func (p *PrescriptionDao) Count() (int64, error) {
	var count int64
	err := p.Db.Model(&models.Prescription{}).Count(&count).Error
	return count, err
}
