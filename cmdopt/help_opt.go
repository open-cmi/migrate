package cmdopt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type HelpOpt struct {
}

func (ho *HelpOpt) Run() {
	executable, _ := os.Executable()
	proc := filepath.Base(executable)
	if strings.HasPrefix(executable, os.TempDir()) {
		proc = "go run main.go"
	}
	fmt.Printf("usage: %s subcommand option\n\n", proc)
	fmt.Printf("subcommand:\n")
	fmt.Printf("  current:     show current database migrations\n")
	fmt.Printf("  list:        show migrations file list\n")
	fmt.Printf("  up:          migrate number files\n")
	fmt.Printf("  down:        rollback number database\n")
	fmt.Printf("  generate:    generate migrate file\n")

	fmt.Printf("option:\n")
	fmt.Printf("  up [num]:\n")
	fmt.Printf("  down [num]:\n")
	fmt.Printf("  generate {filename}\n")
}
