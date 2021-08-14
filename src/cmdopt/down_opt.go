package cmdopt

import (
	"fmt"
	"strconv"

	"github.com/open-cmi/migrate/global"
)

// DownOpt down operation
type DownOpt struct {
	Args []string
}

// Run down operation
func (o *DownOpt) Run() {
	db := global.DB
	co := &CurrentOpt{}
	migrations := co.GetMigrationList()
	if len(migrations) == 0 {
		fmt.Printf("no migration to down\n")
		return
	}

	var count int = len(migrations)
	if len(o.Args) != 0 {
		ct, err := strconv.Atoi(o.Args[0])
		if err == nil && ct > 0 {
			count = ct
		}
	}

	for idx := len(migrations) - 1; idx >= 0 && count > 0; idx-- {
		m := migrations[idx]
		fmt.Printf("start to down migrate: %s %s\n", m.Seq, m.Description)

		var err error
		if m.Ext == "sql" {
			err = ExecSQLMigrate(db, &m, "down")
		}
		if err == nil {
			dbexec := fmt.Sprintf("delete from migrations where seq='%s'", m.Seq)
			db.Exec(dbexec)
			fmt.Println("successfully!!")
		}
		count--
	}

	return
}
