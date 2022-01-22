package cmdopt

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/open-cmi/goutils/pathutil"
)

var gotemplate string = `
package migration

import (
	"fmt"

	"github.com/open-cmi/migrate"
	"github.com/open-cmi/migrate/cmdopt"
	"github.com/open-cmi/migrate/global"
)

// ChangeMeInstance migrate
type ChangeMeInstance struct {
}

// Up up migrate
func (mi ChangeMeInstance) Up() error {
	db := global.DB

	sqlClause := fmt.Sprintf(`+"`"+`
		CREATE TABLE IF NOT EXISTS template (
			id char(64) NOT NULL PRIMARY KEY,
			name VARCHAR(256) NOT NULL unique DEFAULT ''
		)
	`+"`"+`)
	_, err := db.Exec(sqlClause)
	return err
}

// Down down migrate
func (mi ChangeMeInstance) Down() error {
	db := global.DB

	sqlClause := fmt.Sprintf("DROP TABLE IF EXISTS template")
	_, err := db.Exec(sqlClause)
	return err
}

func init() {
	migrate.Register(&cmdopt.SeqInfo{
		Seq:         "00000000000000",
		Description: "example",
		Ext:         "go",
		Instance:    ChangeMeInstance{},
	})
}

`

var sqlup string = ``
var sqldown string = ``

// GenerateOpt generate opt
type GenerateOpt struct {
}

var name string = ""

// Run run
func (g *GenerateOpt) Run() error {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	generateCmd.StringVar(&output, "output", output, "output directory")
	generateCmd.StringVar(&format, "format", output, "format, go or sql")
	generateCmd.StringVar(&name, "name", name, "script name")

	generateCmd.Parse(os.Args[2:])

	if name == "" {
		generateCmd.Usage()
		return errors.New("name cant't be empty")
	}

	if format == "" || format == "go" {
		SetMigrateMode("go")
	} else {
		SetMigrateMode("sql")
	}

	if output == "" {
		rt := pathutil.Getwd()
		output = filepath.Join(rt, "migration")
	}

	SetMigrateDir(output)

	t := time.Now()
	date := fmt.Sprintf("%4d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	if MigrateMode == "go" {
		// 从template模版读取字符串，然后替换日期
		wfile := fmt.Sprintf("%s_%s.go", date, name)
		wfilepath := filepath.Join(MigrateDir, wfile)
		wf, err := os.OpenFile(wfilepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("create file failed in migrations, please confirm migrations directory is exist\n")
			return err
		}

		content := gotemplate

		newcontent := strings.Replace(content, "00000000000000", date, -1)
		newcontent = strings.Replace(newcontent, "example", name, -1)
		io.WriteString(wf, newcontent)
	} else {
		// 从template模版读取字符串，然后替换日期
		upfile := fmt.Sprintf("%s_%s.up.sql", date, name)
		wfilepath := filepath.Join(MigrateDir, upfile)
		wf, err := os.OpenFile(wfilepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("create file failed in migrations\n")
			return err
		}
		io.WriteString(wf, sqlup)

		downfile := fmt.Sprintf("%s_%s.down.sql", date, name)
		wfilepath = filepath.Join(MigrateDir, downfile)
		wf, err = os.OpenFile(wfilepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("create file failed in migrations\n")
			return err
		}
		io.WriteString(wf, sqldown)
	}
	fmt.Printf("generate file successfully!\n")
	return nil
}
