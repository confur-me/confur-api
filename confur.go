package main

import (
	"flag"
	"fmt"
	"github.com/confur-me/confur-api/api/migrator"
	"github.com/confur-me/confur-api/lib/config"
	"github.com/confur-me/confur-api/server"
	"os"
)

var configFile *string

func init() {
	const (
		help_description   = "print this help"
		config_description = "path to config file"
		usage              = `
Usage:
  confur -config=/path/to/config.yml start
  confur -config=/path/to/config.yml migrate
`
	)
	flag.Bool("h", false, help_description)
	configFile = flag.String("c", "confur.yml", config_description)

	flag.Usage = func() {
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Print(usage)
		os.Exit(0)
	}
}

func main() {
	args := parseFlags()
	if err := config.Read(*configFile); err != nil {
		os.Exit(0)
	}

	switch args[0] {
	case "start":
		app := server.Application{}
		app.Run()
	case "db:create":
		migrator.DbCreate()
	case "db:migrate":
		migrator.DbMigrate()
	case "db:drop":
		migrator.DbDrop()
	default:
		flag.Usage()
	}
}

func parseFlags() []string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}
	return args
}
