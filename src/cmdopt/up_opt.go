package cmdopt

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/migrate/global"
)

// UpOpt up operation
type UpOpt struct {
	Args []string
}

// Run up operation run
func (o *UpOpt) Run() {
	db := global.DB
	co := &CurrentOpt{}
	migrations := co.GetMigrationList()

	lo := &ListOpt{}
	filelist := lo.GetMigrationList()

	var count int = len(filelist)
	if len(o.Args) != 0 {
		ct, err := strconv.Atoi(o.Args[0])
		if err == nil && ct > 0 {
			count = ct
		}
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
			return
		}
	}

	for idx := startIndex; idx < len(filelist) && count > 0; idx++ {
		fl := filelist[idx]
		fmt.Printf("start to migrate: %s %s\n", fl.Seq, fl.Description)

		var err error
		if fl.Ext == "sql" {
			rp := common.GetRootPath()
			sqlfile := fl.Seq + "_" + fl.Description + ".up." + fl.Ext
			upfile := filepath.Join(rp, "migrations", sqlfile)
			err = ExecSqlFile(db, upfile)
		} else if fl.Ext == "go" {
			//err = ExecGoScript(db)
		}

		if err == nil {
			dbexec := fmt.Sprintf("insert into migrations(seq, description) values('%s','%s')", fl.Seq, fl.Description)
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
}
