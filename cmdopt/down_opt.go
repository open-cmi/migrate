package cmdopt

import (
	"flag"
	"fmt"
	"os"

	"github.com/open-cmi/migrate/config"
	"github.com/open-cmi/migrate/global"
)

var configfile string = ""
var migratedir string = ""
var count int = -1

// DownOpt down operation
type DownOpt struct {
}

// Run down operation
func (o *DownOpt) Run() error {

	downCmd := flag.NewFlagSet("down", flag.ExitOnError)
	downCmd.StringVar(&configfile, "config", configfile, "config file, default ./etc/db.json")
	downCmd.StringVar(&migratedir, "migrations", migratedir, "migration directory")
	downCmd.IntVar(&count, "count", count, "migrate count")

	downCmd.Parse(os.Args[2:])

	if configfile == "" {
		configfile = GetDefaultConfigFile()
	}
	err := config.Init(configfile)
	if err != nil {
		fmt.Printf("init config failed: %s\n", err.Error())
		return err
	}
	if migratedir == "" {
		SetMigrateMode("go")
	} else {
		SetMigrateMode("sql")
		SetMigrateDir(migratedir)
	}

	db := global.DB
	co := &CurrentOpt{}
	migrations := co.GetMigrationList()
	if len(migrations) == 0 {
		fmt.Printf("no migration to down\n")
		return nil
	}

	if count == -1 {
		count = len(migrations)
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

	return nil
}
