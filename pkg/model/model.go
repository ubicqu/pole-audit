package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

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

type State string

const (
	PENDING     = State("pending")
	APPROVED    = State("approved")
	REMEDIATION = State("remediation")
	REJECTED    = State("rejected")
)

func (b State) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString([]string{string(PENDING), string(APPROVED), string(REMEDIATION), string(REJECTED)}), nil
}

type Device string

const (
	CELL = Device("cell")
	DTM  = Device("dtm")
	TVM  = Device("tvm")
	HUB  = Device("hub")
)

func (b Device) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString([]string{string(CELL), string(DTM), string(TVM), string(HUB)}), nil
}

type Input string

const (
	RadioInput      = Input("radio")
	TextInput       = Input("text")
	SelectionsInput = Input("selections")
	NumberInput     = Input("number")
	PhotoInput      = Input("photo")
)

func (b Input) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString([]string{
		string(RadioInput),
		string(TextInput),
		string(SelectionsInput),
		string(NumberInput),
		string(PhotoInput),
	}), nil
}

type NoteType string

const (
	PHOTO = NoteType("photo")
	AUDIO = NoteType("audio")
	VIDEO = NoteType("video")
	TEXT  = NoteType("text")
)

func (b NoteType) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString([]string{string(PHOTO), string(AUDIO), string(VIDEO), string(TEXT)}), nil
}

type Posterity struct {
	At *time.Time `gorm:"type:datetime" fake:"{pastdate}"`
	By *string    `fake:"{email}"`
}

type Pole struct {
	ID                uint32  `gorm:"primarykey"`
	Latitude          float32 `fake:"{latitude}"`
	Longitude         float32 `fake:"{longitude}"`
	Street            string  `fake:"{street}"`
	Bearing           Bearing
	Kind              *string
	Height            float32 `fake:"{float32range:180,720}"`
	Locked            bool
	Created           Posterity          `gorm:"embedded;embeddedPrefix:created_"`
	Updated           Posterity          `gorm:"embedded;embeddedPrefix:updated_"`
	Deleted           Posterity          `gorm:"embedded;embeddedPrefix:deleted_" fake:"skip"`
	PoleInstallations []PoleInstallation `fake:"skip"`
}

type PoleInstallation struct {
	ID        uint32 `gorm:"primaryKey"`
	PoleID    uint32
	UbiHubSn  *string    `gorm:"column:ubihub_sn"`
	CamerasSn *string    `gorm:"column:cameras_sn"`
	Start     *time.Time `gorm:"type:datetime" fake:"{pastdate}"`
	End       *time.Time `gorm:"type:datetime" fake:"{pastdate}"`
	Created   Posterity  `gorm:"embedded;embeddedPrefix:created_"`
	Updated   Posterity  `gorm:"embedded;embeddedPrefix:updated_"`
	Pole      *Pole      `fake:"skip"`
}

type PoleAudit struct {
	ID               uint32           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	InstallationID   uint32           `gorm:"column:pole_installation_id;not null" json:"installation_id"`
	Attempt          uint32           `gorm:"column:attempt;not null" json:"attempt"`
	State            State            `gorm:"column:state;not null;default:pending" json:"state"`
	Summary          string           `gorm:"column:summary" json:"summary"`
	Auditor          string           `gorm:"column:auditor;not null" json:"auditor"`
	Created          Posterity        `gorm:"embedded;embeddedPrefix:created_"`
	Updated          Posterity        `gorm:"embedded;embeddedPrefix:updated_"`
	PoleInstallation PoleInstallation `fake:"skip"`
	PoleAuditable    PoleAuditable    `fake:"skip"`
}

type PoleAuditNote struct {
	ID        uint32    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AuditID   uint32    `gorm:"column:pole_audit_id;not null" json:"audit_id"`
	Kind      string    `gorm:"column:kind;not null;default:photo" json:"kind"`
	Datum     string    `gorm:"column:datum;not null" json:"datum"`
	Created   Posterity `gorm:"embedded;embeddedPrefix:created_"`
	Updated   Posterity `gorm:"embedded;embeddedPrefix:updated_"`
	PoleAudit PoleAudit `fake:"skip"`
}

type PoleAuditQuestion struct {
	ID            uint32        `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Device        string        `gorm:"column:device;not null" json:"device"`
	Position      uint32        `gorm:"column:position;not null" json:"position"`
	Question      string        `gorm:"column:question;not null" json:"question"`
	Input         string        `gorm:"column:input;not null" json:"input"`
	Answer        string        `gorm:"column:answer;not null" json:"answer"`
	Created       Posterity     `gorm:"embedded;embeddedPrefix:created_"`
	Updated       Posterity     `gorm:"embedded;embeddedPrefix:updated_"`
	Deleted       Posterity     `gorm:"embedded;embeddedPrefix:deleted_" fake:"skip"`
	PoleAuditable PoleAuditable `fake:"skip"`
}

type PoleAuditQuestionAnswer struct {
	ID            uint32        `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Datum         string        `gorm:"column:datum;not null" json:"datum"`
	Created       Posterity     `gorm:"embedded;embeddedPrefix:created_"`
	PoleAuditable PoleAuditable `fake:"skip"`
}

type PoleAuditable struct {
	AuditID    uint32 `gorm:"column:pole_audit_id;not null" json:"audit_id"`
	QuestionID uint32 `gorm:"column:pole_audit_question_id;not null" json:"question_id"`
	AnswerID   uint32 `gorm:"column:pole_audit_question_answer_id" json:"answer_id"`
}
