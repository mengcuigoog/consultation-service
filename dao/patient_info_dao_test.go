package dao

import (
	"consultation-service/models"
	"fmt"
	"testing"
)

var patientInfoDao *PatientInfoDao

func init() {
	Init()
	patientInfoDao = &PatientInfoDao{}
}

func TestCreatePatientInfo(t *testing.T) {
	patientInfos := make([]*models.PatientInfo, 0)
	for i := 0; i < 10; i++ {
		patientInfo := &models.PatientInfo{
			Name:    fmt.Sprintf("test_name_%d", i+1),
			Age:     13,
			Sex:     1,
			Address: "地球",
			Tel:     "13011259132",
		}
		patientInfos = append(patientInfos, patientInfo)
	}
	for _, p := range patientInfos {
		err := patientInfoDao.Create(p)
		// t.Logf("#####: %t\n", errors.Is(err, sqlite3.ErrConstraint))
		if err != nil {
			t.Errorf("create failed,err:%v", err)
		}
		t.Logf("id:%d", p.Id)
	}
}

func TestCreateBatchPatientInfo(t *testing.T) {
	patientInfos := make([]*models.PatientInfo, 0)
	for i := 0; i < 10; i++ {
		patientInfo := &models.PatientInfo{
			Name:    fmt.Sprintf("test_name_%d", i+1),
			Age:     13,
			Sex:     1,
			Address: "地球",
			Tel:     "13011259132",
		}
		patientInfos = append(patientInfos, patientInfo)
	}

	err := patientInfoDao.CreateBatches(patientInfos, 100)
	if err != nil {
		t.Errorf("create failed,err:%v", err)
	}
	for _, p := range patientInfos {
		t.Logf("Id:%d, Name:%s\n", p.Id, p.Name)
	}
}

func TestDelPatientInfo(t *testing.T) {
	pa := &models.PatientInfo{
		// Name: "test_name_5",
		Tel: "13011259132",
	}
	err := patientInfoDao.GetOne(pa)
	if err != nil {
		t.Errorf("get failed,%v", err)
		t.Fail()
	}

	err = patientInfoDao.Del(pa)
	if err != nil {
		t.Errorf("del %v err:%v", pa, err)
	}
}
