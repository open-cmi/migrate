package migration

import (
	"github.com/open-cmi/migrate"
	"github.com/open-cmi/migrate/cmdopt"
	"github.com/open-cmi/migrate/global"
)

// TemplateInstance migrate
type TemplateInstance struct {
}

// Up up migrate
func (mi TemplateInstance) Up() error {
	db := global.DB

	sqlClause := `
		CREATE TABLE IF NOT EXISTS template (
			id char(64) NOT NULL PRIMARY KEY,
			name VARCHAR(256) NOT NULL unique DEFAULT ''
		)
	`
	_, err := db.Exec(sqlClause)
	return err
}

// Down down migrate
func (mi TemplateInstance) Down() error {
	db := global.DB

	sqlClause := `DROP TABLE IF EXISTS template`
	_, err := db.Exec(sqlClause)
	return err
}

func init() {
	migrate.Register(&cmdopt.SeqInfo{
		Seq:         "20220210071830",
		Description: "test_tem",
		Ext:         "go",
		Instance:    TemplateInstance{},
	})
}
