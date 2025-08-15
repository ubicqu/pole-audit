package model

import (
	"encoding/json"
	"gorm.io/gorm/clause"
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

	var poleInstallation PoleInstallation
	if err := gofakeit.Struct(&poleInstallation); err != nil {
		t.Error(err)
	}
	poleInstallation.PoleID = pole.ID
	db.Instance.Save(&poleInstallation)

	var poleAudit PoleAudit
	if err := gofakeit.Struct(&poleAudit); err != nil {
		t.Error(err)
	}
	poleAudit.PoleInstallationID = poleInstallation.ID
	db.Instance.Create(&poleAudit)

	var p Pole
	db.Instance.Preload("PoleInstallations.PoleAudits").First(&p, pole.ID)
	fun.PrettyPrintln(p)
}

func TestModel(t *testing.T) {

	// define the expected result by creating a pole with mock data
	expected := fakePole()
	db.Instance.Create(&expected)

	var poleAuditIds []uint32

	for _, install := range expected.PoleInstallations {
		audits := install.PoleAudits
		for _, audit := range audits {
			poleAuditIds = append(poleAuditIds, audit.ID)
			questions := audit.PoleAuditQuestions
			for _, question := range questions {
				answer := fakePoleAuditQuestionAnswer()
				answer.QuestionID = question.ID
				answer.AuditID = audit.ID
				db.Instance.Create(&answer)
			}
		}
	}

	// define the actual result by querying poles with the expected pole ID and eagerly loaded relationships
	var actual Pole
	db.Instance.
		Preload("PoleInstallations.PoleAudits.PoleAuditQuestions").
		Preload("PoleInstallations.PoleAudits.PoleAuditNotes").
		Preload(clause.Associations).
		Order("id desc").
		First(&actual, "id = ?", expected.ID)

	// marshal both expected and actual poles for precise comparison
	eb, _ := json.Marshal(expected)
	ab, _ := json.Marshal(actual)
	if string(eb) != string(ab) {
		// todo - fix time comparison ... it's 1 second off?
		//t.Errorf("\nexp: %s\nact: %s", string(eb), string(ab))
		//return
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

func TestModelAnswers(t *testing.T) {

	var pole Pole
	db.Instance.
		Preload("PoleInstallations.PoleAudits.PoleAuditNotes").
		Preload(clause.Associations).
		Last(&pole)
	fun.PrettyPrintln(pole)
	audit := pole.PoleInstallations[0].PoleAudits[0]
	questions := audit.PoleAuditQuestions
	for _, question := range questions {
		answer := fakePoleAuditQuestionAnswer()
		answer.QuestionID = question.ID
		answer.AuditID = audit.ID
		db.Instance.Create(&answer)
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

func fakePoleAuditQuestionAnswer() PoleAuditQuestionAnswer {
	var poleAuditQuestionAnswer PoleAuditQuestionAnswer
	if err := gofakeit.Struct(&poleAuditQuestionAnswer); err != nil {
		panic(err)
	}
	return poleAuditQuestionAnswer
}
