package cmdopt

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// ListOpt list operation
type ListOpt struct {
}

// GoMigrationList migration list
var GoMigrationList []SeqInfo

// SQLMigrationList migrate list
var SQLMigrationList []SeqInfo

// GetMigrationList get migration list
func (o *ListOpt) GetMigrationList() (migrations []SeqInfo) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.StringVar(&migratedir, "migrations", migratedir, "migration directory, if migration is emptry, use go mode")

	listCmd.Parse(os.Args[2:])

	if migratedir == "" {
		SetMigrateMode("go")
	} else {
		SetMigrateMode("sql")
		SetMigrateDir(migratedir)
	}

	if MigrateMode == "go" {
		return GoMigrationList
	}

	files, err := ioutil.ReadDir(MigrateDir)
	if err != nil {
		return
	}
	for _, finfo := range files {
		fname := finfo.Name()
		arr := strings.Split(fname, ".")
		if arr[1] != "up" {
			continue
		}
		var item SeqInfo
		sd := strings.SplitN(arr[0], "_", 2)
		if len(sd) != 2 {
			return
		}

		item.Seq = sd[0]
		item.Description = sd[1]
		item.Ext = arr[2]
		migrations = append(migrations, item)
	}
	if len(migrations) != 0 {
		sort.SliceStable(migrations, func(i, j int) bool {
			cmp := strings.Compare(migrations[i].Seq, migrations[j].Seq)
			return cmp == -1
		})
	}
	return migrations
}

// Run list operation run
func (o *ListOpt) Run() {
	migrations := o.GetMigrationList()
	if len(migrations) == 0 {
		fmt.Printf("no migrations found\n")
		return
	}

	for _, m := range migrations {
		fmt.Printf("%s %s %s\n", m.Seq, m.Description, m.Ext)
	}
}
