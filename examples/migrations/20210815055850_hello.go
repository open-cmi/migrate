
package migrations

import (
	"fmt"

	"github.com/open-cmi/migrate"
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

	fmt.Println("hello up")
	fmt.Println(db)
	return nil
}

// Down down migrate
func (mi ChangeMeInstance) Down() error {
	db := global.DB

	fmt.Println("hello down")
	fmt.Println(db)
	return nil
}

func init() {
	migrate.Register(&cmdopt.SeqInfo{
		Seq:         "20210815055850",
		Description: "hello",
		Ext:         "go",
		Instance:    ChangeMeInstance{},
	})
}

