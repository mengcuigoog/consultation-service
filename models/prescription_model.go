package models

import (
	"database/sql/driver"
	"encoding/json"
)

type PrescriptionContext struct {
	Id    uint64
	Name  string
	Price float64
}

type PrescriptionContextSlice []PrescriptionContext

func (t *PrescriptionContextSlice) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t PrescriptionContextSlice) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type Prescription struct {
	PatientInfoId uint64
	Content       PrescriptionContextSlice `json:"prescriptionContextSlice"`
	Uid           string
}
