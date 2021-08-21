package cmdopt

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/open-cmi/migrate/config"
	"github.com/open-cmi/migrate/global"
)

// UpOpt up operation
type UpOpt struct {
	Args []string
}

// Run up operation run
func (o *UpOpt) Run() error {
	upCmd := flag.NewFlagSet("up", flag.ContinueOnError)
	upCmd.StringVar(&migratedir, "migrations", migratedir, "migration directory, if migration is empty, use go mode")
	upCmd.StringVar(&configfile, "config", configfile, "config file, default ./etc/db.json")

	upCmd.IntVar(&count, "count", count, "migrate up count")

	upCmd.Parse(os.Args[2:])

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

	lo := &ListOpt{}
	filelist := lo.GetMigrationList()

	if count == -1 {
		count = len(filelist)
	}

	// find start index
	startIndex := 0
	if len(migrations) != 0 {
		var find bool
		latest := migrations[len(migrations)-1]
		for idx, fm := range filelist {
			if strings.Compare(latest.Seq, fm.Seq) < 0 {
				startIndex = idx
				find = true
				break
			}
		}
		if !find {
			fmt.Printf("no migrations\n")
			return nil
		}
	}

	for idx := startIndex; idx < len(filelist) && count > 0; idx++ {
		fl := filelist[idx]
		fmt.Printf("start to up migrate: %s %s\n", fl.Seq, fl.Description)

		if fl.Ext == "sql" {
			err = ExecSQLMigrate(db, &fl, "up")
		} else if fl.Ext == "go" {
			err = ExecGoMigrate(db, &fl, "up")
		}

		if err == nil {
			dbexec := fmt.Sprintf("insert into migrations(seq, description, service, ext) values('%s','%s','%s','%s')",
				fl.Seq, fl.Description, Service, fl.Ext)
			_, err = db.Exec(dbexec)
			if err != nil {
				fmt.Printf("migrate %s %s failed, %s\n", fl.Seq, fl.Description, err.Error())
				break
			}
			fmt.Println("successfully!!")
		} else {
			fmt.Printf("migrate failed, error: %s\n", err.Error())
			break
		}
		count--
	}
	return err
}
