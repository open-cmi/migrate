package cmdopt

import (
	"flag"
	"fmt"
	"os"

	"github.com/open-cmi/migrate/config"
	"github.com/open-cmi/migrate/global"
)

// CurrentOpt get current migration
type CurrentOpt struct {
}

// GetMigrationList get migration list
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

// Run run
func (o *CurrentOpt) Run() error {
	currentCmd := flag.NewFlagSet("current", flag.ExitOnError)
	currentCmd.StringVar(&configfile, "config", configfile, "config file, default ./etc/db.json")

	currentCmd.Parse(os.Args[2:])
	if configfile == "" {
		configfile = GetDefaultConfigFile()
	}
	err := config.Init(configfile)
	if err != nil {
		fmt.Printf("init config failed: %s\n", err.Error())
		return err
	}
	migrations := o.GetMigrationList()
	if len(migrations) == 0 {
		fmt.Printf("no migrations\n")
		return nil
	}
	for _, m := range migrations {
		fmt.Printf("%s %s %s\n", m.Seq, m.Description, m.Ext)
	}
	return nil
}
