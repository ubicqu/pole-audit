package db

import "testing"

func TestDB(t *testing.T) {
	var results []map[string]interface{}
	Instance.Raw("show databases").Scan(&results)
	for _, result := range results {
		if result["Database"] == "mysql" {
			return
		}
	}
	t.Fail()
}
