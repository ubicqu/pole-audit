package model

import (
	"encoding/json"
	"gorm.io/gorm/clause"
	"pole-audit/pkg/db"
	"testing"
)
import "github.com/brianvoe/gofakeit/v7"

func TestModel(t *testing.T) {

	// define the expected result by creating a pole with mock data
	expected := fakePole()
	db.Instance.Create(&expected)

	// for each persisted installation, collect the audit id and mock audit question answers
	var poleAuditIds []uint32
	for _, install := range expected.PoleInstallations {
		poleAuditIds = append(poleAuditIds, install.PoleAudit.ID)
		for _, question := range install.PoleAudit.PoleAuditQuestions {
			answer := fakePoleAuditQuestionAnswer(install.PoleAudit.ID, question.ID)
			db.Instance.Create(&answer)
		}
	}

	// define the actual result by querying poles with the expected pole ID
	// eagerly load immediate associations and nested relationships
	var actual Pole
	db.Instance.
		Preload("PoleInstallations.PoleAudit.PoleAuditQuestions").
		Preload("PoleInstallations.PoleAudit.PoleAuditNotes").
		Preload(clause.Associations).
		First(&actual, "id = ?", expected.ID)

	// marshal both expected and actual poles for precise comparison
	eb, _ := json.Marshal(expected)
	ab, _ := json.Marshal(actual)
	if string(eb) != string(ab) {
		t.Errorf("\nexp: %s\nact: %s", string(eb), string(ab))
		return
	}

	var answers []PoleAuditQuestionAnswer
	db.Instance.
		Preload("PoleAudit.PoleInstallation.Pole").
		Preload(clause.Associations).
		Find(&answers, "pole_audit_id IN ?", poleAuditIds)

	if len(answers) == 0 {
		t.Fail()
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
	poleInstallation.PoleAudit = fakePoleAudit()
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
	for i := 0; i < 2; i++ {
		poleAudit.PoleAuditQuestions = append(poleAudit.PoleAuditQuestions, fakePoleAuditQuestion())
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

func fakePoleAuditQuestion() PoleAuditQuestion {
	var poleAuditQuestion PoleAuditQuestion
	if err := gofakeit.Struct(&poleAuditQuestion); err != nil {
		panic(err)
	}
	return poleAuditQuestion
}

func fakePoleAuditQuestionAnswer(auditID, questionID uint32) PoleAuditQuestionAnswer {
	var poleAuditQuestionAnswer PoleAuditQuestionAnswer
	if err := gofakeit.Struct(&poleAuditQuestionAnswer); err != nil {
		panic(err)
	}
	poleAuditQuestionAnswer.AuditID = auditID
	poleAuditQuestionAnswer.QuestionID = questionID
	return poleAuditQuestionAnswer
}
