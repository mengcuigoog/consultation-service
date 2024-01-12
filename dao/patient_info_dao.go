package dao

import (
	"consultation-service/models"
)

type PatientInfoDao struct {
	// Db *gorm.DB
}

func (pi *PatientInfoDao) Create(patientInfo *models.PatientInfo) error {
	err := Db.Create(patientInfo).Error
	return err
}

func (pi *PatientInfoDao) CreateBatches(patientInfo []*models.PatientInfo, batchSize int) error {
	err := Db.CreateInBatches(patientInfo, batchSize).Error
	return err
}

func (pi *PatientInfoDao) GetOne(patientInfo *models.PatientInfo) error {
	db := Db.Model(&models.PatientInfo{}).Find(&models.ApiPatientInfo{})
	if patientInfo.Tel != "" {
		db = db.Where("tel = ?", patientInfo.Tel)
	}
	if patientInfo.Name != "" {
		db = db.Where("name = ?", patientInfo.Name)
	}
	err := db.First(patientInfo).Error
	return err
}

func (pi *PatientInfoDao) All() ([]models.PatientInfo, error) {
	prescriptions := make([]models.PatientInfo, 0)
	db := Db.Model(&models.PatientInfo{}).Find(&models.ApiPatientInfo{})
	db = db.Where("id > ?", 0)
	err := db.Find(&prescriptions).Error
	return prescriptions, err
}

func (pi *PatientInfoDao) Del(patientInfo *models.PatientInfo) error {
	db := Db.Model(&models.PatientInfo{})
	if patientInfo.Id > 0 {
		db = db.Where("id = ?", patientInfo.Id)
	}
	if patientInfo.Tel != "" {
		db = db.Where("tel = ?", patientInfo.Tel)
	}
	if patientInfo.Name != "" {
		db = db.Where("name = ?", patientInfo.Name)
	}
	err := db.Delete(&models.Prescription{}).Error
	return err
}

func (pi *PatientInfoDao) Count() (int64, error) {
	var count int64
	err := Db.Model(&models.PatientInfo{}).Count(&count).Error
	return count, err
}
