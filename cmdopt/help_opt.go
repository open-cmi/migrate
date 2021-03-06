package cmdopt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// HelpOpt help operation
type HelpOpt struct {
}

// Run run help command
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

	fmt.Printf("%s subcommand -h for more subcommand info\n", proc)
}
