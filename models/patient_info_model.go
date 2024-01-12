package models

type PatientInfo struct {
	Id      uint
	Name    string `gorm:"unique"`
	Age     uint
	Sex     uint
	Address string
	Tel     string `gorm:"index"`
}

type ApiPatientInfo struct {
	Id      uint
	Name    string `gorm:"unique"`
	Age     uint
	Sex     uint
	Address string
	Tel     string `gorm:"index"`
}
