package cmdopt

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/open-cmi/goutils/common"
)

var template string = `
package migrations

import (
	"fmt"

	"github.com/open-cmi/migrate/cmdopt"
	"github.com/open-cmi/migrate/global"
)

// ChangeMeInstance migrate
type ChangeMeInstance struct {
	Name string
}

// Up up migrate
func (mi ChangeMeInstance) Up() error {
	db := global.DB

	fmt.Println("example up")
	fmt.Println(db)
	return nil
}

// Down down migrate
func (mi ChangeMeInstance) Down() error {
	db := global.DB

	fmt.Println("example down")
	fmt.Println(db)
	return nil
}

func init() {
	cmdopt.MigrationList = append(cmdopt.MigrationList, cmdopt.SeqInfo{
		Seq:         "00000000000000",
		Description: "example",
		Ext:         "go",
		Instance:    ChangeMeInstance{},
	})
}

`

// GenerateOpt generate opt
type GenerateOpt struct {
	Args []string
}

// Run run
func (g *GenerateOpt) Run() {
	// 从template模版读取字符串，然后替换日期
	rt := common.Getwd()

	/*
		templateFile := filepath.Join(rt, "template", "example.go")

		rf, err := os.Open(templateFile)
		if err != nil {
			fmt.Printf("open template file failed, please confirm %s is exist.", templateFile)
			return
		}*/

	t := time.Now()
	date := fmt.Sprintf("%4d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	wfile := fmt.Sprintf("%s_%s.go", date, g.Args[0])
	wfilepath := filepath.Join(rt, "migrations", wfile)
	fmt.Println(date, wfile)
	wf, err := os.OpenFile(wfilepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("create faile failed in migrations")
		return
	}

	/*
		content, err := io.ReadAll(rf)
		if err != nil {
			fmt.Printf("read file content failed\n")
			return
		}*/
	content := template

	newcontent := strings.Replace(content, "00000000000000", date, -1)
	newcontent = strings.Replace(newcontent, "example", g.Args[0], -1)
	newcontent = strings.Replace(newcontent, "MigrateInstance", "ChangeMeInstance", -1)
	io.WriteString(wf, newcontent)
	return
}
