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
	Id                   int
	PatientInfoId        int
	Content              PrescriptionContextSlice `grom:"text" json:"prescriptionContextSlice"` //处方内容
	ConditionDescription string                   // 问诊信息描述
}

type ApiPrescription struct {
	Id                   uint
	PatientInfoId        int
	Content              PrescriptionContextSlice `grom:"text" json:"prescriptionContextSlice"` //处方内容
	ConditionDescription string
}
