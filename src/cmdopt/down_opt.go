package cmdopt

import (
	"fmt"
	"path/filepath"
	"strconv"

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
		sqlfile := m.Seq + "_" + m.Description + ".down." + m.Ext
		downfile := filepath.Join(rp, "migrations", sqlfile)
		var err error
		if m.Ext == "sql" {
			err = ExecSqlFile(dbsql.DBSql, downfile)
		} else if m.Ext == "so" {
			err = ExecSoFile(dbsql.DBSql, downfile)
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
