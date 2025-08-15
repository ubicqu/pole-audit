package main

import (
	"gorm.io/gen"
	"pole-audit/pkg/db"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../pole-audit/pkg/model/gen/query",
		ModelPkgPath: "../pole-audit/pkg/model/gen/models",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		WithUnitTest: true,
	})

	g.UseDB(db.Instance)

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		// Generate struct `User` based on table `users`
		//g.GenerateModel("poles"),
		//g.GenerateModel("pole_installations"),
		//g.GenerateModel("pole_audits"),
		//g.GenerateModel("pole_audit_notes"),
		//g.GenerateModel("pole_audit_questions"),
		//g.GenerateModel("pole_audit_question_answers"),
		//g.GenerateModel("pole_auditables"),
		g.GenerateAllTable()...,
	)

	// Generate the code
	g.Execute()
}
