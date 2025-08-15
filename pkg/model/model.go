package model

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
)
import "time"

type Bearing string

const (
	North = Bearing("n")
	East  = Bearing("e")
	South = Bearing("s")
	West  = Bearing("w")
)

func (b Bearing) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString([]string{string(North), string(East), string(South), string(West)}), nil
}

type Posterity struct {
	At time.Time `json:"at" gorm:"type:datetime" fake:"{pastdate}"`
	By string    `json:"by" fake:"{email}"`
}

func (p Posterity) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"at": p.At.Format("2006-01-02T15:04:05Z"),
		"by": p.By,
	})
}

type Pole struct {
	ID                uint32             `json:"id" gorm:"primaryKey" fake:"skip"`
	Latitude          float32            `json:"latitude" fake:"{latitude}"`
	Longitude         float32            `json:"longitude" fake:"{longitude}"`
	Street            string             `json:"street" fake:"{street}"`
	Bearing           Bearing            `json:"bearing"`
	Kind              *string            `json:"kind"`
	Height            int32              `json:"height" `
	Locked            bool               `json:"locked"`
	Created           Posterity          `json:"created" gorm:"embedded;embeddedPrefix:created_"`
	Updated           Posterity          `json:"updated" gorm:"embedded;embeddedPrefix:updated_"`
	Deleted           *Posterity         `json:"deleted,omitempty" gorm:"embedded;embeddedPrefix:deleted_" fake:"skip"`
	PoleInstallations []PoleInstallation `json:"installations,omitempty" fake:"skip"`
}

type PoleInstallation struct {
	ID         uint32      `json:"id" gorm:"primaryKey" fake:"skip"`
	PoleID     uint32      `json:"-"`
	UbiHubSn   *string     `json:"ubihub_sn" gorm:"column:ubihub_sn"`
	CamerasSn  *string     `json:"cameras_sn" gorm:"column:cameras_sn"`
	Start      *time.Time  `json:"start" gorm:"type:datetime" fake:"skip"`
	End        *time.Time  `json:"end" gorm:"type:datetime" fake:"skip"`
	Created    Posterity   `json:"created" gorm:"embedded;embeddedPrefix:created_"`
	Updated    Posterity   `json:"updated" gorm:"embedded;embeddedPrefix:updated_"`
	PoleAudits []PoleAudit `json:"audits" fake:"skip"`
}

type PoleAudit struct {
	ID                 uint32              `json:"id" gorm:"primaryKey" fake:"skip"`
	PoleInstallationID uint32              `json:"-"`
	Attempt            uint32              `json:"attempt"`
	State              string              `json:"state" fake:"{randomstring:[pending,approved,remediation,rejected]}"`
	Summary            *string             `json:"summary"`
	Auditor            string              `json:"auditor"`
	Created            Posterity           `json:"created" gorm:"embedded;embeddedPrefix:created_"`
	Updated            Posterity           `json:"updated" gorm:"embedded;embeddedPrefix:updated_"`
	PoleAuditNotes     []PoleAuditNote     `json:"notes,omitempty" fake:"skip"`
	PoleAuditQuestions []PoleAuditQuestion `json:"questions" gorm:"many2many:pole_auditables" fake:"skip"`
}

type PoleAuditNote struct {
	ID          uint32    `json:"id" gorm:"primaryKey" fake:"skip"`
	PoleAuditID uint32    `json:"-"`
	Type        string    `json:"type" gorm:"column:kind;default:photo" fake:"{randomstring:[photo,audio,video,text]}"`
	Datum       string    `json:"datum"`
	Created     Posterity `json:"created" gorm:"embedded;embeddedPrefix:created_"`
	Updated     Posterity `json:"updated" gorm:"embedded;embeddedPrefix:updated_"`
}

type PoleAuditQuestion struct {
	ID         uint32       `json:"id" gorm:"primaryKey" fake:"skip"`
	Device     string       `json:"device" fake:"{randomstring:[cell,hub,dtm,tvm]}"`
	Position   uint32       `json:"position"`
	Question   string       `json:"question"`
	Input      string       `json:"input" fake:"{randomstring:[radio,text,selections,number,photo]}"`
	Answer     string       `json:"answer"`
	Created    Posterity    `json:"created" gorm:"embedded;embeddedPrefix:created_"`
	Updated    Posterity    `json:"updated" gorm:"embedded;embeddedPrefix:updated_"`
	Deleted    Posterity    `json:"deleted" gorm:"embedded;embeddedPrefix:deleted_" fake:"skip"`
	PoleAudits *[]PoleAudit `json:"audits,omitempty" gorm:"many2many:pole_auditables" fake:"skip"`
}

type PoleAuditQuestionAnswer struct {
	ID      uint32    `json:"id" gorm:"primaryKey" fake:"skip"`
	Datum   string    `json:"datum"`
	Created Posterity `json:"created" gorm:"embedded;embeddedPrefix:created_"`
	//PoleAuditable PoleAuditable `json:"pole_auditable" fake:"skip"`
}

type PoleAuditable struct {
	AuditID    uint32 `json:"audit_id" gorm:"column:pole_audit_id"`
	QuestionID uint32 `json:"question_id" gorm:"column:pole_audit_question_id"`
	AnswerID   uint32 `json:"answer_id" gorm:"column:pole_audit_question_answer_id"`
}
