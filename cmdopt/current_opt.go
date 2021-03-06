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

// GetInstanceFromSeq get instance from seq
func GetInstanceFromSeq(seq string) interface{} {
	for _, ins := range GoMigrationList {
		if ins.Seq == seq {
			return ins.Instance
		}
	}
	return nil
}

// GetMigrationList get migration list
func (o *CurrentOpt) GetMigrationList() (migrations []SeqInfo) {
	db := global.DB
	dbquery := fmt.Sprintf(`select * from migrations where service='%s' order by seq asc`, Service)
	r, err := db.Query(dbquery)
	if err != nil {
		// select error , init table
		var init string = `CREATE TABLE IF NOT EXISTS migrations (
							seq varchar(14) UNIQUE NOT NULL,
							description varchar(100) NOT NULL default '',
							ext varchar(100) NOT NULL default 'sql',
							service varchar(100) NOT NULL default ''
							)`
		_, err = db.Exec(init)
		if err != nil {
			fmt.Println("init migrations table failed")
		}
		return
	}

	for r.Next() {
		var row SeqInfo
		err = r.Scan(&row.Seq, &row.Description, &row.Ext, &row.Service)
		if err != nil {
			break
		}
		if row.Ext == "go" {
			row.Instance = GetInstanceFromSeq(row.Seq)
			if row.Instance == nil {
				fmt.Printf("\033[1;31;40mWarning\033[0m seq %s's Ext is go, but it's instance is not found\n", row.Seq)
			}
		}
		migrations = append(migrations, row)
	}
	return migrations
}

// Run run
func (o *CurrentOpt) Run() error {
	currentCmd := flag.NewFlagSet("current", flag.ExitOnError)
	currentCmd.StringVar(&configfile, "config", configfile, "config file, default ./etc/db.json")

	err := currentCmd.Parse(os.Args[2:])
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
