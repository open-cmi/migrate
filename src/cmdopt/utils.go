package cmdopt

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"plugin"
	"strings"

	"github.com/open-cmi/goutils"
)

type SeqInfo struct {
	Seq         string
	Description string
	Ext         string
}

func ExecSqlFile(db *sql.DB, sqlfile string) (err error) {
	if !goutils.IsExist(sqlfile) {
		errmsg := fmt.Sprintf("migrate file %s not exist\n", sqlfile)
		return errors.New(errmsg)
	}

	// exec file content
	f, err := os.Open(sqlfile)
	if err != nil {
		errmsg := fmt.Sprintf("open %s failed\n", sqlfile)
		return errors.New(errmsg)
	}

	sqlContent, err := ioutil.ReadAll(f)
	if err != nil {
		errmsg := fmt.Sprintf("read %s failed\n", sqlfile)
		return errors.New(errmsg)
	}

	arr := strings.SplitAfter(string(sqlContent), ";")
	for _, sentence := range arr {
		if strings.Trim(sentence, "") == "" {
			continue
		}
		_, err = db.Exec(sentence)
		if err != nil {
			errmsg := fmt.Sprintf("migrate failed %s\n", err.Error())
			return errors.New(errmsg)
		}
	}
	return
}

func ExecSoFile(db *sql.DB, sqlfile string) (err error) {
	p, err := plugin.Open(sqlfile)
	if err != nil {
		errmsg := fmt.Sprintf("open %s failed", sqlfile)
		return errors.New(errmsg)
	}
	migrate, err := p.Lookup("Migrate")
	if err != nil {
		errmsg := "look up Migrate function failed"
		return errors.New(errmsg)
	}

	migrate.(func(*sql.DB))(db)
	return
}

func ExecGoFile(db *sql.DB, seqmod SeqInfo, migrate string) {

}
