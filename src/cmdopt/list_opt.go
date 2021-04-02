package cmdopt

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/open-cmi/goutils/common"
)

type ListOpt struct {
}

func (o *ListOpt) GetMigrationList() (migrations []SeqInfo) {
	// find migrations dir
	rp := common.GetRootPath()
	files, err := ioutil.ReadDir(filepath.Join(rp, "migrations"))
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
