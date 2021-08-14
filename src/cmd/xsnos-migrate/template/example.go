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
