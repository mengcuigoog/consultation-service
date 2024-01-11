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
	v, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type Prescription struct {
	Id            uint
	PatientInfoId uint64
	Content       PrescriptionContextSlice `grom:"text" json:"prescriptionContextSlice"`
}

type ApiPrescription struct {
	Id            uint
	PatientInfoId uint64
	Content       PrescriptionContextSlice `grom:"text" json:"prescriptionContextSlice"`
}
