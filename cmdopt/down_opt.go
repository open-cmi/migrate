package cmdopt

import (
	"flag"
	"fmt"
	"os"

	"github.com/open-cmi/migrate/config"
	"github.com/open-cmi/migrate/global"
)

// DownOpt down operation
type DownOpt struct {
}

// Run down operation
func (o *DownOpt) Run() error {

	downCmd := flag.NewFlagSet("down", flag.ExitOnError)
	downCmd.StringVar(&configfile, "config", configfile, "config file, default ./etc/db.json")
	downCmd.StringVar(&format, "format", format, "format, go or sql")
	downCmd.StringVar(&input, "input", input, "if use sql, should specify sql directory")
	downCmd.IntVar(&count, "count", count, "migrate count")

	err := downCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if configfile == "" {
		configfile = GetDefaultConfigFile()
	}
	err = config.Init(configfile)
	if err != nil {
		fmt.Printf("init config failed: %s\n", err.Error())
		return err
	}
	if format == "" || format == "go" {
		SetMigrateMode("go")
	} else {
		SetMigrateMode("sql")
	}

	if input != "" {
		SetMigrateDir(input)
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
		} else if m.Ext == "go" {
			err = ExecGoMigrate(db, &m, "down")
		}
		if err == nil {
			dbexec := fmt.Sprintf("delete from migrations where seq='%s'", m.Seq)
			db.Exec(dbexec)
			fmt.Println("successfully!!")
		} else {
			fmt.Printf("migrate down failed: %s\n", err.Error())
		}
		count--
	}

	return nil
}
