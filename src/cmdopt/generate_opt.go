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

// GenerateOpt generate opt
type GenerateOpt struct {
	Args []string
}

// Run run
func (g *GenerateOpt) Run() {
	// 从template模版读取字符串，然后替换日期
	rt := common.GetRootPath()

	templateFile := filepath.Join(rt, "template", "example.go")
	rf, err := os.Open(templateFile)
	if err != nil {
		fmt.Printf("open template file failed, please confirm %s is exist.", templateFile)
		return
	}

	t := time.Now()
	date := fmt.Sprintf("%4d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	wfile := fmt.Sprintf("%s_%s.go", date, g.Args[0])
	wfilepath := filepath.Join(rt, "src", "migrations", wfile)
	fmt.Println(date, wfile)
	wf, err := os.OpenFile(wfilepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("create faile failed in migrations")
		return
	}

	content, err := io.ReadAll(rf)
	if err != nil {
		fmt.Printf("read file content failed\n")
		return
	}

	newcontent := strings.Replace(string(content), "00000000000000", date, -1)
	newcontent = strings.Replace(newcontent, "example", g.Args[0], -1)
	newcontent = strings.Replace(newcontent, "MigrateInstance", "ChangeMeInstance", -1)
	io.WriteString(wf, newcontent)
	return
}
