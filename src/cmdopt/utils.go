package cmdopt

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"reflect"
	"strings"

	"github.com/open-cmi/goutils"
	"github.com/open-cmi/goutils/common"
)

// SeqInfo migrate seq info
type SeqInfo struct {
	Seq         string
	Description string
	Ext         string
	Instance    interface{}
}

// ExecSQLMigrate exec sql mod
func ExecSQLMigrate(db *sql.DB, si *SeqInfo, updown string) (err error) {
	rp := common.GetRootPath()
	sqlfile := si.Seq + "_" + si.Description + "." + updown + "." + si.Ext
	sqlfilepath := filepath.Join(rp, "migrations", sqlfile)

	if !goutils.IsExist(sqlfilepath) {
		errmsg := fmt.Sprintf("migrate file %s not exist\n", sqlfilepath)
		return errors.New(errmsg)
	}

	// exec file content
	f, err := os.Open(sqlfilepath)
	if err != nil {
		errmsg := fmt.Sprintf("open %s failed\n", sqlfilepath)
		return errors.New(errmsg)
	}

	sqlContent, err := ioutil.ReadAll(f)
	if err != nil {
		errmsg := fmt.Sprintf("read %s failed\n", sqlfilepath)
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

// ExecGoMigrate exec go migrate
func ExecGoMigrate(db *sql.DB, si *SeqInfo, updown string) (err error) {
	instance := si.Instance
	ref := reflect.ValueOf(instance)
	var fun reflect.Value
	if updown == "up" {
		fun = ref.MethodByName("Up")
	} else if updown == "down" {
		fun = ref.MethodByName("Down")
	}
	var params []reflect.Value = []reflect.Value{}
	retlist := fun.Call(params)
	if retlist[0].Interface() != nil {
		return retlist[0].Interface().(error)
	}
	return nil
}

// ExecSoFile exec plugin so file
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
