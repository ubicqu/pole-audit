package model

import (
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
	db.Instance.Create(&poleInstallation)

	var p Pole
	db.Instance.Preload("PoleInstallations").First(&p, "id = ?", pole.ID)
	fun.PrettyPrintln(p)

	var pi PoleInstallation
	db.Instance.Preload("Pole").First(&pi, "id = ?", poleInstallation.ID)
	fun.PrettyPrintln(pi)
}
