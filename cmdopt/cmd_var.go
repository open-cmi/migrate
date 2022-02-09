package cmdopt

// file format, use go or use sql
var format string = ""

// where the output file locate wheen use generate command
var output string = ""

// when use sql format, you should
// use -input specify the path of sql directory
var input string = ""

// where db config saved
var configfile string = ""

// when up or down, you can specify the count
var count int = -1

// GoMigrationList migration list
var GoMigrationList []SeqInfo

// Service service belong
var Service string = ""
