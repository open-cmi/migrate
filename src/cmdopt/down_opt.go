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

type DownOpt struct {
	Args []string
}

func (o *DownOpt) Run() {
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
		rp := common.GetRootPath()
		sqlfile := m.Seq + "_" + m.Description + ".down.sql"
		downfile := filepath.Join(rp, "migrations", sqlfile)
		if !goutils.IsExist(downfile) {
			fmt.Printf("migrate file %s not exist\n", sqlfile)
			return
		}
		// exec file content
		f, err := os.Open(downfile)
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
			dbexec := fmt.Sprintf("delete from migrations where seq='%s'", m.Seq)
			dbsql.DBSql.Exec(dbexec)
			fmt.Println("successfully!!")
		}
		count--
	}

	return
}
