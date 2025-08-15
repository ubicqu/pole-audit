package model

import (
	"encoding/json"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/utils"
	"pole-audit/pkg/db"
	"pole-audit/pkg/fun"
	"testing"
)
import "github.com/brianvoe/gofakeit/v7"

func TestPole(t *testing.T) {
	var p Pole
	if err := gofakeit.Struct(&p); err != nil {
		t.Error(err)
	}
	db.Instance.Create(&p)
}

func TestPoleAndPoleInstallation(t *testing.T) {
	var pole Pole
	if err := gofakeit.Struct(&pole); err != nil {
		t.Error(err)
	}
	db.Instance.Create(&pole)
	fun.PrettyPrintln(pole)
	var poleInstallation PoleInstallation
	if err := gofakeit.Struct(&poleInstallation); err != nil {
		t.Error(err)
	}
	poleInstallation.PoleID = pole.ID
	db.Instance.Create(&poleInstallation)

	var p Pole
	db.Instance.Preload("PoleInstallations").First(&p, "id = ?", pole.ID)
	fun.PrettyPrintln(p)

	var pi PoleInstallation
	db.Instance.Preload("Pole").First(&pi, "id = ?", poleInstallation.ID)
	fun.PrettyPrintln(pi)
}

func TestModel(t *testing.T) {

	expected := fakePole()
	db.Instance.Create(&expected)

	var actual Pole
	db.Instance.
		Preload("PoleInstallations.PoleAudits.PoleAuditNotes").
		Preload(clause.Associations).
		Order("id DESC").
		First(&actual, "id = ?", expected.ID)

	utils.AssertEqual(expected, actual)
	eb, _ := json.Marshal(expected)
	ab, _ := json.Marshal(actual)
	if string(eb) != string(ab) {
		t.Errorf("exp: %s", string(eb))
		t.Errorf("act: %s", string(ab))
	}
}

func fakePole() Pole {
	var pole Pole
	if err := gofakeit.Struct(&pole); err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		pole.PoleInstallations = append(pole.PoleInstallations, fakePoleInstallation())
	}
	return pole
}

func fakePoleInstallation() PoleInstallation {
	var poleInstallation PoleInstallation
	if err := gofakeit.Struct(&poleInstallation); err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		poleInstallation.PoleAudits = append(poleInstallation.PoleAudits, fakePoleAudit())
	}
	return poleInstallation
}

func fakePoleAudit() PoleAudit {
	var poleAudit PoleAudit
	if err := gofakeit.Struct(&poleAudit); err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		poleAudit.PoleAuditNotes = append(poleAudit.PoleAuditNotes, fakePoleAuditNote())
	}
	return poleAudit
}

func fakePoleAuditNote() PoleAuditNote {
	var poleAuditNote PoleAuditNote
	if err := gofakeit.Struct(&poleAuditNote); err != nil {
		panic(err)
	}
	return poleAuditNote
}
