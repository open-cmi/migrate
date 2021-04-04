package cmdopt

import (
	"fmt"

	"github.com/open-cmi/migrate/global"
)

type CurrentOpt struct {
}

func (o *CurrentOpt) GetMigrationList() (migrations []SeqInfo) {
	db := global.DB
	dbquery := `select * from migrations order by seq asc`
	r, err := db.Query(dbquery)
	if err != nil {
		// select error , init table
		var init string = `CREATE TABLE IF NOT EXISTS migrations (
      seq varchar(14) UNIQUE NOT NULL,
      description varchar(100) NOT NULL default '',
			ext varchar(100) NOT NULL default 'sql'
    )`
		_, err = db.Exec(init)
		if err != nil {
			fmt.Println("init migrations table failed")
		}
		return
	}

	for r.Next() {
		var row SeqInfo
		err = r.Scan(&row.Seq, &row.Description, &row.Ext)
		if err != nil {
			break
		}
		migrations = append(migrations, row)
	}
	return migrations
}

func (o *CurrentOpt) Run() {
	migrations := o.GetMigrationList()
	if len(migrations) == 0 {
		fmt.Printf("no migrations\n")
		return
	}
	for _, m := range migrations {
		fmt.Printf("%s %s\n", m.Seq, m.Description)
	}
}
