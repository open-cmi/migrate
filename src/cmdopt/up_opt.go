package cmdopt

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/database/dbsql"
)

type UpOpt struct {
	Args []string
}

func (o *UpOpt) Run() {
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
			}
		}
		if !find {
			fmt.Printf("no migrations\n")
			return
		}
	}

	for idx := startIndex; idx < len(filelist) && count > 0; idx++ {
		fl := filelist[idx]
		fmt.Printf("start to up migrate: %s %s\n", fl.Seq, fl.Description)
		rp := common.GetRootPath()
		sqlfile := fl.Seq + "_" + fl.Description + ".up.sql"
		upfile := filepath.Join(rp, "migrations", sqlfile)
		if !goutils.IsExist(upfile) {
			fmt.Printf("migrate file %s not exist\n", sqlfile)
			return
		}
		// exec file content
		f, err := os.Open(upfile)
		if err != nil {
			fmt.Printf("open %s failed\n", sqlfile)
			return
		}
		sqlContent, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Printf("read %s failed\n", sqlfile)
			return
		}

		arr := strings.SplitAfter(string(sqlContent), ";")
		for _, sentence := range arr {
			if strings.Trim(sentence, "") == "" {
				continue
			}
			_, err = dbsql.DBSql.Exec(sentence)
			if err != nil {
				break
			}
		}
		if err == nil {
			dbexec := fmt.Sprintf("insert into migrations(seq, description) values('%s','%s')", fl.Seq, fl.Description)
			_, err = dbsql.DBSql.Exec(dbexec)
			if err != nil {
				fmt.Printf("migrate %s failed, %s\n", sqlfile, err.Error())
				return
			}
			fmt.Println("successfully!!")
		}
		count--
	}

	return
}
