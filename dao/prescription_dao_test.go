package dao

import (
	"consultation-service/models"
	"fmt"
	"testing"
)

var (
	p *PrescriptionDao
)

func init() {
	Init()
	p = NewPrescriptionDao()
}

func TestCreatePrescription(t *testing.T) {
	c := make([]models.PrescriptionContext, 0)
	for i := 0; i < 10; i++ {
		tmp := models.PrescriptionContext{
			Id:    uint64(i + 1),
			Name:  fmt.Sprintf("huangqi_%d", i),
			Price: 33.33,
		}
		c = append(c, tmp)
	}
	m := models.Prescription{
		PatientInfoId: 12,
		Content:       c,
		Uid:           "uuid12312314",
	}
	err := p.Create(&m)
	t.Logf("####:%v\n", err)
}

func TestGetPrescription(t *testing.T) {
	m := models.Prescription{
		Uid: "uuid12312314",
	}
	err := p.Get(&m)
	if err != nil {
		t.Errorf("AAAAAA: %v\n", err)
	}
	for _, c := range m.Content {
		t.Logf("####:%d %s %f\n", c.Id, c.Name, c.Price)
	}

}

func TestDelPrescription(t *testing.T) {
	m := models.Prescription{
		Uid: "uuid12312314",
	}
	err := p.Get(&m)
	t.Logf("get ####:%q %v\n", err, m)

	err = p.Del(&m)
	if err != nil {
		t.Errorf("del ####:%v %v\n", err, m)
	}
}
